package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/satoshi-tahara-st/todo_backend/pkg/application"
	"github.com/satoshi-tahara-st/todo_backend/pkg/di"
	"github.com/satoshi-tahara-st/todo_backend/pkg/handlers"
	dbClient "github.com/satoshi-tahara-st/todo_backend/pkg/infra/db"
	"gorm.io/gorm"

	customMiddleware "github.com/satoshi-tahara-st/todo_backend/pkg/middleware"
)

func NewRouter(conf application.Config) *echo.Echo {
	e := echo.New()

	// middlewareの設定
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisablePrintStack: true,
	}))

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogLatency:       true,
		LogProtocol:      true,
		LogRemoteIP:      false,
		LogHost:          false,
		LogMethod:        true,
		LogURI:           true,
		LogURIPath:       true,
		LogRoutePath:     false,
		LogRequestID:     false,
		LogReferer:       true,
		LogUserAgent:     true,
		LogStatus:        true,
		LogError:         false,
		LogContentLength: true,
		LogResponseSize:  true,
		LogHeaders: []string{
			echo.HeaderAcceptEncoding,
			echo.HeaderContentType,
			// application.RequestHeaderAmznTraceId,
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			// logger.Log.Debug("request_logger", "request", v)
			return nil
		},
	}))

	e.Use(customMiddleware.RequestContextMiddleware)

	e.HTTPErrorHandler = httpErrorHandler

	// DBの設定
	dbs := make(map[string]map[string]*gorm.DB)
	for k, v := range conf.DB {
		refConfig := v.Ref
		refDb, err := dbClient.NewDB(
			dbClient.NewDBConfig(v.Driver, refConfig.Username, refConfig.Password, refConfig.Host, v.DBName, refConfig.MaxOpenConns, refConfig.MaxIdleConns, refConfig.ConnMaxLifetimeSeconds),
		)
		if err != nil {
			panic(fmt.Errorf("fail to connect ref DB %s:\n %w", k, err))
		}

		// updConfig := v.Upd
		// updDb, err := dbClient.NewDB(
		// 	dbClient.NewDBConfig(v.Driver, updConfig.Username, updConfig.Password, updConfig.Host, v.DBName, updConfig.MaxOpenConns, updConfig.MaxIdleConns, updConfig.ConnMaxLifetimeSeconds),
		// )
		// if err != nil {
		// 	panic(fmt.Errorf("fail to connect upd DB %s:\n %w", k, err))
		// }

		db := map[string]*gorm.DB{
			application.ConnectionRef: refDb,
		}
		dbs[k] = db
	}

	// 各APIの設定
	adminHandler := handlers.NewAdminHandler()
	todoHandler := di.TodoHandler(dbs)
	handler := newHandler(adminHandler, todoHandler)

	return e
}

func httpErrorHandler(err error, c echo.Context) {
	// logger.Log.Debug("httpErrorHandler", "err", err)
	sendErr := false
	var he *echo.HTTPError
	if errors.As(err, &he) {
		if he.Code >= 404 && he.Code != http.StatusUnauthorized {
			sendErr = true
		}
	} else {
		sendErr = true
	}
	if sendErr {
		// logger.SendExceptionToSentry(request.RequestContext(c), err)
	}

	c.Echo().DefaultHTTPErrorHandler(err, c)
}

type handler struct {
	adminHandler handlers.AdminHandler
	todoHandler  handlers.TodoHandler
}

func newHandler(
	ah handlers.AdminHandler,
	th handlers.TodoHandler,
) handler {
	return handler{
		adminHandler: ah,
		todoHandler:  th,
	}
}

func (h handler) Health(etx echo.Context) error {
	return h.adminHandler.HealthCheck(etx)
}

func (h handler) GetTodo(etx echo.Context) error {
	return h.todoHandler.Get()
}
