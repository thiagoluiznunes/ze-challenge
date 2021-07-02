package datamongo

import (
	"context"

	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

// Schema name constants
const (
	SchemaPartner = "Partners"
)

type partnerRepo struct {
	client *mongo.Client
}

func newPartnerRepo(client *mongo.Client) *partnerRepo {
	return &partnerRepo{
		client: client,
	}
}

func (r *partnerRepo) Add(ctx context.Context, partner entity.Partner) (err error) {
	return nil
}
