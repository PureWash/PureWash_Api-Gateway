package handlers

import (
	pbp "api_gateway/genproto/pure_wash"
	"api_gateway/internal/domain"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// CreateOrderHandler godoc
// @Router       /api/order [POST]
// @Security     ApiKeyAuth
// @Summary      Create an order
// @Description  Endpoint to create a new order
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        order body pure_wash.CreateOrderReq true "Order Request"
// @Success      201  {object}  pure_wash.CreateOrderResp
// @Failure      400  {object}  domain.Response
// @Failure      401  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) CreateOrderHandler(ctx *gin.Context) {
	var (
		payload pbp.CreateOrderReq
		err     error
	)
	if err = ctx.ShouldBindJSON(&payload); err != nil {
		handleResponse(ctx, h.log, "Invalid order data", http.StatusBadRequest, err.Error())
		return
	}

	_, err = ParseUuId(payload.ServiceId, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "Invalid service ID format", http.StatusBadRequest, err.Error())
		return
	}

	if payload.Area <= 0 {
		handleResponse(ctx, h.log, "Area must be greater than zero", http.StatusBadRequest, "Invalid area")
		return
	}

	if payload.TotalPrice < 0 {
		payload.TotalPrice = 0
	}

	response, err := h.services.OrderService().CreateOrder(ctx, &payload)
	if err != nil {
		handleResponse(ctx, h.log, "Failed to create order", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(ctx, h.log, "Order created successfully", http.StatusCreated, response)
}

// UpdateOrderHandler   godoc
// @Router       /api/order/{id} [put]
// @Security     ApiKeyAuth
// @Summary      Update  Order
// @Description  Updates the details of an existing Order .
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        id path string true "Order  ID"
// @Param        order body pure_wash.UpdateOrderReq true "Order Type Update Request"
// @Success      200  {object}  pure_wash.UpdateOrderResp
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) UpdateOrderHandler(ctx *gin.Context) {
	var (
		payload pbp.UpdateOrderReq
		err     error
		id      string
	)

	id = ctx.Param("id")
	ID, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is order_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}

	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		handleResponse(ctx, h.log, "Failed to parse payload body", http.StatusBadRequest, err.Error())
		return
	}

	if payload.TotalPrice <= 0 {
		payload.TotalPrice = 0
	}
	if payload.Area < 0 {
		handleResponse(ctx, h.log, "error is while you don't you area and you have to area>zero ---~~~~~~~ERROR===", http.StatusBadRequest, "error is while you don't you area and you have to area>zero ")
		return
	}

	response, err := h.services.OrderService().UpdateOrder(ctx, &pbp.UpdateOrderReq{
		Id:         cast.ToString(ID),
		Latitude:   cast.ToFloat32(payload.Latitude),
		Longitude:  cast.ToFloat32(payload.Longitude),
		Area:       payload.Area,
		TotalPrice: float32(payload.TotalPrice),
		Status:     payload.Status,
	})
	if err != nil {
		handleResponse(ctx, h.log, "Failed to update Order ", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(ctx, h.log, "SUCCESSES", http.StatusOK, gin.H{
		"order":   response,
		"message": "Order added successfully",
		"success": true,
	})
}

// DeleteOrderHandler   godoc
// @Router       /api/order/{id} [delete]
// @Security     ApiKeyAuth
// @Summary      Order
// @Description  Order  Delete
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param       id path string true "Order  ID"
// @Success      200  {object}  domain.Response
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) DeleteOrderHandler(ctx *gin.Context) {
	var (
		payload pbp.PrimaryKey
		err     error
		id      string
	)

	id = ctx.Param("id")
	orderId, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is order_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	payload.Id = cast.ToString(orderId)
	_, err = h.services.OrderService().DeleteOrder(ctx, &payload)
	if err != nil {
		handleResponse(ctx, h.log, "error is while update  order by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, "order success that delete")
}

// GetOrderHandler   godoc
// @Router       /api/order/{id} [GET]
// @Security     ApiKeyAuth
// @Summary      Order
// @Description  Order
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        id path string true "Order  ID"
// @Success      200  {object}  pure_wash.GetOrderResp
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) GetOrderHandler(ctx *gin.Context) {
	var (
		id  string
		err error
	)

	id = ctx.Param("id")
	orderId, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is order_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	response, err := h.services.OrderService().GetOrder(ctx, &pbp.PrimaryKey{
		Id: cast.ToString(orderId),
	})
	if err != nil {
		handleResponse(ctx, h.log, "error is while update  order by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "SUCCESSES", http.StatusOK, response)
}

// GetAllOrderForCouriers godoc
// @Security     ApiKeyAuth
// @Router       /api/courier_orders [GET]
// @Summary      Get all couriers orders
// @Description  get all couriers orders
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Success      200  {object}  pure_wash.GetOrdersResp
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h Handler) GetAllOrderForCouriers(c *gin.Context) {
	var (
		err          error
		defaultPage  = "1"
		defaultLimit = "10"
	)

	page := cast.ToInt(c.DefaultQuery("page", defaultPage))
	limit := cast.ToInt(c.DefaultQuery("limit", defaultLimit))
	//search := fmt.Sprintf("%%%s%%", c.DefaultQuery("search", ""))
	fmt.Println(defaultLimit, defaultPage)
	response, err := h.services.OrderService().GetAllOrder(context.Background(), &pbp.GetListRequest{
		Page:  int64((page - 1) * limit),
		Limit: int64(limit),
		//Search: search,
	})
	if err != nil {
		handleResponse(c, h.log, "error is while getting all baskets", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "SUCCESSES", http.StatusOK, response)
}

// GetAllOrders godoc
// @Security     ApiKeyAuth
// @Router       /api/orders [GET]
// @Summary      Get all orders
// @Description  get all orders
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        status query string false "status"
// @Param        time query string false "time"
// @Param        full_name query string false "full_name"
// @Success      200  {object}  pure_wash.GetOrdersResp
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h Handler) GetAllOrders(c *gin.Context) {
	var (
		err          error
		defaultPage  = "1"
		defaultLimit = "10"
	)

	page := cast.ToInt(c.DefaultQuery("page", defaultPage))
	limit := cast.ToInt(c.DefaultQuery("limit", defaultLimit))
	req := domain.GetAllOrderReq{
		FullName: c.DefaultQuery("full_name", ""),
		Ontime:   c.DefaultQuery("time", ""),
		Status:   c.DefaultQuery("status", ""),
	}

	response, err := h.services.OrderService().GetAllOrderForCurier(context.Background(), &pbp.GetAllOrdersReq{
		Offset:   int32((page - 1) * limit),
		Limit:    int32(limit),
		FullName: req.FullName,
		OnTime:   req.Ontime,
		Status:   req.Status,
	})
	if err != nil {
		handleResponse(c, h.log, "error is while getting all baskets", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "SUCCESSES", http.StatusOK, response)
}

// UpdateOrderStatusHandler   godoc
// @Router       /api/order_status/{id} [put]
// @Security     ApiKeyAuth
// @Summary      Update  Order
// @Description  Updates the details of an existing Order
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        id path string true "Order  ID"
// @Param        status query string true "status"
// @Success      200  {object}  domain.Response
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) UpdateOrderStatusHandler(ctx *gin.Context) {
	var (
		err    error
		id     string
		status string
	)

	id = ctx.Param("id")
	status = ctx.Query("status")
	ID, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is order_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.services.OrderService().UpdateOrderStatus(ctx, &pbp.StatusOrderReq{
		Id:     cast.ToString(ID),
		Status: status,
	})
	if err != nil {
		handleResponse(ctx, h.log, "Failed to update Order ", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "SUCCESSES", http.StatusOK, fmt.Errorf("%s status updated successfully", resp.GetId()))
}
