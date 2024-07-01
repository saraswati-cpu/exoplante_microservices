package controllers

import (
	"net/http"

	"exoplant_services/pkg/Api/models"
	"exoplant_services/pkg/services"

	"github.com/gin-gonic/gin"
)

func Exoplanate(group *gin.RouterGroup) {

	group.POST("/exoplanets", AddExoplanet)

}

// @Summary Create a new product
// @Description Create a new product
// @Accept json
// @Tags Product
// @Produce json
// @Param input body models.Exoplanet true "Product information"
// @Success 200 {object} models.Exoplanet
// @Router /exoplanets [POST]
func AddExoplanet(c *gin.Context) {
	var exoplanet models.Exoplanet
	if err := c.ShouldBindJSON(&exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdExoplanet, err := services.AddExoplanet(&exoplanet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdExoplanet)
}
