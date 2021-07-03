package viewmodel

import (
	"errors"

	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
)

type PartnerRequest struct {
	TradingName  string              `json:"tradingName,omitempty"`
	OwnerName    string              `json:"ownerName,omitempty"`
	Document     string              `json:"document,omitempty"`
	CoverageArea entity.MultiPolygon `json:"coverageArea,omitempty"`
	Address      entity.Point        `json:"address,omitempty"`
}

type PartnerResponse struct {
	ID           string              `json:"id,omitempty" bson:"_id"`
	TradingName  string              `json:"tradingName,omitempty" bson:"tradingName"`
	OwnerName    string              `json:"ownerName,omitempty" bson:"ownerName"`
	Document     string              `json:"document,omitempty" bson:"document"`
	CoverageArea entity.MultiPolygon `json:"coverageArea,omitempty" bson:"coverageArea"`
	Address      entity.Point        `json:"address,omitempty" bson:"address"`
}

func NewPartner(viewmodel PartnerRequest) (partner entity.Partner, err error) {

	defer func() {
		if recover() != nil {
			err = errors.New("fail to create new partner")
		}
	}()

	partner.TradingName = viewmodel.TradingName
	partner.OwnerName = viewmodel.OwnerName
	partner.Document = viewmodel.Document
	partner.CoverageArea = viewmodel.CoverageArea
	partner.Address = viewmodel.Address

	return partner, nil
}

func ModelsToView(partners []entity.PartnerDoc) (partnersView []PartnerResponse, err error) {

	defer func() {
		if recover() != nil {
			err = errors.New("fail to parse partner model to view")
		}
	}()

	for _, value := range partners {
		response := PartnerResponse{
			ID:           value.ID,
			TradingName:  value.TradingName,
			OwnerName:    value.OwnerName,
			Document:     value.Document,
			CoverageArea: value.CoverageArea,
			Address:      value.Address,
		}
		partnersView = append(partnersView, response)
	}

	return partnersView, nil
}
