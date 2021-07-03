package entity

import "errors"

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
	ID           string       `json:"id,omitempty" bson:"_id"`
	TradingName  string       `json:"tradingName,omitempty" bson:"tradingName"`
	OwnerName    string       `json:"ownerName,omitempty" bson:"ownerName"`
	Document     string       `json:"document,omitempty" bson:"document"`
	CoverageArea MultiPolygon `json:"coverageArea,omitempty" bson:"coverageArea"`
	Address      Point        `json:"address,omitempty" bson:"address"`
}

func (p *Partner) Validate() (err error) {

	if p.TradingName == "" {
		return errors.New("tradingName: required field")
	}

	if p.OwnerName == "" {
		return errors.New("ownerName: required field")
	}

	if p.Document == "" {
		return errors.New("document: required field")
	}

	if p.CoverageArea.Type == "" {
		return errors.New("coverageArea.type: required field")
	}

	if p.CoverageArea.Coordinates == nil {
		return errors.New("coverageArea.coordinates: required field")
	}

	return nil
}
