package viewmodel

import (
	"errors"

	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
)

type PartnerRequest struct {
	Name string `json:"name,omitempty"`
}

type PartnerResponse struct {
	Name string `json:"name,omitempty"`
}

func NewPartner(viewmodel PartnerRequest) (partner entity.Partner, err error) {

	defer func() {
		if recover() != nil {
			err = errors.New("fail to create new partner")
		}
	}()

	partner.Name = viewmodel.Name

	return partner, nil
}

func ModelToView(partners []entity.Partner) (partnersView []PartnerResponse, err error) {

	defer func() {
		if recover() != nil {
			err = errors.New("fail to parse partner model to view")
		}
	}()

	for _, value := range partners {
		response := PartnerResponse{
			Name: value.Name,
		}
		partnersView = append(partnersView, response)
	}

	return partnersView, nil
}
