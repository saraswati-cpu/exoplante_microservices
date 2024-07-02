package controllers

import (
	"net/http"
	"strconv"

	"exoplant_services/pkg/Api/models"
	"exoplant_services/pkg/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func SetLogger(l *logrus.Logger) {
	logger = l
}

func Exoplanate(group *gin.RouterGroup) {

	group.POST("/exoplanets", AddExoplanet)
	group.GET("/exoplanets", GetAllExoplanet)
	group.GET("/exoplanets/:id", GetExoplanetByID)
	group.PUT("/exoplanete/:id", UpdateExoplanet)
	group.DELETE("/exoplanete/:id", DeleteExoplanet)

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
	logger.Info("Exoplanete retrieved successfully")
	c.JSON(http.StatusOK, exop)

}

// @Summary Get Explonets by ID
// @Produce json
// @Tags Exoplanets
// @Param id path int true "Exoplanet ID"
// @Success 200 {object} models.Exoplanet
// @Failure 404 {string} string "error"
// @Router /exoplanets/{id} [GET]
func GetExoplanetByID(c *gin.Context) {
	// Get the ID from the URL parameters
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	// Call the service layer
	exoplanet, status, err := services.GetExoplanetByID(id)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	// Return the exoplanet data as JSON
	c.JSON(http.StatusOK, exoplanet)
}

// @Summary Update Explonets by ID
// @Produce json
// @Tags Exoplanets
// @Param id path int true "Exoplanet ID"
// @Success 200 {object} models.Exoplanet
// @Failure 404 {string} string "error"
// @Router /exoplanets/{id} [PUT]
func UpdateExoplanet(c *gin.Context) {
	var exoplanet models.Exoplanet

	// Bind the JSON body to the exoplanet model
	if err := c.ShouldBindJSON(&exoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service layer to update the exoplanet
	status, err := services.UpdateExoplanet(&exoplanet)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exoplanet updated successfully"})
}

// DeleteExoplanet handles the HTTP DELETE request to delete an exoplanet by ID
// @Summary Delete Explonets by ID
// @Router /exoplanets/{id} [DELETE]
func DeleteExoplanet(c *gin.Context) {
	// Extract exoplanet ID from URL parameters
	idParam := c.Param("id")
	exoplanetID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exoplanet ID"})
		return
	}

	// Call the service layer to delete the exoplanet
	status, err := services.DeleteExoplanet(exoplanetID)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message if delete was successful
	c.JSON(http.StatusOK, gin.H{"message": "Exoplanet deleted successfully"})
}
