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
func ResponseAPIError(c echo.Context, status int, message string) error {

	returnValue := resultWrapper{
		Error:   true,
		Message: message,
	}

	return c.JSON(status, returnValue)
}

// HandleAPIError applies the default error handling to the response
func HandleAPIError(c echo.Context, errorToHandle error) error {

	statusCode := http.StatusServiceUnavailable
	errorMessage := "Service Unavailable"

	if errorToHandle != nil {
		errorString := errorToHandle.Error()

		switch e := errorToHandle.(type) {
		case *zerrors.GeneralError:
			statusCode = e.Code
			errorMessage = errorString
		}
	}

	return ResponseAPIError(c, statusCode, errorMessage)
}
