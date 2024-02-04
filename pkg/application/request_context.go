package application

import "github.com/labstack/echo/v4"

// type RequestContext struct {
// 	echo.Context
// 	AmznTraceId string
// }

// AWS x-rayを一旦使わないため
type RequestContext struct {
	echo.Context
}
