package models

import "github.com/jinzhu/gorm"

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

// models structure for exoplanets
type Exoplanet struct {
	gorm.Model
	Name        string        `json:"name" binding:"required"`
	Description string        `json:"description" binding:"required"`
	Distance    int           `json:"distance" binding:"required,gte=10,lte=1000"`
	Radius      float64       `json:"radius" binding:"required,gte=0.1,lte=10"`
	Mass        float64       `json:"mass,omitempty" binding:"omitempty,gte=0.1,lte=10"`
	Type        ExoplanetType `json:"type" binding:"required,oneof=GasGiant Terrestrial"`
}
