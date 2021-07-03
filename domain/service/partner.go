package service

import (
	"context"

	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
)

type PartnerService struct {
	svc *Service
}

func NewPartnerService(svc *Service) (service *PartnerService) {
	return &PartnerService{
		svc: svc,
	}
}

func (s *PartnerService) Add(ctx context.Context, partner entity.Partner) (err error) {

	err = s.svc.db.Partner().Add(ctx, partner)
	if err != nil {
		return err
	}

	return nil
}

func (s *PartnerService) GetAll(ctx context.Context) (partners []entity.Partner, err error) {

	partners, err = s.svc.db.Partner().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return partners, nil
}
