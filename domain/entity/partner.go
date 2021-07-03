package entity

import "errors"

type CoverageArea struct {
	Type        string `json:"type,omitempty" bson:"type"`
	Coordinates [][][][]float64
}

type Address struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Partner struct {
	ID           int64        `json:"id,omitempty" bson:"_id"`
	TradingName  string       `json:"tradingName,omitempty" bson:"tradingName"`
	OwnerName    string       `json:"ownerName,omitempty" bson:"ownerName"`
	Document     string       `json:"document,omitempty" bson:"document"`
	CoverageArea CoverageArea `json:"coverageArea,omitempty" bson:"coverageArea"`
	Address      Address      `json:"address,omitempty" bson:"address"`
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
