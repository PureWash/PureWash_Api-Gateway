package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping   godoc
// @Router       /api/ping [get]
// @Security     ApiKeyAuth
// @Summary      Ping
// @Description  Ping
// @Tags         Ping
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  domain.Response
// @Failure      404  {object}  domain.Response
// @Failure      500  {object}  domain.Response
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"Success": "Hi PING PONG",
	})
}
