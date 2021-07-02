package datamongo

import (
	"context"

	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Schema name constants
const (
	PartnerCollection = "Partners"
)

type partnerRepo struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func newPartnerRepo(db *mongo.Database) *partnerRepo {
	return &partnerRepo{
		db:         db,
		collection: db.Collection(PartnerCollection),
	}
}

func (r *partnerRepo) Add(ctx context.Context, partner entity.Partner) (err error) {

	document, err := bson.Marshal(partner)
	if err != nil {
		return err
	}

	_, err = r.collection.InsertOne(ctx, document)
	if err != nil {
		return err
	}

	return nil
}
