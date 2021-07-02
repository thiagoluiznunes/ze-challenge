package contract

import (
	"context"

	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
)

type RepoManager interface {
	Partner() PartnerRepo
}

type PartnerRepo interface {
	Add(ctx context.Context, partner entity.Partner) (err error)
}
