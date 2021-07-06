package contract

import (
	"context"

	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
)

type PartnerService interface {
	Add(ctx context.Context, partner entity.Partner) (err error)
	AddInBatch(ctx context.Context, partners []entity.Partner) (err error)
	GetByID(ctx context.Context, id string) (partner entity.Partner, err error)
	GetAll(ctx context.Context) (partners []entity.Partner, err error)
	GetNearby(ctx context.Context, point entity.Point) (partner entity.Partner, err error)
}
