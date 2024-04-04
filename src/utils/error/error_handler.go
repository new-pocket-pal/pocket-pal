package error

import (
	"pocket-pal/src/utils/response"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func CustomHTTPErrorHandler(l *logrus.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			l.WithError(err).WithFields(logrus.Fields{
				"trace_id": c.Get("trace_id"),
			}).Error("echo http error")

			c.JSON(he.Code, response.Response{
				Meta: response.Meta{
					Code:    he.Code,
					Message: he.Message,
					Msg:     he.Internal.Error(),
				},
				Data: nil,
			})
			return
		}

		// If it's not a validation error, return the default error response
		c.Echo().DefaultHTTPErrorHandler(err, c)

	}
}
