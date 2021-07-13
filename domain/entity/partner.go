package entity

import (
	"github.com/thiagoluiznunes/ze-challenge/domain"
	"github.com/thiagoluiznunes/ze-challenge/infra/zerrors"
)

//Coordinates is a set of coordinate
type Coordinates []Coordinate

//Coordinate is a [longitude, latitude]
type Coordinate [2]float64

// Point rapresent a geojson point geometry object
type Point struct {
	Type        string     `json:"type"`
	Coordinates Coordinate `json:"coordinates"`
}

// MultiPolygon rapresent a geojson mulitpolygon  geometry object
type MultiPolygon struct {
	Type        string          `json:"type,omitempty"`
	Coordinates [][]Coordinates `json:"coordinates,omitempty"`
}

type Partner struct {
	ID           string       `json:"id,omitempty" bson:"_id,omitempty"`
	TradingName  string       `json:"tradingName,omitempty" bson:"tradingName"`
	OwnerName    string       `json:"ownerName,omitempty" bson:"ownerName"`
	Document     string       `json:"document,omitempty" bson:"document"`
	CoverageArea MultiPolygon `json:"coverageArea,omitempty" bson:"coverageArea"`
	Address      Point        `json:"address,omitempty" bson:"address"`
}

func (a *Point) validate() (err error) {

	if a.Type == "" {
		return zerrors.NewValidationError("address.type", "required field")
	}

	if a.Type != domain.MapGeometryType[domain.Point] {
		return zerrors.NewValidationError("address.type", "invalid type")
	}

	if a.Coordinates[0] == 0 && a.Coordinates[1] == 0 {
		return zerrors.NewValidationError("address.coordinates", "required field")
	}

	if len(a.Coordinates) <= 0 {
		return zerrors.NewValidationError("address.coordinates", "null or empty")
	}

	return nil
}

func (m *MultiPolygon) validate() (err error) {

	if m.Type == "" {
		return zerrors.NewValidationError("coverageArea.type", "required field")
	}

	if m.Type != domain.MapGeometryType[domain.MultiPolygon] {
		return zerrors.NewValidationError("coverageArea.type", "invalid type")
	}

	if len(m.Coordinates) <= 0 {
		return zerrors.NewValidationError("coverageArea.coordinates", "required field")
	} else if len(m.Coordinates[0]) <= 0 {
		return zerrors.NewValidationError("coverageArea.coordinates", "null or empty")
	}

	return nil
}

func (p *Partner) Validate() (err error) {

	if p.TradingName == "" {
		return zerrors.NewValidationError("tradingName", " required field")
	}

	if p.OwnerName == "" {
		return zerrors.NewValidationError("ownerName", " required field")
	}

	if p.Document == "" {
		return zerrors.NewValidationError("document", " required field")
	}

	err = p.CoverageArea.validate()
	if err != nil {
		return err
	}
	err = p.Address.validate()
	if err != nil {
		return err
	}

	return nil
}
