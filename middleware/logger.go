package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	format = `%s request:"%s" response_code:%d response_timez_ms:%d bytes_sent:%d query_string:"%s"` + "\n"
)

var (
	blacklist = []string{
		"/health/**",
		"/metrics",
	}
)

//the below function is used to log details and we donot need loggers for black list paths
func Logger() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: formatter,
		SkipPaths: blacklist,
	})
}

func formatter(params gin.LogFormatterParams) string {
	return fmt.Sprintf(
		format,
		params.ClientIP,
		params.Request.URL,
		params.StatusCode,
		params.Latency,
		params.BodySize,
		params.Request.URL.Query(),
	)
}
