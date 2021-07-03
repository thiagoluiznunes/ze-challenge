package service

import (
	"context"
	"errors"

	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
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
	if mongo.IsDuplicateKeyError(err) && err != nil {
		return errors.New("partner already registered")
	} else if err != nil {
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
