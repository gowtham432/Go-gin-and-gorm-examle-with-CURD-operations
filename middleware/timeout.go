package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

const timeoutDuration = 5000 * time.Millisecond

func timeoutHandler(c *gin.Context){
	c.Next()
}

func timeoutResponse(c *gin.Context){
	code := strconv.Itoa(http.StatusGatewayTimeout)
	msg := http.StatusText(http.StatusGatewayTimeout)
	c.JSON(http.StatusGatewayTimeout, gin.H{
		"code": code,
		"msg":msg,
	})
	c.Writer.WriteHeaderNow()
	c.Abort()
}

func Timeout() gin.HandlerFunc{
	return timeout.New(
		timeout.WithTimeout(timeoutDuration),
		timeout.WithHandler(timeoutHandler),
		timeout.WithResponse(timeoutResponse),
	)
}