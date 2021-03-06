package contract

import (
	"context"

	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
)

type RepoManager interface {
	Partner() PartnerRepo
}

type PartnerRepo interface {
	Add(ctx context.Context, partner entity.Partner) (partnerID string, err error)
	AddInBatch(ctx context.Context, partners []entity.Partner) (err error)
	GetByID(ctx context.Context, id string) (partner entity.Partner, err error)
	GetAll(ctx context.Context) (partners []entity.Partner, err error)
}
