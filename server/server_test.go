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
	t.Run("fail: test service", TestEcho)
	t.Run("fail: test controller", TestController)
	t.Run("fail: test route", TestRouter)
	t.Run("fail: test echo", TestService)
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

func TestController(t *testing.T) {

}

func TestRouter(t *testing.T) {

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
	})

	t.Run("fail: add middleware", func(t *testing.T) {
		srv.AddMiddleware(middleware.Logger())
		assert.Len(t, srv.middlewares, 1)
	})

	t.Run("fail: add router", func(t *testing.T) {
		var svc *service.Service
		var partnerService *service.PartnerService

		t.Run("fail: instance service", func(t *testing.T) {
			svc, err = service.New(nil, &cfg)
			assert.Nil(t, err)
			assert.NotEmpty(t, svc)
		})

		t.Run("fail: instance partner service", func(t *testing.T) {
			partnerService = service.NewPartnerService(svc)
			assert.NotEmpty(t, partnerService)
		})

		// srv.AddAppRouter()
		assert.Len(t, srv.middlewares, 1)
	})

}
