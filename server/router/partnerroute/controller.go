package partnerroute

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/thiagoluiznunes/ze-challenge/domain/contract"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
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
		return ctx.JSON(503, "service unavailable")
	}

	partner, err := viewmodel.NewPartner(input)
	if err != nil {
		return ctx.JSON(503, "service unavailable")
	}

	err = c.partnerService.Add(ctx.Request().Context(), partner)
	if err != nil {
		return ctx.JSON(503, "service unavailable")
	}

	return ctx.JSON(200, "OK")
}

func (c *Controller) handleGetAllPartners(ctx echo.Context) (err error) {

	partners, err := c.partnerService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(503, "service unavailable")
	}

	response, err := viewmodel.ModelToView(partners)
	if err != nil {
		return ctx.JSON(503, err.Error())
	}

	return ctx.JSON(200, response)
}
