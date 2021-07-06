package service

import (
	"context"
	"errors"
	"fmt"

	geo "github.com/kellydunn/golang-geo"
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
func (s *PartnerService) AddInBatch(ctx context.Context, partners []entity.Partner) (err error) {

	err = s.svc.db.Partner().AddInBatch(ctx, partners)
	if mongo.IsDuplicateKeyError(err) && err != nil {
		return errors.New("partner already registered")
	} else if err != nil {
		return err
	}

	return nil
}

func (s *PartnerService) GetByID(ctx context.Context, id string) (partner entity.Partner, err error) {

	partner, err = s.svc.db.Partner().GetByID(ctx, id)
	if err == mongo.ErrNoDocuments {
		return partner, errors.New("not found")
	} else if err != nil {
		return partner, err
	}

	return partner, nil
}

func (s *PartnerService) GetAll(ctx context.Context) (partners []entity.Partner, err error) {

	partners, err = s.svc.db.Partner().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return partners, nil
}

func (s *PartnerService) GetNearby(ctx context.Context, point entity.Point) (partner entity.Partner, err error) {

	partners, err := s.svc.db.Partner().GetAll(ctx)
	if err != nil {
		return partner, err
	}

	geoPoint := geo.NewPoint(point.Coordinates[1], point.Coordinates[0])
	coverageAreas, _ := includedInCoverageAreas(geoPoint, partners)
	fmt.Println(coverageAreas)

	return partner, nil
}

func includedInCoverageAreas(point *geo.Point, partners []entity.Partner) (areas []*geo.Polygon, err error) {

	for _, value := range partners {
		for _, zvalue := range value.CoverageArea.Coordinates {
			var arrayPoints []*geo.Point
			for _, zcoord := range zvalue[0] {
				arrayPoints = append(arrayPoints, geo.NewPoint(zcoord[1], zcoord[0]))
			}
			polygon := geo.NewPolygon(arrayPoints)
			if polygon.Contains(point) {
				areas = append(areas, polygon)
			}
		}
	}

	return areas, nil
}
