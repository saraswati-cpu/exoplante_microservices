package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckRoutes(group *gin.RouterGroup) {

	group.GET("/ping", Ping)
}

// @Summary ping
// @Description This end point respond to pings
// @Produce json
// @Tags ping
// @Success 200
// @Failure 404 {string} string "error"
// @Router /ping [GET]
func Ping(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "okay",
	})
}
