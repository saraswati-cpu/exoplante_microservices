package database

import (
	"errors"
	"exoplant_services/config"
	"exoplant_services/pkg/Api/models"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var logger *logrus.Logger

func SetLogger(l *logrus.Logger) {
	logger = l
}

func AddExoplanet(exopl *models.Exoplanet) (int, error) {
	exoplanet := *exopl
	result := config.DB.Create(&exoplanet)
	if result.Error != nil {

		if strings.Contains(result.Error.Error(), "unique constraint") {

			return http.StatusBadRequest, errors.New("exoplanete with the same name already exists")
		}
		return http.StatusInternalServerError, errors.New("error adding exoplanete in database")
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

func GetExoplanetByID(id int) (models.Exoplanet, int, error) {
	var exop models.Exoplanet
	result := config.DB.Select("name, description, distance, radius, mass, type").Where("id = ?", id).First(&exop)
	if result.Error != nil {
		return exop, http.StatusNotFound, errors.New("failed to get exoplanete for this exoplanet ID")
	}
	return exop, http.StatusOK, nil

}

func UpdateExoplanet(exop *models.Exoplanet) (int, error) {
	result := config.DB.Save(&exop)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "unique constraint") {
			logger.Error(result.Error.Error())
			return http.StatusBadRequest, errors.New("exoplanet name with same name already exists")
		}
		logger.Error("Exoplanet name with same name already exists", result.Error)
		return http.StatusBadRequest, errors.New("exoplanete name with same name already exists")
	}
	return http.StatusOK, nil
}

func DeleteExoplanet(id int) (int, error) {
	// Check if the exoplanet exists
	var exoplanet models.Exoplanet
	result := config.DB.Where("id = ?", id).First(&exoplanet)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, errors.New("exoplanet not found")
		}
		return http.StatusInternalServerError, errors.New("database error")
	}

	// Delete the exoplanet
	result = config.DB.Delete(&exoplanet)
	if result.Error != nil {
		return http.StatusInternalServerError, errors.New("failed to delete exoplanet")
	}

	return http.StatusOK, nil
}
