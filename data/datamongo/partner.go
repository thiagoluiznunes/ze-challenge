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

func (r *partnerRepo) GetAll(ctx context.Context) (partners []entity.Partner, err error) {

	cur, err := r.collection.Find(ctx, bson.D{})
	partners, err = parsePartnerEntitySet(ctx, cur, err)
	if err != nil {
		return nil, err
	}

	return partners, err
}

func parsePartnerEntitySet(ctx context.Context, rows *mongo.Cursor, err error) ([]entity.Partner, error) {

	if err != nil {
		return nil, err
	}

	result := make([]entity.Partner, 0)
	for rows.Next(ctx) {
		var document entity.Partner
		err := rows.Decode(&document)
		if err != nil {
			return nil, err
		}
		result = append(result, document)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, err
}
