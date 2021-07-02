package partnerroute

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
)

var (
	instance *Controller
	once     sync.Once
)

// Controller holds user controller operations
type Controller struct {
	cfg *config.Config
}

//NewController returns an instance of Controller
func NewController(cfg *config.Config) *Controller {
	once.Do(func() {
		instance = &Controller{
			cfg: cfg,
		}
	})
	return instance
}

func (c *Controller) handlePartner(ctx echo.Context) error {

	return ctx.JSON(200, "OK")
}
