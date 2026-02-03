package recorder

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type Traffic struct {
	Method   string            `json:"method"`
	Path     string            `json:"path"`
	Headers  map[string]string `json:"headers"`
	Body     string            `json:"body"`
	Response int               `json:"response"`
	Time     time.Time         `json:"time"`
}

var Records []Traffic

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()

		headers := map[string]string{}
		for k, v := range c.Request.Header {
			headers[k] = v[0]
		}

		Records = append(Records, Traffic{
			Method:   c.Request.Method,
			Path:     c.FullPath(),
			Headers:  headers,
			Body:     string(bodyBytes),
			Response: c.Writer.Status(),
			Time:     time.Now(),
		})
	}
}
