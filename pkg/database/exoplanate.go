package database

import (
	"errors"
	"exoplant_services/config"
	"exoplant_services/pkg/Api/models"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func SetLogger(l *logrus.Logger) {
	logger = l
}

var exoplanets = make(map[string]models.Exoplanet)

//	func AddExoplanet(exoplanet models.Exoplanet) (models.Exoplanet, error) {
//		exoplanet.ID = uuid.New().String()
//		exoplanets[exoplanet.ID] = exoplanet
//		return exoplanet, nil
//	}
func AddExoplanet(exopl *models.Exoplanet) (int, error) {
	exoplanet := *exopl
	result := config.DB.Create(&exoplanet)
	if result.Error != nil {

		if strings.Contains(result.Error.Error(), "unique constraint") {
			//logger.Error("Product with the same name already exists:", result.Error)
			return http.StatusBadRequest, errors.New("product with the same name already exists")
		}
		return http.StatusInternalServerError, errors.New("error adding product in database")
	}
	return http.StatusCreated, nil

}

func GetAllExoplanet() ([]models.Exoplanet, int, error) {
	var exop []models.Exoplanet
	result := config.DB.Find(&exop)
	if result.Error != nil {
		logger.Error("Error retrieving exoplanete from the database:", result.Error)
		return nil, http.StatusInternalServerError, errors.New("failed to retrieve products from the database")
	}
	return exop, http.StatusOK, nil
}
