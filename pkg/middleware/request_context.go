package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/satoshi-tahara-st/todo_backend/pkg/application"
)

func RequestContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// AWS x-rayを一旦使わないため
		// header := c.Request().Header

		// return next(&application.RequestContext{
		// 	Context:     c,
		// 	AmznTraceId: header.Get(application.RequestHeaderAmznTraceId),
		// })

		return next(&application.RequestContext{
			Context: c,
		})
	}
}
