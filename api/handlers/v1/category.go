package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	pb "github.com/muhriddinsalohiddin/online_store_api/genproto/catalog_service"
	l "github.com/muhriddinsalohiddin/online_store_api/pkg/logger"
	"github.com/muhriddinsalohiddin/online_store_api/pkg/utils"
)

// CreateCategory ...
// @Summary CreateCategory
// @Description This API for creating a new category
// @Tags category
// @Accept  json
// @Produce  json
// @Param category request body models.Category true "categoryCreateRequest"
// @Success 200 {object} models.Category
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/categories/ [post]
func (h *handlerV1) CreateCategory(c *gin.Context) {
	var (
		body        pb.Category
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	resp, err := h.serviceManager.CatalogService().CreateCategory(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create category", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// UpdateCategory ...
// @Summary UpdateCategory
// @Description This API for updating category
// @Tags category
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Category request body models.Category true "categoryUpdateRequest"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/categories/{id} [put]
func (h *handlerV1) UpdateCategory(c *gin.Context) {
	var (
		body        pb.Category
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	resp, err := h.serviceManager.CatalogService().UpdateCategory(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update category", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetCategoryById ...
// @Summary GetCategoryById
// @Description This API for getting category detail
// @Tags category
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Category
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/categories/{id} [get]
func (h *handlerV1) GetCategoryById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	resp, err := h.serviceManager.CatalogService().GetCategoryById(ctx, &pb.GetCategoryByIdReq{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get category", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// DeleteCategoryById ...
// @Summary DeleteCategoryById
// @Description This API for deleting category
// @Tags category
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/categories/{id} [delete]
func (h *handlerV1) DeleteCategoryById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	resp, err := h.serviceManager.CatalogService().DeleteCategoryById(ctx, &pb.GetCategoryByIdReq{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete category", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// ListCategories ...
// @Summary ListCategories
// @Description This API for getting list of categories
// @Tags category
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.ListCategories
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/categories [get]
func (h *handlerV1) ListCategories(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse qurey params json" + errStr[0])
	}
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	resp, err := h.serviceManager.CatalogService().ListCategories(ctx,
		&pb.ListCategoryReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list category", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}
