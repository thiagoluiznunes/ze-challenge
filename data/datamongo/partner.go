package datamongo

import (
	"context"
	"fmt"

	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
	"github.com/thiagoluiznunes/ze-challenge/server/serverconfig"
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

func (r *partnerRepo) Add(ctx context.Context, partner entity.Partner) (partnerID string, err error) {

	document, err := bson.Marshal(partner)
	if err != nil {
		return partnerID, err
	}

	defer serverconfig.CreateNewRelicMongoDBSegment(ctx, PartnerCollection, "InsertOne")

	cur, err := r.collection.InsertOne(ctx, document)
	if err != nil {
		return partnerID, err
	}
	partnerID = fmt.Sprintf("%v", cur.InsertedID)

	return partnerID, nil
}

func (r *partnerRepo) AddInBatch(ctx context.Context, partners []entity.Partner) (err error) {

	var batch []interface{}
	for _, partner := range partners {
		batch = append(batch, partner)
	}

	defer serverconfig.CreateNewRelicMongoDBSegment(ctx, PartnerCollection, "InsertMany")

	_, err = r.collection.InsertMany(ctx, batch)
	if err != nil {
		return err
	}

	return nil
}

func (r *partnerRepo) GetByID(ctx context.Context, id string) (partner entity.Partner, err error) {

	var filter bson.D
	var objectID primitive.ObjectID

	if primitive.IsValidObjectID(id) {
		objectID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return partner, err
		}
		filter = bson.D{{Key: "_id", Value: objectID}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}}
	}

	defer serverconfig.CreateNewRelicMongoDBSegment(ctx, PartnerCollection, "FindOne")

	err = r.collection.FindOne(ctx, filter).Decode(&partner)
	if err == mongo.ErrNoDocuments {
		return partner, err
	} else if err != nil {
		return partner, err
	}

	return partner, nil
}

func (r *partnerRepo) GetAll(ctx context.Context) (partners []entity.Partner, err error) {

	defer serverconfig.CreateNewRelicMongoDBSegment(ctx, PartnerCollection, "Find")

	cur, err := r.collection.Find(ctx, bson.D{})
	partners, err = parsePartnerEntitySet(ctx, cur, err)
	if err != nil {
		return nil, err
	}

	return partners, err
}

func parsePartnerEntitySet(ctx context.Context, rows *mongo.Cursor, err error) ([]entity.Partner, error) {

	defer rows.Close(ctx)

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
