package controllers

import (
	"net/http"

	"exoplant_services/pkg/Api/models"
	"exoplant_services/pkg/services"

	"github.com/gin-gonic/gin"
)

func Exoplanate(group *gin.RouterGroup) {

	group.POST("/exoplanets", AddExoplanet)
	group.GET("/exoplanets", GetAllExoplanet)
	// group.GET("/exoplanete/:id", GetExoplanetById)
	// group.PUT("/exoplanete/:id", EditExoplanet)
	// group.DELETE("/exoplanete/:id", DeleteExoplanet)

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

// @Summary Get All exoplanete
// @Description This end point returns list of all exoplanetts
// @Produce json
// @Tags Exoplanets
// @Success 200 {object} models.Exoplanet
// @Failure 404 {string} string "error"
// @Router /exoplanets [GET]
func GetAllExoplanet(c *gin.Context) {
	exop, statusCode, err := services.GetAllExoplanet() // calling services Layer
	if err != nil {
		switch statusCode {
		case http.StatusBadRequest:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		case http.StatusNotFound:
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	//logger.Info("Exoplanete retrieved successfully")
	c.JSON(http.StatusOK, exop)

}

// // @Summary Get a exoplanete by product ID
// // @Description Get a exoplanete using its ID
// // @Produce json
// // @Tags Product
// // @Param id path int true "exoplanete id"
// // @Success 200 {object} models.ExoplaneteById
// // @Failure 404 {string} string "error"
// // @Router /exoplanete/:id[GET]
// func GetProductVersionById(c *gin.Context) {
// 	id := c.Param("id")
// 	exoplaneteId, err := strconv.Atoi(id)
// 	if err != nil {
// 		logger.Error("Invalid type for exoplante ID:", err)
// 		c.JSON(400, gin.H{"error": "Invalid type for exoplanete ID"})
// 		return
// 	}
// 	prod, statusCode, err := services.GetExoplaneteById(exId)
// 	if err != nil {
// 		switch statusCode {
// 		case http.StatusBadRequest:
// 			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		case http.StatusNotFound:
// 			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
// 		default:
// 			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 		}
// 		return
// 	}
// 	c.JSON(http.StatusOK, prod)
// }
