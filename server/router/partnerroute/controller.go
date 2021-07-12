package partnerroute

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/thiagoluiznunes/ze-challenge/domain/contract"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
	"github.com/thiagoluiznunes/ze-challenge/server/routeutils"
	"github.com/thiagoluiznunes/ze-challenge/server/viewmodel"
)

var (
	instance *Controller
	once     sync.Once
)

// Controller holds user controller operations
type Controller struct {
	cfg            *config.Config
	partnerService contract.PartnerService
}

//NewController returns an instance of Controller
func NewController(cfg *config.Config, partnerService contract.PartnerService) *Controller {
	once.Do(func() {
		instance = &Controller{
			cfg:            cfg,
			partnerService: partnerService,
		}
	})
	return instance
}

func (c *Controller) handleAddPartner(ctx echo.Context) (err error) {

	input := viewmodel.PartnerRequest{}
	err = ctx.Bind(&input)
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	partner, err := viewmodel.NewPartner(input)
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	err = partner.Validate()
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	err = c.partnerService.Add(ctx.Request().Context(), partner)
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	return routeutils.ResponseAPIOK(ctx, "OK")
}

func (c *Controller) handleAddPartnerInBatch(ctx echo.Context) (err error) {

	input := viewmodel.PartnerInBatchRequest{}
	err = ctx.Bind(&input)
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	partners, err := viewmodel.NewPartners(input)
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	err = c.partnerService.AddInBatch(ctx.Request().Context(), partners)
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	return routeutils.ResponseAPIOK(ctx, "OK")
}

func (c *Controller) handleGetPartnerByID(ctx echo.Context) (err error) {

	id := ctx.Param("id")
	partner, err := c.partnerService.GetByID(ctx.Request().Context(), id)
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	return routeutils.ResponseAPIOK(ctx, partner)
}

func (c *Controller) handleGetAllPartners(ctx echo.Context) (err error) {

	partners, err := c.partnerService.GetAll(ctx.Request().Context())
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	response, err := viewmodel.ModelsToView(partners)
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	return routeutils.ResponseAPIOK(ctx, response)
}

func (c *Controller) handleGetPartnerNearby(ctx echo.Context) (err error) {

	point, err := viewmodel.NewPoint(ctx.QueryParams())
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	partner, err := c.partnerService.GetNearby(ctx.Request().Context(), point)
	if err != nil {
		return routeutils.HandleAPIError(ctx, err)
	}

	return routeutils.ResponseAPIOK(ctx, partner)
}
