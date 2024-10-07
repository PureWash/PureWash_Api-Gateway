package handlers

import (
	pbp "api_gateway/genproto/pure_wash"
	"api_gateway/internal/domain"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// CreateServiceHandler   godoc
// @Router       /api/service [POST]
// @Security     ApiKeyAuth
// @Summary      Service
// @Description  Service
// @Tags         Service
// @Accept       json
// @Produce      json
// @Param        service body domain.ServiceRequest true "Service  Request"
// @Success      200  {object}  domain.Service
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) CreateServiceHandler(ctx *gin.Context) {
	var (
		payload domain.ServiceRequest
		err     error
	)

	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		handleResponse(ctx, h.log, "error is while reading service information by  body ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}

	if payload.Price <= 0 {
		payload.Price = 0
	}

	response, err := h.services.ServiceService().CreateService(ctx, &pbp.ServiceRequest{
		Tariffs:     payload.Traffic,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       float32(payload.Price),
	})
	if err != nil {
		handleResponse(ctx, h.log, "error is while create  service by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "SUCCESSES", http.StatusCreated, gin.H{
		"service": response,
		"message": "Service added successfully",
		"success": true,
	})
}

// UpdateServiceHandler   godoc
// @Router       /api/service/{id} [put]
// @Security     ApiKeyAuth
// @Summary      Update  Service
// @Description  Updates the details of an existing Service .
// @Tags         Service
// @Accept       json
// @Produce      json
// @Param        id path string true "Service  ID"
// @Param        service body domain.ServiceRequest true "Service Type Update Request"
// @Success      200  {object}  domain.Service
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) UpdateServiceHandler(ctx *gin.Context) {
	var (
		payload domain.ServiceRequest
		err     error
		id      string
	)

	id = ctx.Param("id")
	serviceId, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is service_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}

	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		handleResponse(ctx, h.log, "Failed to parse payload body", http.StatusBadRequest, err.Error())
		return
	}
	if payload.Price <= 0 {
		payload.Price = 0
	}

	response, err := h.services.ServiceService().UpdateService(ctx, &pbp.Service{
		Id:          cast.ToString(serviceId),
		Tariffs:     payload.Traffic,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       float32(payload.Price),
	})
	if err != nil {
		handleResponse(ctx, h.log, "Failed to update Service ", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(ctx, h.log, "SUCCESSES", http.StatusOK, gin.H{
		"service": response,
		"message": "Service added successfully",
		"success": true,
	})
}

// DeleteServiceHandler   godoc
// @Router       /api/service/{id} [delete]
// @Security     ApiKeyAuth
// @Summary      Service
// @Description  Service  Delete
// @Tags         Service
// @Accept       json
// @Produce      json
// @Param       id path string true "Service  ID"
// @Success      200  {object}  domain.Response
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) DeleteServiceHandler(ctx *gin.Context) {
	var (
		payload pbp.PrimaryKey
		err     error
		id      string
	)

	id = ctx.Param("id")
	serviceId, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is service_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	payload.Id = cast.ToString(serviceId)
	_, err = h.services.ServiceService().DeleteService(ctx, &payload)
	if err != nil {
		handleResponse(ctx, h.log, "error is while update  service by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, "service success that delete")
}

// GetServiceHandler   godoc
// @Router       /api/service/{id} [GET]
// @Security     ApiKeyAuth
// @Summary      Service
// @Description  Service
// @Tags         Service
// @Accept       json
// @Produce      json
// @Param        id path string true "Service  ID"
// @Success      200  {object}  domain.Service
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) GetServiceHandler(ctx *gin.Context) {
	var (
		id  string
		err error
	)

	id = ctx.Param("id")
	serviceId, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is service_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	response, err := h.services.ServiceService().GetService(ctx, &pbp.PrimaryKey{
		Id: cast.ToString(serviceId),
	})
	if err != nil {
		handleResponse(ctx, h.log, "error is while update  service by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "SUCCESSES", http.StatusOK, response)
}

// GetAllServices godoc
// @Security     ApiKeyAuth
// @Router       /api/services [GET]
// @Summary      Get all services
// @Description  get all services
// @Tags         Service
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  domain.ServicesResponse
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h Handler) GetAllServices(c *gin.Context) {
	var (
		err          error
		defaultPage  = "1"
		defaultLimit = "10"
	)

	page := cast.ToInt(c.DefaultQuery("page", defaultPage))
	limit := cast.ToInt(c.DefaultQuery("limit", defaultLimit))
	response, err := h.services.ServiceService().GetAllService(context.Background(), &pbp.GetListRequest{
		Page:  int64((page - 1) * limit),
		Limit: int64(limit),
	})
	if err != nil {
		handleResponse(c, h.log, "error is while getting all baskets", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, response)

}
