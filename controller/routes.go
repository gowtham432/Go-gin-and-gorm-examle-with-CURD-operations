package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Init(router *gin.RouterGroup, metricsHandler gin.HandlerFunc) {
	router.GET("/health/liveness", liveness)
	router.GET("/health/readiness", readiness)
	router.GET("/metrics", metricsHandler)
	router.POST("/createUser", createUserProfile)
	router.GET("/getAllUsers", getAllUsers)
	router.GET("/getUserById/:id", getUserById)
	router.PUT("/updateUserById/:id", updateUserById)
	router.DELETE("/deleteUserById/:id", deleteUserById)
}

func defaultResponse(c *gin.Context, status int) {
	code := strconv.Itoa(status)
	msg := http.StatusText(status)
	c.JSON(status, Response{Code: code, Message: msg})
}

func errorResponse(c *gin.Context, status int, err error) {
	if err != nil {
		c.Error(err)
	}
	code := strconv.Itoa(status)
	msg := http.StatusText(status)
	c.AbortWithStatusJSON(status, Response{Code: code, Message: msg})
}
func statusOk(c *gin.Context) {
	defaultResponse(c, http.StatusOK)
}

func StatusNotFound(c *gin.Context) {
	defaultResponse(c, http.StatusNotFound)
}

func statusInternalServerError(c *gin.Context, err error) {
	errorResponse(c, http.StatusInternalServerError, err)
}