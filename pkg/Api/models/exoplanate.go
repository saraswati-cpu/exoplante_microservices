package models

import "exoplant_services/config"

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	ID          string        `json:"id"`
	Name        string        `json:"name" binding:"required"`
	Description string        `json:"description" binding:"required"`
	Distance    int           `json:"distance" binding:"required"` // distance from Earth in light years
	Radius      float64       `json:"radius" binding:"required"`   // radius in Earth-radius units
	Mass        float64       `json:"mass,omitempty"`              // mass in Earth-mass units, only for Terrestrial
	Type        ExoplanetType `json:"type"`                        // GasGiant or Terrestrial
}

var schemaName = config.NewConfig().SchemaName

// define custom table for AddProduct model
func (*Exoplanet) TableName() string {
	return schemaName + ".exoplanet"
}

type OutputResponse struct {
	Products []Exoplanet `json:"exoplanet"`
}
