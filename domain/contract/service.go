package contract

import (
	"context"

	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
)

type PartnerService interface {
	Add(ctx context.Context, partner entity.Partner) (err error)
}
