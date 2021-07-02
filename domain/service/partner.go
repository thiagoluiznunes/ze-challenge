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
