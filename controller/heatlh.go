package controller

import (
	"github.com/gin-gonic/gin"
)

func readiness(c *gin.Context) {
	statusOk(c)
}

func liveness(c *gin.Context) {
	statusOk(c)
}
