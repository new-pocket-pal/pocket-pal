package routers

import (
	"net/http"

	"pocket-pal/src/domain/service"

	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo, s *service.Service) {
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
}
