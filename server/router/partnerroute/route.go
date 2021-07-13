package partnerroute

import (
	"github.com/labstack/echo/v4"
)

const (
	rootRoute        = ""
	addInBatch       = "/batch"
	getPartnerByID   = "/:id"
	getAllPartners   = "/all"
	getPartnerNearby = "/"
)

//PartnerRoute holds user router
type PartnerRoute struct {
	ctrl *Controller
	name string
}

//NewRouter returns an instance of PartnerRoute
func NewRouter(routeName string, ctrl *Controller) *PartnerRoute {
	return &PartnerRoute{
		ctrl: ctrl,
		name: routeName,
	}
}

// Register registers the routes in the echo group
func (r *PartnerRoute) Register(e *echo.Group) {

	router := e.Group("/" + r.name)

	router.POST(rootRoute, r.ctrl.handleAddPartner)
	router.POST(addInBatch, r.ctrl.handleAddPartnerInBatch)
	router.GET(getPartnerByID, r.ctrl.handleGetPartnerByID)
	router.GET(getPartnerNearby, r.ctrl.handleGetPartnerNearby)
	router.GET(getAllPartners, r.ctrl.handleGetAllPartners)
}
