package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func LogCollect(l *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bodyBytes, _ := io.ReadAll(c.Request().Body)

			if !strings.Contains(fmt.Sprint(c.Request().URL), "/health") {
				checkMap := map[string]interface{}{}
				logData := bodyBytes
				json.Unmarshal(logData, &checkMap)

				l.WithFields(logrus.Fields{
					"trace_id":   c.Get("trace_id"),
					"method":     c.Request().Method,
					"uri":        fmt.Sprint(c.Request().URL),
					"ip":         fmt.Sprint(c.Request().RemoteAddr),
					"request":    string(logData),
					"user_agent": c.Request().UserAgent(),
					"host":       c.Request().Host,
				}).Info("log request collect")
			}

			c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			return next(c)
		}
	}
}

func LoggingResponseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Create a new buffer for the response body
		buf := new(bytes.Buffer)

		// Use a custom writer to write the response to the buffer
		c.Response().Writer = &BodyDumpResponseWriter{
			ResponseWriter: c.Response().Writer,
			Buffer:         buf,
		}

		// Call the next middleware or handler
		err := next(c)

		// Log the contents of the response body
		logrus.WithFields(logrus.Fields{
			"response": buf.String(),
		}).Info("Response Body")
		buf.Reset()
		return err
	}
}

type BodyDumpResponseWriter struct {
	http.ResponseWriter
	Buffer *bytes.Buffer
}

func (w *BodyDumpResponseWriter) Write(b []byte) (int, error) {
	w.Buffer.Write(b)
	return w.ResponseWriter.Write(b)
}

func TraceIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			traceID := req.Header.Get("x-trace-id")
			if traceID == "" {
				traceID = uuid.NewV4().String()
			}
			c.Set("trace_id", traceID)
			return next(c)
		}

	}
}
