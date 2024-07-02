package services

import (
	"errors"
	"exoplant_services/pkg/Api/models"
	"exoplant_services/pkg/database"
	"net/http"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func SetLogger(l *logrus.Logger) {
	logger = l
}

var exoplanets = make(map[string]models.Exoplanet)

func CalculateGravity(exoplanet models.Exoplanet) float64 {
	if exoplanet.Type == models.GasGiant {
		return 0.5 / (exoplanet.Radius * exoplanet.Radius)
	}
	return exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)
}

func CalculateFuel(exoplanet models.Exoplanet, crewCapacity int) (float64, error) {
	if crewCapacity <= 0 {
		return 0, errors.New("invalid crew capacity")
	}
	gravity := CalculateGravity(exoplanet)
	return float64(exoplanet.Distance) / (gravity * gravity) * float64(crewCapacity), nil
}

func AddExoplanet(exoplanet *models.Exoplanet) (int, error) {
	if exoplanet.Distance < 10 || exoplanet.Distance > 1000 {
		return http.StatusBadRequest, errors.New("distance must be between 10 and 1000 light years")
	}
	if exoplanet.Radius < 0.1 || exoplanet.Radius > 10 {
		return http.StatusBadRequest, errors.New("radius must be between 0.1 and 10 Earth-radius units")
	}
	if exoplanet.Type == "Terrestrial" && (exoplanet.Mass < 0.1 || exoplanet.Mass > 10) {
		return http.StatusBadRequest, errors.New("mass must be between 0.1 and 10 Earth-mass units for Terrestrial planets")
	}

	statusCode, err := database.AddExoplanet(exoplanet)
	if err != nil {
		return statusCode, err
	}
	return http.StatusCreated, nil
}

func GetAllExoplanet() (models.OutputResponse, int, error) {

	var output models.OutputResponse

	products, statusCode, err := database.GetAllExoplanet()
	if err != nil {
		logger.Error("Error retrieving products from database:", err)
		return output, statusCode, err
	}

	output.Products = append(output.Products, products...)

	return output, statusCode, nil
}

func GetExoplanetByID(id int) (models.Exoplanet, int, error) {
	exoplanet, status, err := database.GetExoplanetByID(id)
	if err != nil {
		return exoplanet, status, err
	}

	return exoplanet, http.StatusOK, nil
}

func UpdateExoplanet(product *models.Exoplanet) (int, error) {
	// Validate the input fields
	if product.Name == "" {
		return http.StatusBadRequest, errors.New("exoplanet name is required")
	}
	if product.Description == "" {
		return http.StatusBadRequest, errors.New("exoplanet description is required")
	}
	if product.Distance <= 0 {
		return http.StatusBadRequest, errors.New("exoplanet distance must be a positive integer")
	}
	if product.Radius <= 0 {
		return http.StatusBadRequest, errors.New("exoplanet radius must be a positive number")
	}
	if product.Type == "" {
		return http.StatusBadRequest, errors.New("exoplanet type is required")
	}
	if product.Type == models.Terrestrial && product.Mass <= 0 {
		return http.StatusBadRequest, errors.New("exoplanet mass is required for terrestrial type and must be a positive number")
	}

	// Call the data layer to update the exoplanet
	status, err := database.UpdateExoplanet(product)
	if err != nil {
		// Additional logging or error handling if needed
		logger.Error("Failed to update exoplanet: ", err)
		return status, err
	}

	return http.StatusOK, nil
}

// DeleteExoplanet deletes an exoplanet by its ID
func DeleteExoplanet(id int) (int, error) {
	// Call the data layer to delete the exoplanet
	status, err := database.DeleteExoplanet(id)
	if err != nil {
		return status, err
	}

	return http.StatusOK, nil
}
