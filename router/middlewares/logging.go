package middlewares

import (
	"bytes"
	"github.com/mssola/user_agent"
	"github.com/tomasen/realip"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-xweb/log"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}
		// Read the Body content
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		bodyStr := strings.Replace(strings.Replace(string(bodyBytes), "	", "", -1), "\n", "", -1)
		if bodyStr != "" {
			c.Set("bodyStr", bodyStr)
		}
		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		method := c.Request.Method
		clientIP := realip.FromRequest(c.Request)
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		statusCode := c.Writer.Status()
		ua := user_agent.New(c.GetHeader("User-Agent"))
		log.Infof("| %d | %v | %s | %s %s | %s | %s | %s",
			statusCode,
			latency,
			clientIP,
			method,
			path,
			ua.Platform(),
			ua.OS(),
			"",
		)
	}
}
