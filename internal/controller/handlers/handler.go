package handlers

import (
	"api_gateway/internal/domain"
	"api_gateway/internal/pkg/logger"
	"github.com/gin-gonic/gin"
)

func handleResponse(c *gin.Context, log logger.ILogger, msg string, statusCode int, data interface{}) {

	resp := domain.Response{}

	switch code := statusCode; {
	case code < 400:
		resp.Description = "OK"
		log.Info("~~~~> OK", logger.String("msg", msg), logger.Any("status", code))

	case code == 401:
		resp.Description = "Unauthorized"
		log.Error("???? Unauthorized", logger.String("msg", msg), logger.Any("status", code))

	case code < 500:
		resp.Description = "Bad Request"
		log.Error("!!!!! BAD REQUEST", logger.String("msg", msg), logger.Any("status", code))
	default:
		resp.Description = "Internal Server Error"
		log.Error("!!!!! INTERNAL SERVER ERROR", logger.String("msg", msg), logger.Any("status", code), logger.Any("error", data))
	}

	resp.StatusCode = statusCode
	resp.Data = data

	c.JSON(resp.StatusCode, resp)

}
