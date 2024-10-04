package handlers

import (
	pbp "api_gateway/genproto/carpet_service"
	"api_gateway/internal/domain"
	token "api_gateway/internal/pkg/jwt"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// CreateOrderHandler   godoc
// @Router       /api/order [POST]
// @Security     ApiKeyAuth
// @Summary      Order
// @Description  Order
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param        order body domain.OrderRequest true "Order  Request"
// @Success      200  {object}  domain.Order
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) CreateOrderHandler(ctx *gin.Context) {
	var (
		payload domain.OrderRequest
		err     error
	)

	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		handleResponse(ctx, h.log, "error is while reading order information by  body ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	userID, err := ParseUuId(payload.UserID, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parsing to uuid in userID ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}

	serviceID, err := ParseUuId(payload.ServiceID, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parsing to uuid in serviceID ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	if payload.TotalPrice <= 0 {
		payload.TotalPrice = 0
	}
	if payload.Area < 0 {
		handleResponse(ctx, h.log, "error is while you don't you area and you have to area>zero ---~~~~~~~ERROR===", http.StatusBadRequest, "error is while you don't you area and you have to area>zero ")
		return
	}
	response, err := h.services.OrderService().CreateOrder(ctx, &pbp.OrderRequest{
		UserId:     cast.ToString(userID),
		ServiceId:  cast.ToString(serviceID),
		Area:       payload.Area,
		TotalPrice: float32(payload.TotalPrice),
		Status:     payload.Status,
	})
	if err != nil {
		handleResponse(ctx, h.log, "error is while create  order by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "SUCCESSES", http.StatusCreated, gin.H{
		"order":   response,
		"message": "Order added successfully",
		"success": true,
	})
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
// @Param        order body domain.OrderRequest true "Order Type Update Request"
// @Success      200  {object}  domain.Order
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) UpdateOrderHandler(ctx *gin.Context) {
	var (
		payload domain.OrderRequest
		err     error
		id      string
	)

	id = ctx.Param("id")
	orderId, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is order_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}

	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		handleResponse(ctx, h.log, "Failed to parse payload body", http.StatusBadRequest, err.Error())
		return
	}
	userID, err := ParseUuId(payload.UserID, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parsing to uuid in userID ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}

	serviceID, err := ParseUuId(payload.ServiceID, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parsing to uuid in serviceID ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	if payload.TotalPrice <= 0 {
		payload.TotalPrice = 0
	}
	if payload.Area < 0 {
		handleResponse(ctx, h.log, "error is while you don't you area and you have to area>zero ---~~~~~~~ERROR===", http.StatusBadRequest, "error is while you don't you area and you have to area>zero ")
		return
	}

	response, err := h.services.OrderService().UpdateOrder(ctx, &pbp.Order{
		Id:         cast.ToString(orderId),
		UserId:     cast.ToString(userID),
		ServiceId:  cast.ToString(serviceID),
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
// @Success      200  {object}  domain.Order
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
// @Param        search query string false "search"
// @Success      200  {object}  domain.OrdersResponse
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
	search := fmt.Sprintf("%%%s%%", c.DefaultQuery("search", ""))
	response, err := h.services.OrderService().GetAllOrder(context.Background(), &pbp.GetListRequest{
		Page:   int64((page - 1) * limit),
		Limit:  int64(limit),
		Search: search,
	})
	if err != nil {
		handleResponse(c, h.log, "error is while getting all baskets", http.StatusInternalServerError, err.Error())
		return
	}
	var services []*pbp.Service
	for i := 0; i < len(response.Orders); i++ {
		serviceResp, err := h.services.ServiceService().GetService(c, &pbp.PrimaryKey{
			Id: cast.ToString(response.Orders[i].ServiceId),
		})
		if err != nil {
			handleResponse(c, h.log, "error is while getting all services", http.StatusInternalServerError, err.Error())
			return
		}
		services = append(services, serviceResp)

	}

	handleResponse(c, h.log, "SUCCESSES", http.StatusOK, gin.H{
		"Order": gin.H{
			"order":   response,
			"message": "Order get-all successfully",
			"success": true,
		},
		"services": gin.H{
			"services": services,
			"message":  "services get all successfully",
			"success":  true,
		},
		"message": "successfully",
		"success": true,
	})
}

// user order information

// CreateOrderForUserHandler   godoc
// @Router       /api/user_order [POST]
// @Security     ApiKeyAuth
// @Summary      User_Order
// @Description  User_Order
// @Tags         User_Order
// @Accept       json
// @Produce      json
// @Param        order body domain.OrderForUserRequest true "User_Order  Request"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) CreateOrderForUserHandler(ctx *gin.Context) {
	var (
		payload domain.OrderForUserRequest
		err     error
	)

	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		handleResponse(ctx, h.log, "error is while reading order information by  body ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}

	serviceID, err := ParseUuId(payload.ServiceID, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parsing to uuid in serviceID ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	if payload.TotalPrice <= 0 {
		payload.TotalPrice = 0
	}
	if payload.Area < 0 {
		handleResponse(ctx, h.log, "error is while you don't you area and you have to area>zero ---~~~~~~~ERROR===", http.StatusBadRequest, "error is while you don't you area and you have to area>zero ")
		return
	}
	userId, err := token.GetUserIdByClaims(ctx, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while reading userId from authorization ---~~~~~~~ERROR===", http.StatusUnauthorized, err.Error())
		return
	}

	response, err := h.services.OrderService().CreateOrder(ctx, &pbp.OrderRequest{
		UserId:     cast.ToString(userId),
		ServiceId:  cast.ToString(serviceID),
		Area:       payload.Area,
		TotalPrice: float32(payload.TotalPrice),
		Status:     payload.Status,
	})
	if err != nil {
		handleResponse(ctx, h.log, "error is while create  order by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "SUCCESSES", http.StatusOK, gin.H{
		"user": cast.ToString(userId),
		"Order": gin.H{
			"order":   response,
			"message": "Order create successfully",
			"success": true,
		},
		"message": "User of order create successfully",
		"success": true,
	})
}

// UpdateOrderForUserHandler   godoc
// @Router       /api/user_order_canceled/{id} [PUT]
// @Security     ApiKeyAuth
// @Summary      Update  User_Order
// @Description  Updates the details of an existing User_Order .
// @Tags         User_Order
// @Accept       json
// @Produce      json
// @Param        id path string true "Order Id canceled Request"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) UpdateOrderForUserHandler(ctx *gin.Context) {
	var (
		err error
		id  string
	)

	id = ctx.Param("id")
	orderId, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is order_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	order, err := h.services.OrderService().GetOrder(ctx, &pbp.PrimaryKey{
		Id: cast.ToString(orderId),
	})
	if err != nil {
		handleResponse(ctx, h.log, "error is while update  order by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}

	userId, err := token.GetUserIdByClaims(ctx, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while reading userId from authorization ---~~~~~~~ERROR===", http.StatusUnauthorized, err.Error())
		return
	}

	response, err := h.services.OrderService().UpdateOrder(ctx, &pbp.Order{
		Id:         order.GetId(),
		UserId:     cast.ToString(userId),
		AddressId:  order.GetAddressId(),
		ServiceId:  order.GetServiceId(),
		Area:       order.GetArea(),
		TotalPrice: order.GetTotalPrice(),
		Status:     "canceled",
	})
	if err != nil {
		handleResponse(ctx, h.log, "Failed to update Order ", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(ctx, h.log, "SUCCESSES", http.StatusOK, gin.H{
		"user": cast.ToString(userId),
		"Order": gin.H{
			"order":   response,
			"message": "Order canceled successfully",
			"success": true,
		},
		"message": "User of order canceled successfully",
		"success": true,
	})

}

// GetOrderForUserHandler   godoc
// @Security     ApiKeyAuth
// @Router       /api/user_order/{id} [GET]
// @Summary      User_Order
// @Description  User_Order
// @Tags         User_Order
// @Accept       json
// @Produce      json
// @Param        id path string true "User_Order  ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) GetOrderForUserHandler(ctx *gin.Context) {
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
	userId, err := token.GetUserIdByClaims(ctx, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while reading userId from authorization ---~~~~~~~ERROR===", http.StatusUnauthorized, err.Error())
		return
	}

	if response.GetUserId() != cast.ToString(userId) {
		handleResponse(ctx, h.log, "error is while get userId  not found in user", http.StatusNotFound, fmt.Errorf("user not found"))
		return

	}
	handleResponse(ctx, h.log, "SUCCESSES", http.StatusOK, gin.H{
		"user": cast.ToString(userId),
		"Order": gin.H{
			"order":   response,
			"message": "Order get successfully",
			"success": true,
		},
		"message": "User of order get successfully",
		"success": true,
	})
}

// GetAllOrdersForUser godoc
// @Security     ApiKeyAuth
// @Router       /api/user_orders [GET]
// @Summary      Get all orders
// @Description  get all orders
// @Tags         User_Order
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h Handler) GetAllOrdersForUser(c *gin.Context) {
	var (
		err          error
		defaultPage  = "1"
		defaultLimit = "10"
	)

	page := cast.ToInt(c.DefaultQuery("page", defaultPage))
	limit := cast.ToInt(c.DefaultQuery("limit", defaultLimit))
	search := fmt.Sprintf("%%%s%%", c.DefaultQuery("search", ""))
	response, err := h.services.OrderService().GetAllOrder(context.Background(), &pbp.GetListRequest{
		Page:   int64((page - 1) * limit),
		Limit:  int64(limit),
		Search: search,
	})
	if err != nil {
		handleResponse(c, h.log, "error is while getting all baskets", http.StatusInternalServerError, err.Error())
		return
	}

	userId, err := token.GetUserIdByClaims(c, h.log)
	if err != nil {
		handleResponse(c, h.log, "error is while reading userId from authorization ---~~~~~~~ERROR===", http.StatusUnauthorized, err.Error())
		return
	}

	var orders []*pbp.Order
	for _, order := range response.Orders {
		if order.GetUserId() == cast.ToString(userId) {
			orders = append(orders, order)
		}
	}

	handleResponse(c, h.log, "Success", http.StatusOK, gin.H{
		"user": cast.ToString(userId),
		"Order": gin.H{
			"orders":  orders,
			"message": "Orders get all successfully",
			"success": true,
		},
		"message": "User get-all successfully",
		"success": true,
	})
}
