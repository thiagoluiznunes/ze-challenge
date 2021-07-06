package routeutils

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thiagoluiznunes/ze-challenge/infra/zerrors"
)

// resultWrapper has fields for standard message responses
type resultWrapper struct {
	Error   bool        `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseAPIOK returns a standard API success response
func ResponseAPIOK(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

// ResponseAPIError returns a standard API error to the response
func ResponseAPIError(c echo.Context, data *zerrors.GeneralError) error {

	returnValue := resultWrapper{
		Error:   true,
		Message: data.Message,
	}

	return c.JSON(data.Code, returnValue)
}
