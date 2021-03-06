package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/muhriddinsalohiddin/online_store_api/api/handlers/models"
	pb "github.com/muhriddinsalohiddin/online_store_api/genproto/order_service"
	l "github.com/muhriddinsalohiddin/online_store_api/pkg/logger"
	"github.com/muhriddinsalohiddin/online_store_api/pkg/utils"
)

// CreateOrder ...
// @Summary CreateOrder
// @Description This API for creating a new order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param Order request body models.Order true "orderCreateRequest"
// @Success 200 {object} models.Order
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orders/ [post]
func (h *handlerV1) CreateOrder(c *gin.Context) {
	var (
		body        pb.Order
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
	response, err := h.serviceManager.OrderService().CreateOrder(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create order", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// GetOrder ...
// @Summary GetOrder
// @Description This API for getting order detail
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Order
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orders/{id} [get]
func (h *handlerV1) GetOrderById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.OrderService().GetOrderById(
		ctx, &pb.GetOrderByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get order", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateOrder ...
// @Summary UpdateOrder
// @Description This API for updating order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Order request body models.Order true "OrderUpdateRequest"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orders/{id} [put]
func (h *handlerV1) UpdateOrder(c *gin.Context) {
	var (
		body        pb.Order
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

	response, err := h.serviceManager.OrderService().UpdateOrder(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update order", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteOrder ...
// @Summary DeleteOrder
// @Description This API for deleting Order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orders/{id} [delete]
func (h *handlerV1) DeleteOrder(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.OrderService().DeleteById(
		ctx, &pb.GetOrderByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete Order", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListOrders ...
// @Summary ListOrders
// @Description This API for getting list of Orders
// @Tags Order
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.ListOrders
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orders [get]
func (h *handlerV1) ListOrders(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.OrderService().ListOrders(
		ctx, &pb.ListOrderReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list Orders", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
