package router

import (
	"github.com/labstack/echo/v4"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
)

// IRouter interface for routers
type IRouter interface {
	Register(appGroup *echo.Group)
}

// Router holds application's routers
type Router struct {
	e       *echo.Echo
	cfg     *config.Config
	routers []IRouter
	appName string
}

// New returns an instance of Router
func New(e *echo.Echo, cfg *config.Config, appName string) *Router {
	return &Router{
		e:       e,
		cfg:     cfg,
		appName: appName,
	}
}

// AddRouters add new routers to register
func (r *Router) AddRouters(router IRouter) {
	r.routers = append(r.routers, router)
}

// Register registers router to an application
func (r *Router) Register(e *echo.Echo) {

	prefixGroup := r.e.Group(r.cfg.HTTPPrefix)
	appGroup := prefixGroup.Group("")

	for _, r := range r.routers {
		r.Register(appGroup)
	}
}
