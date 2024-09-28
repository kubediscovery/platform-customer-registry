package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kubediscovery/platform-customer-registry/internal/core/entity"
	"github.com/kubediscovery/platform-customer-registry/pkg/kd_utils"
	"github.com/kubediscovery/platform-customer-registry/pkg/otelpkg"
)

type LabDestroyHandlerHttpInterface interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
}

type LabDestroyHandlerHttp struct {
	Service entity.LabDestroyInterface
	Tracer  *otelpkg.OtelPkgInstrument
}

func NewLabDestroyHandlerHttp(svc *entity.LabDestroyInterface, otl *otelpkg.OtelPkgInstrument, routerGroup *gin.RouterGroup, middleware ...func(c *gin.Context)) LabDestroyHandlerHttpInterface {

	lab := &LabDestroyHandlerHttp{
		Service: *svc,
		Tracer:  otl,
	}

	lab.handlers(routerGroup, middleware...)

	return lab
}

func (c *LabDestroyHandlerHttp) handlers(routerGroup *gin.RouterGroup, middleware ...func(c *gin.Context)) {
	middlewareList := make([]gin.HandlerFunc, len(middleware))
	for i, mw := range middleware {
		middlewareList[i] = mw
	}

	routerGroup.POST("/labdestroy", append(middlewareList, c.Create)...)
	routerGroup.GET("/labdestroy/:id", append(middlewareList, c.GetByID)...)
	routerGroup.GET("/labdestroy/search", append(middlewareList, c.GetByFilter)...)
	routerGroup.GET("/labdestroy", append(middlewareList, c.GetAll)...)
	routerGroup.PATCH("/labdestroy/:id", append(middlewareList, c.Update)...)
}

// CreateLabDestroyResponse    godoc
// @Summary     create a new lab destroy
// @Tags        labDestroy
// @Accept       json
// @Produce     json
// @Description create a new lab destroy
// @Success     200 {object} entity.LabDestroyResponse
// @Failure     404 {object} string
// @Failure     500 {object} string
// @Router      /labdestroy [post]
func (obj *LabDestroyHandlerHttp) Create(c *gin.Context) {
	ctx, span := obj.Tracer.Tracer.Start(c.Request.Context(), "LabDestroyHandlerHttp.ListResources")
	defer span.End()

	var lab *entity.LabDestroy
	if err := c.ShouldBindBodyWithJSON(&lab); err != nil {
		span.RecordError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	labResponse, err := entity.NewLabDestroy(lab)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := obj.Service.Create(ctx, labResponse)
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

// GetAllLabDestroyResponse    godoc
// @Summary     get all lab destroy
// @Tags        labDestroy
// @Accept       json
// @Produce     json
// @Description get all lab destroy
// @Success     200 {object} []entity.LabDestroyResponse
// @Failure     404 {object} string
// @Failure     500 {object} string
// @Router      /labdestroy [get]
func (obj *LabDestroyHandlerHttp) GetAll(c *gin.Context) {
	ctx, span := obj.Tracer.Tracer.Start(c.Request.Context(), "LabDestroyHandlerHttp.ListResources")
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
// @Summary     get a lab destroy by ID
// @Tags        labDestroy
// @Accept      json
// @Produce     json
// @Param       id path string true "found"
// @Description get a lab destroy by ID
// @Success     200 {object} entity.LabDestroyResponse
// @Failure     404 {object} string
// @Failure     500 {object} string
// @Router      /labdestroy/{id} [get]
func (obj *LabDestroyHandlerHttp) GetByID(c *gin.Context) {
	_, span := obj.Tracer.Tracer.Start(c.Request.Context(), "LabDestroyHandlerHttp.ListResources")
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

// GetByFilter  godoc
// @Summary     get a lab destroy by email or project
// @Tags        labDestroy
// @Accept      json
// @Produce     json
// @Param       email query string false "email@domain.com"
// @Param       project query string false "project_name"
// @Param       username query string false "username"
// @Param       available query boolean false "string default" default(false)
// @Description get a lab destroy by email or project
// @Success     200 {object} []entity.LabDestroyResponse
// @Failure     404 {object} string
// @Failure     500 {object} string
// @Router      /labdestroy/search [get]
func (obj *LabDestroyHandlerHttp) GetByFilter(c *gin.Context) {
	_, span := obj.Tracer.Tracer.Start(c.Request.Context(), "LabDestroyHandlerHttp.ListResources")
	defer span.End()

	if len(c.Request.URL.Query()) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or project is required"})
		c.Abort()
		return
	}

	if c.Request.URL.Query().Get("email") == "" && c.Request.URL.Query().Get("project") == "" && c.Request.URL.Query().Get("username") == "" && c.Request.URL.Query().Get("available") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or project is required"})
		c.Abort()
		return
	}

	var available *bool

	if c.Request.URL.Query().Get("available") != "" {
		b := kd_utils.ConvertStringToBool(c.Request.URL.Query().Get("available"))
		available = &b
	}

	result, err := obj.Service.GetByFilter(c.Request.Context(), &entity.LabDestroyResponse{
		Avaliable: available,
		LabDestroy: entity.LabDestroy{
			CustomerRegistry: entity.CustomerRegistry{
				UserEmail:   c.Request.URL.Query().Get("email"),
				ProjectName: c.Request.URL.Query().Get("project"),
				UserName:    c.Request.URL.Query().Get("username"),
			},
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (obj *LabDestroyHandlerHttp) Update(c *gin.Context) {
	_, span := obj.Tracer.Tracer.Start(c.Request.Context(), "LabDestroyHandlerHttp.ListResources")
	defer span.End()

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		c.Abort()
		return
	}

	var input struct {
		Available    bool   `json:"available" binding:"required"`
		ErrorMessage string `json:"error_message"`
	}

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	result, err := obj.Service.PATCH(c.Request.Context(), &id, &input.ErrorMessage, &input.Available)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusAccepted, result)
}
