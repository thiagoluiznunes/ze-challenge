package service

import (
	"context"
	"errors"
	"math"

	geo "github.com/kellydunn/golang-geo"
	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
	"github.com/thiagoluiznunes/ze-challenge/infra/zerrors"
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

func checkIsDuplicateKeyError(err error) error {
	if mongo.IsDuplicateKeyError(err) && err != nil {
		return zerrors.NewConflictError(zerrors.PartnerAlreadyExistError)
	} else if err != nil {
		return err
	}
	return nil
}

func (s *PartnerService) Add(ctx context.Context, partner entity.Partner) (err error) {

	err = s.svc.db.Partner().Add(ctx, partner)
	err = checkIsDuplicateKeyError(err)
	if err != nil {
		return err
	}

	return nil
}
func (s *PartnerService) AddInBatch(ctx context.Context, partners []entity.Partner) (err error) {

	err = s.svc.db.Partner().AddInBatch(ctx, partners)
	err = checkIsDuplicateKeyError(err)
	if err != nil {
		return zerrors.NewApplicationError(err)
	}

	return nil
}

func (s *PartnerService) GetByID(ctx context.Context, id string) (partner entity.Partner, err error) {

	partner, err = s.svc.db.Partner().GetByID(ctx, id)
	if err == mongo.ErrNoDocuments {
		return partner, zerrors.NewNotFoundError(err)
	} else if err != nil {
		return partner, zerrors.NewApplicationError(err)
	}

	return partner, nil
}

func (s *PartnerService) GetAll(ctx context.Context) (partners []entity.Partner, err error) {

	partners, err = s.svc.db.Partner().GetAll(ctx)
	if len(partners) == 0 {
		return partners, zerrors.NewNotFoundError(err)
	} else if err != nil {
		return partners, zerrors.NewApplicationError(err)
	}

	return partners, nil
}

func (s *PartnerService) GetNearby(ctx context.Context, point entity.Point) (partner entity.Partner, err error) {

	partners, err := s.svc.db.Partner().GetAll(ctx)
	if err != nil {
		return partner, zerrors.NewApplicationError(err)
	} else if len(partners) <= 0 {
		return partner, zerrors.NewNotFoundError(errors.New(zerrors.PartnerNotFoundError))
	}

	geoPoint := geo.NewPoint(point.Coordinates[0], point.Coordinates[1])
	partner, err = getClosestPartnerByArea(geoPoint, partners)
	if err != nil {
		return partner, zerrors.NewApplicationError(err)
	} else if partner.ID == "" {
		return partner, zerrors.NewNotFoundError(errors.New(zerrors.PartnerNotFoundError))
	}

	return partner, nil
}

func getClosestPartnerByArea(point *geo.Point, partners []entity.Partner) (closestPartner entity.Partner, err error) {

	var closestDistance float64
	defer func() {
		if recover() != nil {
			err = zerrors.NewApplicationError(errors.New(zerrors.GetClosestPartnerByAreaError))
		}
	}()

	for _, value := range partners {
		for _, zvalue := range value.CoverageArea.Coordinates {
			var arrayPoints []*geo.Point
			for _, zcoord := range zvalue[0] {
				arrayPoints = append(arrayPoints, geo.NewPoint(zcoord[0], zcoord[1]))
			}
			polygon := geo.NewPolygon(arrayPoints)
			if polygon.Contains(point) {
				if closestDistance == 0 {
					closestPartner = value
					closestDistance = distance(point.Lat(), point.Lng(), closestPartner.Address.Coordinates[1], closestPartner.Address.Coordinates[1])
				} else {
					valueDistance := distance(point.Lat(), point.Lng(), value.Address.Coordinates[1], value.Address.Coordinates[1])
					if valueDistance < closestDistance {
						closestDistance = valueDistance
						closestPartner = value
					}
				}
			}
		}
	}

	return closestPartner, err
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, earthRadius float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	earthRadius = 6378100 // Earth radius in METERS

	// calculate
	val := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * earthRadius * math.Asin(math.Sqrt(val))
}
