package server

import (
	"encoding/json"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/thiagoluiznunes/ze-challenge/domain"
	"github.com/thiagoluiznunes/ze-challenge/domain/service"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
	"github.com/thiagoluiznunes/ze-challenge/server/router"
	"github.com/thiagoluiznunes/ze-challenge/server/router/partnerroute"
)

var err error
var cfg config.Config

var ech *echo.Echo
var srv *Server

var partnerCtrl *partnerroute.Controller
var partnerRoute *partnerroute.PartnerRoute

var svc *service.Service
var partnerService *service.PartnerService

func TestServer(t *testing.T) {
	// <setup code>
	t.Run("fail: test service", TestService)
	t.Run("fail: test echo", TestEcho)
	// <tear-down code>
}

func TestService(t *testing.T) {

	t.Run("fail: instance service", func(t *testing.T) {
		svc, err = service.New(nil, &cfg)
		assert.Nil(t, err)
		assert.NotEmpty(t, svc)
	})

	t.Run("fail: instance partner service", func(t *testing.T) {
		partnerService = service.NewPartnerService(svc)
		assert.NotEmpty(t, partnerService)
	})
}

func TestEcho(t *testing.T) {

	t.Run("fail: instance echo", func(t *testing.T) {
		ech = echo.New()
		assert.NotEmpty(t, ech)
	})

	t.Run("fail: instance config", func(t *testing.T) {
		err := json.Unmarshal([]byte(domain.MockConfig), &cfg)
		assert.Nil(t, err)
	})

	t.Run("fail: instance server", func(t *testing.T) {
		srv = Instance(ech, &cfg)
		assert.NotEmpty(t, srv)
		assert.NotNil(t, srv)
	})

	t.Run("fail: instance controller", func(t *testing.T) {
		partnerCtrl = partnerroute.NewController(&cfg, partnerService)
		assert.NotEmpty(t, partnerCtrl)
		assert.NotNil(t, partnerCtrl)
	})

	t.Run("fail: instance partner route", func(t *testing.T) {
		partnerRoute = partnerroute.NewRouter("partner", partnerCtrl)
		assert.NotEmpty(t, partnerRoute)
		assert.NotNil(t, partnerRoute)
	})

	var appRouter *router.Router
	t.Run("fail: instance app router", func(t *testing.T) {
		appRouter = router.New(srv.Echo, &cfg, "ze-delivery")
		assert.NotEmpty(t, appRouter)
		assert.NotNil(t, appRouter)
		appRouter.AddRouters(partnerRoute)
	})

	t.Run("fail: add app router ", func(t *testing.T) {
		srv.AddAppRouter(appRouter)
		assert.Len(t, srv.appRouters, 1)

	})

	t.Run("fail: add middleware", func(t *testing.T) {
		srv.AddMiddleware(middleware.Logger())
		assert.Len(t, srv.middlewares, 1)
	})
}
