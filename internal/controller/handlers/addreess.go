package handlers

import (
	pbp "api_gateway/genproto/carpet_service"
	"api_gateway/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// CreateAddressHandler   godoc
// @Router       /api/address [POST]
// @Security     ApiKeyAuth
// @Summary      Address
// @Description  Address
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param        address body domain.AddressRequest true "Address  Request"
// @Success      200  {object}  domain.Address
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) CreateAddressHandler(ctx *gin.Context) {
	var (
		payload domain.AddressRequest
		err     error
	)

	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		handleResponse(ctx, h.log, "error is while reading user information by  body ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	id, err := ParseUuId(payload.UserID, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is user_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.AddressService().CreateAddress(ctx, &pbp.AddressRequest{
		UserId:    cast.ToString(id),
		Latitude:  payload.Latitude,
		Longitude: payload.Longitude,
	})
	if err != nil {
		handleResponse(ctx, h.log, "error is while create  address by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "SUCCESSES", http.StatusCreated, gin.H{
		"address": response,
		"message": "Address added successfully",
		"success": true,
	})
}

// UpdateAddressHandler   godoc
// @Router       /api/address/{id} [put]
// @Security     ApiKeyAuth
// @Summary      Update  Address
// @Description  Updates the details of an existing Address .
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param        id path string true "Address  ID"
// @Param        address body domain.AddressRequest true "Address Type Update Request"
// @Success      200  {object}  domain.Address
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) UpdateAddressHandler(ctx *gin.Context) {
	var (
		payload domain.AddressRequest
		err     error
		id      string
	)

	id = ctx.Param("id")
	addressId, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is address_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}

	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		handleResponse(ctx, h.log, "Failed to parse payload body", http.StatusBadRequest, err.Error())
		return
	}

	userId, err := ParseUuId(payload.UserID, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is user_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.AddressService().UpdateAddress(ctx, &pbp.Address{
		Id:        cast.ToString(addressId),
		UserId:    cast.ToString(userId),
		Longitude: payload.Longitude,
		Latitude:  payload.Latitude,
	})
	if err != nil {
		handleResponse(ctx, h.log, "Failed to update Address ", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(ctx, h.log, "SUCCESSES", http.StatusOK, gin.H{
		"address": response,
		"message": "Address added successfully",
		"success": true,
	})
}

// DeleteAddressHandler   godoc
// @Router       /api/address/{id} [delete]
// @Security     ApiKeyAuth
// @Summary      Address
// @Description  Address  Delete
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param       id path string true "Address  ID"
// @Success      200  {object}  domain.Response
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) DeleteAddressHandler(ctx *gin.Context) {
	var (
		payload pbp.PrimaryKey
		err     error
		id      string
	)

	id = ctx.Param("id")
	addressId, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is address_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	payload.Id = cast.ToString(addressId)
	_, err = h.services.AddressService().DeleteAddress(ctx, &payload)
	if err != nil {
		handleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "SUCCESSS", http.StatusOK, "address success that delete")
}

//
//// GetAllAddressesHandler   godoc
//// @Router       /api/address [GET]
//// @Summary      Address
//// @Description  Address
//// @Tags         Address
//// @Accept       json
//// @Produce      json
//// @Param        page query string false "page"
//// @Param        limit query string false "limit"
//// @Param        search query string false "search"
//// @Success      200  {object}  domain.AddressesResponse
//// @Failure      400  {object}  domain.Response
//// @Failure      404  {object}  domain.Response
//// @Failure      500  {object}  domain.Response
//// @Security BearerAuth
//func (h *Handler) GetAllAddressesHandler(ctx *gin.Context) {
//	var (
//		err          error
//		defaultPage  = "1"
//		defaultLimit = "10"
//	)
//
//	page := cast.ToInt(ctx.DefaultQuery("page", defaultPage))
//	limit := cast.ToInt(ctx.DefaultQuery("limit", defaultLimit))
//	search := fmt.Sprintf("%%%s%%", ctx.DefaultQuery("search", ""))
//	response, err := h.services.AddressService().GetAll(context.Background(), &domain.GetListRequest{
//		Page: int32((page - 1) * limit),
//		Limit:  int32(limit),
//		Search: search,
//	})
//	if err != nil {
//		handleResponse(ctx, h.log, "error is while getting all baskets", http.StatusInternalServerError, err)
//		return
//	}
//
//	handleResponse(ctx, h.log, "", http.StatusOK, response)
//
//}

// GetAddressHandler   godoc
// @Router       /api/address/{id} [GET]
// @Security     ApiKeyAuth
// @Summary      Address
// @Description  Address
// @Tags         Address
// @Accept       json
// @Produce      json
// @Param        id path string true "Address  ID"
// @Success      200  {object}  domain.Address
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) GetAddressHandler(ctx *gin.Context) {
	var (
		id  string
		err error
	)

	id = ctx.Param("id")
	addressId, err := ParseUuId(id, h.log)
	if err != nil {
		handleResponse(ctx, h.log, "error is while parse to uuid  id that is address_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
		return
	}
	response, err := h.services.AddressService().GetAddress(ctx, &pbp.PrimaryKey{
		Id: cast.ToString(addressId),
	})
	if err != nil {
		handleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(ctx, h.log, "SUCCESSES", http.StatusOK, response)
}

// // User Ning malumotlari

// // CreateAddressForUserHandler   godoc
// // @Router       /api/user_address [POST]
// // @Security     ApiKeyAuth
// // @Summary      User_Address
// // @Description  User_Address
// // @Tags         User_Address
// // @Accept       json
// // @Produce      json
// // @Param        address body domain.AddressForUserRequest true "Address  Request"
// // @Success      200  {object}  map[string]string
// // @Failure      400  {object}  domain.Response
// // @Failure      404  {object}  domain.Response
// // @Failure      500  {object}  domain.Response
// func (h *Handler) CreateAddressForUserHandler(ctx *gin.Context) {
// 	var (
// 		payload domain.AddressForUserRequest
// 		err     error
// 	)

// 	err = ctx.ShouldBindJSON(&payload)
// 	if err != nil {
// 		handleResponse(ctx, h.log, "error is while reading user information by  body ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	userId, err := token.GetUserIdByClaims(ctx, h.log)
// 	if err != nil {
// 		handleResponse(ctx, h.log, "error is while reading userId from authorization ---~~~~~~~ERROR===", http.StatusUnauthorized, err)
// 		return
// 	}

// 	response, err := h.services.AddressService().CreateAddress(ctx, &pbp.AddressRequest{
// 		UserId:    cast.ToString(userId),
// 		Latitude:  payload.Latitude,
// 		Longitude: payload.Longitude,
// 	})
// 	if err != nil {
// 		handleResponse(ctx, h.log, "error is while create  address by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	handleResponse(ctx, h.log, "SUCCESSES", http.StatusCreated, gin.H{
// 		"user": cast.ToString(userId),
// 		"Address": gin.H{
// 			"address": response,
// 			"message": "Address added successfully",
// 			"success": true,
// 		},
// 		"message": "User of Address added successfully",
// 		"success": true,
// 	})
// }

// // UpdateAddressForUserHandler   godoc
// // @Router       /api/user_address/{id} [put]
// // @Security     ApiKeyAuth
// // @Summary      Update  User_Address
// // @Description  Updates the details of an existing User_Address .
// // @Tags         User_Address
// // @Accept       json
// // @Produce      json
// // @Param        id path string true "User_Address  ID"
// // @Param        address body domain.AddressForUserRequest true "Address Type Update Request"
// // @Success      200  {object}  map[string]string
// // @Failure      400  {object}  domain.Response
// // @Failure      404  {object}  domain.Response
// // @Failure      500  {object}  domain.Response
// func (h *Handler) UpdateAddressForUserHandler(ctx *gin.Context) {
// 	var (
// 		payload domain.AddressForUserRequest
// 		err     error
// 		id      string
// 	)

// 	id = ctx.Param("id")
// 	addressId, err := ParseUuId(id, h.log)
// 	if err != nil {
// 		handleResponse(ctx, h.log, "error is while parse to uuid  id that is address_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	err = ctx.ShouldBindJSON(&payload)
// 	if err != nil {
// 		handleResponse(ctx, h.log, "Failed to parse payload body", http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	userId, err := token.GetUserIdByClaims(ctx, h.log)
// 	if err != nil {
// 		handleResponse(ctx, h.log, "error is while reading userId from authorization ---~~~~~~~ERROR===", http.StatusUnauthorized, err.Error())
// 		return
// 	}

// 	response, err := h.services.AddressService().UpdateAddress(ctx, &pbp.Address{
// 		Id:        cast.ToString(addressId),
// 		UserId:    cast.ToString(userId),
// 		Longitude: payload.Longitude,
// 		Latitude:  payload.Latitude,
// 	})
// 	if err != nil {
// 		handleResponse(ctx, h.log, "Failed to update Address ", http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	handleResponse(ctx, h.log, "SUCCESSES", http.StatusOK, gin.H{
// 		"user": cast.ToString(userId),
// 		"Address": gin.H{
// 			"address": response,
// 			"message": "Address updated successfully",
// 			"success": true,
// 		},
// 		"message": "User of Address updated successfully",
// 		"success": true,
// 	})
// }

// // GetAddressForUserHandler   godoc
// // @Router       /api/user_address/{id} [GET]
// // @Security     ApiKeyAuth
// // @Summary      User_Address
// // @Description  User_Address
// // @Tags         User_Address
// // @Accept       json
// // @Produce      json
// // @Param        id path string true "User_Address  ID"
// // @Success      200  {object}  map[string]string
// // @Failure      400  {object}  domain.Response
// // @Failure      404  {object}  domain.Response
// // @Failure      500  {object}  domain.Response
// func (h *Handler) GetAddressForUserHandler(ctx *gin.Context) {
// 	var (
// 		id  string
// 		err error
// 	)

// 	id = ctx.Param("id")
// 	addressId, err := ParseUuId(id, h.log)
// 	if err != nil {
// 		handleResponse(ctx, h.log, "error is while parse to uuid  id that is address_id ---~~~~~~~ERROR===", http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	response, err := h.services.AddressService().GetAddress(ctx, &pbp.PrimaryKey{
// 		Id: cast.ToString(addressId),
// 	})
// 	if err != nil {
// 		handleResponse(ctx, h.log, "error is while update  user by  storage ---~~~~~~~ERROR===", http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	userId, err := token.GetUserIdByClaims(ctx, h.log)
// 	if err != nil {
// 		handleResponse(ctx, h.log, "error is while reading userId from authorization ---~~~~~~~ERROR===", http.StatusUnauthorized, err.Error())
// 		return
// 	}

// 	if response.GetUserId() != cast.ToString(userId) {
// 		handleResponse(ctx, h.log, "error is while get userId  not found in user", http.StatusNotFound, fmt.Errorf("user not foun in authoration"))

// 	}
// 	handleResponse(ctx, h.log, "SUCCESSES", http.StatusOK, gin.H{
// 		"user": cast.ToString(userId),
// 		"Address": gin.H{
// 			"address": response,
// 			"message": "Address get successfully",
// 			"success": true,
// 		},
// 		"message": "User of Address updated successfully",
// 		"success": true,
// 	})
// }
