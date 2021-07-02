package entity

import (
	"errors"

	"github.com/thiagoluiznunes/ze-challenge/server/viewmodel"
)

type Partner struct {
	Name string `json:"name,omitempty"`
}

func NewPartner(viewmodel viewmodel.PartnerRequest) (partner Partner, err error) {

	defer func() {
		if recover() != nil {
			err = errors.New("failt to create new partner")
		}
	}()

	partner.Name = viewmodel.Name

	return partner, nil
}

func (p *Partner) validate() (err error) {

	return nil
}
