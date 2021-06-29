package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
	"github.com/thiagoluiznunes/ze-challenge/server/router"
)

// Server runs the application
type Server struct {
	Echo           *echo.Echo
	cfg            *config.Config
	middlewares    []echo.MiddlewareFunc
	preMiddlewares []echo.MiddlewareFunc
	appRouters     []router.Router
}

// Instance returns an instance of Service
func Instance(e *echo.Echo, cfg *config.Config) *Server {
	return &Server{
		Echo: e,
		cfg:  cfg,
	}
}

// Run setup and executes the server
func (s *Server) Run() error {

	s.Echo.HideBanner = true
	s.Echo.Debug = s.cfg.Debug

	s.registerMiddlewares()
	s.registerAppRoutes()

	err := s.Echo.Start(fmt.Sprintf(":%d", s.cfg.HTTPPort))
	if err != nil {
		return errors.New("server: fail to start echo")
	}

	return nil
}

// AddMiddleware add middleware to server
func (s *Server) AddMiddleware(md echo.MiddlewareFunc) {
	s.middlewares = append(s.middlewares, md)
}

func (s *Server) registerMiddlewares() {
	for _, md := range s.middlewares {
		s.Echo.Use(md)
	}
}

func (s *Server) registerAppRoutes() {
	for _, appRouter := range s.appRouters {
		appRouter.Register(s.Echo)
	}
}

// AddAppRouter adds an app router to an server
func (s *Server) AddAppRouter(r *router.Router) {
	s.appRouters = append(s.appRouters, *r)
}
