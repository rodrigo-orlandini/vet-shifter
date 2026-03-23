package middleware

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *bodyLogWriter) WriteHeaderNow() {
	w.ResponseWriter.WriteHeaderNow()
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	_, _ = w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *bodyLogWriter) WriteString(s string) (int, error) {
	_, _ = w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func ErrorLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}

		c.Writer = blw
		c.Next()

		statusCode := blw.Status()
		if statusCode < http.StatusBadRequest {
			return
		}

		bodyStr := strings.TrimSpace(blw.body.String())
		var payload map[string]any

		const maxBodyLen = 700
		if len(bodyStr) > maxBodyLen {
			bodyStr = bodyStr[:maxBodyLen] + "...(truncated)"
		}

		var codeVal, errVal string

		if bodyStr != "" && json.Unmarshal([]byte(bodyStr), &payload) == nil {
			if v, ok := payload["code"].(string); ok {
				codeVal = v
			}

			if v, ok := payload["error"].(string); ok {
				errVal = v
			}
		}

		var goErr string
		if len(c.Errors) > 0 {
			goErr = c.Errors.String()
		}

		if codeVal != "" || errVal != "" {
			slog.Warn("erro na API",
				"method", c.Request.Method,
				"path", c.Request.URL.Path,
				"status", statusCode,
				"code", codeVal,
				"error", errVal,
			)
		} else {
			slog.Warn("erro na API",
				"method", c.Request.Method,
				"path", c.Request.URL.Path,
				"status", statusCode,
				"body", bodyStr,
			)
		}

		if goErr != "" {
			slog.Warn("erro na API (gin)",
				"method", c.Request.Method,
				"path", c.Request.URL.Path,
				"cause", goErr,
			)
		}
	}
}
