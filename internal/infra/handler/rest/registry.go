package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kubediscovery/platform-customer-registry/internal/core/entity"
	"github.com/kubediscovery/platform-customer-registry/internal/core/service"
	"github.com/kubediscovery/platform-customer-registry/pkg/otelpkg"
)

type CustomerRegistryHandlerHttpInterface interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
}

type CustomerRegistryHandlerHttp struct {
	Service service.RegistryServiceInterface
	Tracer  *otelpkg.OtelPkgInstrument
}

func NewCustomerRegistryHandlerHttp(svc service.RegistryServiceInterface, otl *otelpkg.OtelPkgInstrument, routerGroup *gin.RouterGroup, middleware ...func(c *gin.Context)) CustomerRegistryHandlerHttpInterface {

	registry := &CustomerRegistryHandlerHttp{
		Service: svc,
		Tracer:  otl,
	}

	registry.handlers(routerGroup, middleware...)

	return registry
}

func (c *CustomerRegistryHandlerHttp) handlers(routerGroup *gin.RouterGroup, middleware ...func(c *gin.Context)) {
	middlewareList := make([]gin.HandlerFunc, len(middleware))
	for i, mw := range middleware {
		middlewareList[i] = mw
	}

	routerGroup.POST("/registry", append(middlewareList, c.Create)...)
	routerGroup.GET("/registry/:id", append(middlewareList, c.GetByID)...)
	routerGroup.GET("/registry", append(middlewareList, c.GetAll)...)
}

// CreateCustomerRegistry    godoc
// @Summary     create a new customer registry
// @Tags        registry
// @Accept       json
// @Produce     json
// @Description create a new registry
// @Success     200 {object} entity.CustomerRegistryResponse
// @Failure     404 {object} string
// @Failure     500 {object} string
// @Router      /registry [post]
func (obj *CustomerRegistryHandlerHttp) Create(c *gin.Context) {
	ctx, span := obj.Tracer.Tracer.Start(c.Request.Context(), "CustomerRegistryHandlerHttp.ListResources")
	defer span.End()

	var registry *entity.CustomerRegistry
	if err := c.ShouldBindJSON(&registry); err != nil {
		span.RecordError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := obj.Service.Create(ctx, &entity.CustomerRegistryResponse{CustomerRegistry: *registry})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found")})
		return
	}

	c.JSON(http.StatusAccepted, result)
}

// GetAllCustomerRegistry    godoc
// @Summary     get all customer registry
// @Tags        registry
// @Accept       json
// @Produce     json
// @Description get all customer registry
// @Success     200 {object} []entity.CustomerRegistryResponse
// @Failure     404 {object} string
// @Failure     500 {object} string
// @Router      /registry [get]
func (obj *CustomerRegistryHandlerHttp) GetAll(c *gin.Context) {
	ctx, span := obj.Tracer.Tracer.Start(c.Request.Context(), "CustomerRegistryHandlerHttp.ListResources")
	defer span.End()

	result, err := obj.Service.GetAll(ctx)
	if err != nil {
		span.RecordError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found")})
		return
	}

	c.JSON(http.StatusAccepted, result)
}

// GetByID      godoc
// @Summary     get a customer registry by ID
// @Tags        registry
// @Accept      json
// @Produce     json
// @Param       id path string true "found"
// @Description get a customer registry by ID
// @Success     200 {object} entity.CustomerRegistryResponse
// @Failure     404 {object} string
// @Failure     500 {object} string
// @Router      /registry/{id} [get]
func (obj *CustomerRegistryHandlerHttp) GetByID(c *gin.Context) {
	_, span := obj.Tracer.Tracer.Start(c.Request.Context(), "CustomerRegistryHandlerHttp.ListResources")
	defer span.End()

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		c.Abort()
		return
	}

	result, err := obj.Service.GetByID(c.Request.Context(), &id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusAccepted, result)
}
