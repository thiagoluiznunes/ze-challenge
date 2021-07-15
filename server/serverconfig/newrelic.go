package serverconfig

import (
	"context"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/thiagoluiznunes/ze-challenge/domain"
)

func NewRelicMiddleware(app *newrelic.Application) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			name := c.Request().Method + " " + c.Path()
			txn := app.StartTransaction(name)
			defer txn.End()

			if strings.HasPrefix(c.Path(), "/health") {
				txn.Ignore()
			}

			c.Set(domain.NRTransactionKey, txn)

			return next(c)
		}
	}
}

func CreateNewRelicMongoDBSegment(ctx context.Context, collectionName string, operation string) *newrelic.DatastoreSegment {

	s := newrelic.DatastoreSegment{
		Product:    newrelic.DatastoreMongoDB,
		Collection: collectionName,
		Operation:  operation,
	}

	tx := ctx.Value(domain.NRTransactionKey)
	if tx != nil {
		s.StartTime = tx.(*newrelic.Transaction).StartSegmentNow()
	}

	return &s
}
