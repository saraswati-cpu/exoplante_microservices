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
