package main

import (
	"context"
	"log"

	// "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/kubediscovery/platform-customer-registry/configs"
	_ "github.com/kubediscovery/platform-customer-registry/docs/swagger"
	"github.com/kubediscovery/platform-customer-registry/internal/core/repository"
	"github.com/kubediscovery/platform-customer-registry/internal/core/service"
	handler "github.com/kubediscovery/platform-customer-registry/internal/infra/handler/rest"
	"github.com/kubediscovery/platform-customer-registry/pkg/cache"
	"github.com/kubediscovery/platform-customer-registry/pkg/kb_db/kb_psql"
	"github.com/kubediscovery/platform-customer-registry/pkg/otelpkg"
	http_server "github.com/kubediscovery/platform-customer-registry/pkg/service_http/server"
)

// @title        platform-customer-registry
// @version      1.0
// @description  This service registry and controller new customer lab

// @contact.name   Rafael Tomelin
// @contact.url    https://local
// @contact.email  contato@synera.com.br

// @schemes   http
// @BasePath  /api
func main() {

	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("error loading configs: %v", err)
	}

	// STARTS OTEL
	ctx := context.Background()
	otl, err := otelpkg.NewOtel(ctx, cfg.FileConfig.ConfigPath, cfg.FileConfig.FileName, cfg.FileConfig.Extentsion)
	if err != nil {
		log.Fatalln(err)
	}

	ctx, span := otl.Tracer.Start(ctx, "main")
	defer span.End()

	defer otl.TracerSdk.Shutdown(ctx)

	_, err = newrelic.NewApplication(
		newrelic.ConfigAppName(otl.Parameters.AppName),
		newrelic.ConfigLicense(otl.Parameters.License),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		log.Fatalln(err)
	}
	// ENDS OTEL

	// STARTS DB
	dbPool, err := kb_psql.NewDBConnection(ctx, cfg.FileConfig.ConfigPath, cfg.FileConfig.FileName, cfg.FileConfig.Extentsion)
	if err != nil {
		log.Fatalln(err)
	}

	defer dbPool.Close()

	err = dbPool.Ping()
	if err != nil {
		log.Fatalln("error is ", err.Error())
	}
	// ENDS DB

	// STARTS CACHE
	cc, err := cache.NewCacheConnection(cfg.FileConfig.ConfigPath, cfg.FileConfig.FileName, cfg.FileConfig.Extentsion)
	if err != nil {
		log.Fatalln(err)
	}
	// ENDS CACHE

	// STARTS HTTP SERVER
	rest, err := http_server.NewRestApi(cfg.FileConfig.ConfigPath, cfg.FileConfig.FileName, cfg.FileConfig.Extentsion)
	if err != nil {
		log.Fatalln(err)
	}
	// ENDS HTTP SERVER

	// STARTS CUSTOMER REGISTRY
	resitryRepository, err := repository.NewRegistryRepository(dbPool, otl)
	if err != nil {
		log.Fatalln(err)
	}
	resitryService, err := service.NewRegistryService(resitryRepository, &cc, otl)
	if err != nil {
		log.Fatalln(err)
	}

	handler.NewCustomerRegistryHandlerHttp(resitryService, otl, rest.RouterGroup, rest.ValidateToken)
	if err != nil {
		log.Fatalln("error is: ", err.Error())
	}
	// ENDS CUSTOMER REGISTRY

	// STARTS LAB DESTROY
	labRepository, err := repository.NewLabDestroyRepository(dbPool, otl)
	if err != nil {
		log.Fatalln(err)
	}
	labService, err := service.NewServiceLabDestroy(&labRepository, &cc, otl)
	if err != nil {
		log.Fatalln(err)
	}

	handler.NewLabDestroyHandlerHttp(&labService, otl, rest.RouterGroup, rest.ValidateToken)
	if err != nil {
		log.Fatalln("error is: ", err.Error())
	}
	// ENDS LAB DESTROY

	rest.Run(rest.Route.Handler())
}
