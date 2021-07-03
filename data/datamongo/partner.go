package datamongo

import (
	"context"

	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func setPartnerIndexes(db *mongo.Database) (err error) {
	_, err = db.Collection(PartnerCollection).Indexes().CreateMany(
		context.TODO(),
		[]mongo.IndexModel{
			{
				Keys:    bson.D{{Key: "document", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
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

func (r *partnerRepo) GetByID(ctx context.Context, id string) (partner entity.PartnerDoc, err error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return partner, err
	}
	filter := bson.D{{"_id", objectID}}

	err = r.collection.FindOne(ctx, filter).Decode(&partner)
	if err == mongo.ErrNoDocuments {
		return partner, err
	} else if err != nil {
		return partner, err
	}

	return partner, nil
}

func (r *partnerRepo) GetAll(ctx context.Context) (partners []entity.PartnerDoc, err error) {

	cur, err := r.collection.Find(ctx, bson.D{})
	partners, err = parsePartnerEntitySet(ctx, cur, err)
	if err != nil {
		return nil, err
	}

	return partners, err
}

func parsePartnerEntitySet(ctx context.Context, rows *mongo.Cursor, err error) ([]entity.PartnerDoc, error) {

	defer rows.Close(ctx)

	if err != nil {
		return nil, err
	}

	result := make([]entity.PartnerDoc, 0)
	for rows.Next(ctx) {
		var document entity.PartnerDoc
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
