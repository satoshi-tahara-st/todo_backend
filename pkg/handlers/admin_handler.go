package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
}

func NewAdminHandler() AdminHandler {
	return AdminHandler{}
}

func (h AdminHandler) HealthCheck(ec echo.Context) error {
	return ec.String(http.StatusOK, "OK")
}
