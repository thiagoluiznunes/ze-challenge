package partnerroute

import (
	"github.com/labstack/echo/v4"
)

const (
	rootRoute = "/"
)

//FormRoute holds user router
type FormRoute struct {
	ctrl *Controller
	name string
}

//NewRouter returns an instance of FormRoute
func NewRouter(routeName string, ctrl *Controller) *FormRoute {
	return &FormRoute{
		ctrl: ctrl,
		name: routeName,
	}
}

// Register registers the routes in the echo group
func (r *FormRoute) Register(e *echo.Group) {

	router := e.Group("/" + r.name)

	router.POST(rootRoute, r.ctrl.handlePartner)
}
