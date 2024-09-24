package main

import (
	"log"
	"context"

	// "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/kubediscovery/platform-customer-registry/configs"
	"github.com/kubediscovery/platform-customer-registry/pkg/otelpkg"
	"github.com/kubediscovery/platform-customer-registry/pkg/cache"
	"github.com/kubediscovery/platform-customer-registry/internal/core/service"
	_ "github.com/kubediscovery/platform-customer-registry/docs/swagger"
	handler "github.com/kubediscovery/platform-customer-registry/internal/infra/handler/rest"
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


	resitryService, err := service.NewRegistryService(nil, &cc, otl)
	if err != nil {
		log.Fatalln(err)
	}

	handler.NewCustomerRegistryHandlerHttp(resitryService, otl, rest.RouterGroup, rest.ValidateToken)
	if err != nil {
		log.Fatalln("error is: ", err.Error())
	}

}
