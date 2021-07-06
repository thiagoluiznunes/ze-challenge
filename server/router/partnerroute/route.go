package partnerroute

import (
	"github.com/labstack/echo/v4"
)

const (
	rootRoute        = ""
	getPartnerByID   = "/:id"
	getAllPartners   = "/all"
	getPartnerNearby = "/"
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

	router.POST(rootRoute, r.ctrl.handleAddPartner)
	router.GET(getPartnerByID, r.ctrl.handleGetPartnerByID)
	router.GET(getPartnerNearby, r.ctrl.handleGetPartnerNearby)
	router.GET(getAllPartners, r.ctrl.handleGetAllPartners)
}
