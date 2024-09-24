package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synera-br/golang-cloud-collector/internal/core/service"
	_ "github.com/synera-br/golang-cloud-collector/internal/core/entity"
	"github.com/synera-br/golang-cloud-collector/pkg/otelpkg"
)

type CustomerRegistryHandlerHttpInterface interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Search(c *gin.Context)
	IsValid(c *gin.Context)
}

type CustomerRegistryHandlerHttp struct {
	Service entity.CustomerRegistryInterface
	Tracer  *otelpkg.OtelPkgInstrument
}

func NewCustomerRegistryHandlerHttp(svc entity.CustomerRegistryInterface, otl *otelpkg.OtelPkgInstrument, routerGroup *gin.RouterGroup, middleware ...func(c *gin.Context)) CustomerRegistryHandlerHttpInterface {

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
}

// CreateCustomerRegistry    godoc
// @Summary     create a new customer registry
// @Tags        registry
// @Accept       json
// @Produce     json
// @Description create a new registry
// @Success     200 {object} interface{}
// @Failure     404 {object} string
// @Failure     500 {object} string
// @Router      /registry [post]
func (obj *CustomerRegistryHandlerHttp) Create(c *gin.Context) {
	ctx, span := obj.Tracer.Tracer.Start(c.Request.Context(), "CustomerRegistryHandlerHttp.ListResources")
	defer span.End()


	var registry *entity.CustomerRegistry
	if err := c.ShouldBindJSON(&registry); err != nil {
		span.RecordError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := obj.Service.Create(ctx, registry)
	if err != nil {
		span.RecordError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found")})
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (obj *CustomerRegistryHandlerHttp) List(c *gin.Context) {
	ctx, span := obj.Tracer.Tracer.Start(c.Request.Context(), "CustomerRegistryHandlerHttp.ListResources")
	defer span.End()


	var registry *entity.CustomerRegistry
	if err := c.ShouldBindJSON(&registry); err != nil {
		span.RecordError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := obj.Service.Create(ctx, registry)
	if err != nil {
		span.RecordError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found")})
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (obj *CustomerRegistryHandlerHttp) Get(c *gin.Context) {
	ctx, span := obj.Tracer.Tracer.Start(c.Request.Context(), "CustomerRegistryHandlerHttp.ListResources")
	defer span.End()


	var registry *entity.CustomerRegistry
	if err := c.ShouldBindJSON(&registry); err != nil {
		span.RecordError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := obj.Service.Create(ctx, registry)
	if err != nil {
		span.RecordError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found")})
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (obj *CustomerRegistryHandlerHttp) Search(c *gin.Context) {
	ctx, span := obj.Tracer.Tracer.Start(c.Request.Context(), "CustomerRegistryHandlerHttp.ListResources")
	defer span.End()


	var registry *entity.CustomerRegistry
	if err := c.ShouldBindJSON(&registry); err != nil {
		span.RecordError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := obj.Service.Create(ctx, registry)
	if err != nil {
		span.RecordError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found")})
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (obj *CustomerRegistryHandlerHttp) IsValid(c *gin.Context) {
	ctx, span := obj.Tracer.Tracer.Start(c.Request.Context(), "CustomerRegistryHandlerHttp.ListResources")
	defer span.End()


	var registry *entity.CustomerRegistry
	if err := c.ShouldBindJSON(&registry); err != nil {
		span.RecordError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := obj.Service.Create(ctx, registry)
	if err != nil {
		span.RecordError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("not found")})
		return
	}

	c.JSON(http.StatusAccepted, result)
}