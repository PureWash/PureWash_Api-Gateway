package handlers

import (
	pbp "api_gateway/genproto/carpet_service"
	"api_gateway/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

// CreateCompanyHandler   godoc
// @Router       /api/company [POST]
// @Security     ApiKeyAuth
// @Summary      Company
// @Description  Company
// @Tags         Company
// @Accept       json
// @Produce      json
// @Param        company body domain.CompanyRequest true "Company  Request"
// @Success      200  {object}  domain.Company
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) CreateCompanyHandler(ctx *gin.Context) {
	var (
		payload domain.CompanyRequest
		err     error
	)

	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		handleResponse(ctx, h.log, "error is while reading company information by  body ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.CompanyService().CreateCompany(ctx, &pbp.CompanyRequest{
		Name:        payload.Name,
		Description: payload.Description,
	})
	if err != nil {
		handleResponse(ctx, h.log, "error is while create  company by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "SUCCESSES", http.StatusCreated, gin.H{
		"company": response,
		"message": "Company added successfully",
		"success": true,
	})
}

// UpdateCompanyHandler   godoc
// @Router       /api/company/{id} [put]
// @Security     ApiKeyAuth
// @Summary      Update  Company
// @Description  Updates the details of an existing Company .
// @Tags         Company
// @Accept       json
// @Produce      json
// @Param        id path string true "Company  ID"
// @Param        company body domain.CompanyRequest true "Company Type Update Request"
// @Success      200  {object}  domain.Company
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) UpdateCompanyHandler(ctx *gin.Context) {
	var (
		payload domain.CompanyRequest
		err     error
		id      string
	)

	id = ctx.Param("id")
	companyId, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is company_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}

	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		handleResponse(ctx, h.log, "Failed to parse payload body", http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.CompanyService().UpdateCompany(ctx, &pbp.Company{
		Id:          cast.ToString(companyId),
		Name:        payload.Name,
		Description: payload.Description,
	})
	if err != nil {
		handleResponse(ctx, h.log, "Failed to update Company ", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(ctx, h.log, "SUCCESSES", http.StatusOK, gin.H{
		"company": response,
		"message": "Company added successfully",
		"success": true,
	})
}

// DeleteCompanyHandler   godoc
// @Router       /api/company/{id} [delete]
// @Security     ApiKeyAuth
// @Summary      Company
// @Description  Company  Delete
// @Tags         Company
// @Accept       json
// @Produce      json
// @Param       id path string true "Company  ID"
// @Success      200  {object}  domain.Response
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) DeleteCompanyHandler(ctx *gin.Context) {
	var (
		payload pbp.PrimaryKey
		err     error
		id      string
	)

	id = ctx.Param("id")
	companyId, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is company_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	payload.Id = cast.ToString(companyId)
	_, err = h.services.CompanyService().DeleteCompany(ctx, &payload)
	if err != nil {
		handleResponse(ctx, h.log, "error is while update  company by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, "company success that delete")
}

// GetCompanyHandler   godoc
// @Router       /api/company/{id} [GET]
// @Security     ApiKeyAuth
// @Summary      Company
// @Description  Company
// @Tags         Company
// @Accept       json
// @Produce      json
// @Param        id path string true "Company  ID"
// @Success      200  {object}  domain.Company
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) GetCompanyHandler(ctx *gin.Context) {
	var (
		id  string
		err error
	)

	id = ctx.Param("id")
	companyId, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is company_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	response, err := h.services.CompanyService().GetCompany(ctx, &pbp.PrimaryKey{
		Id: cast.ToString(companyId),
	})
	if err != nil {
		handleResponse(ctx, h.log, "error is while update  company by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "SUCCESSES", http.StatusOK, response)
}
