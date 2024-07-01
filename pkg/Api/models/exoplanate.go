package models

import "exoplant_services/config"

type Exoplanet struct {
	ID          string  `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Distance    int     `json:"distance" binding:"required"` // distance from Earth in light years
	Radius      float64 `json:"radius" binding:"required"`   // radius in Earth-radius units
	Mass        float64 `json:"mass,omitempty"`              // mass in Earth-mass units, only for Terrestrial
	Type        string  `json:"type" binding:"required"`     // GasGiant or Terrestrial
}

var schemaName = config.NewConfig().SchemaName

// define custom table for AddProduct model
func (*Exoplanet) TableName() string {
	return schemaName + ".exoplanet"
}
