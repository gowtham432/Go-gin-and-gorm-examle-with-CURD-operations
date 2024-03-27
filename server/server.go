package server

import (
	"gogin-basics/controller"
	"gogin-basics/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Address  string
	BasePath string
}

var (
	engine *gin.Engine
	config = Config{
		Address:  ":4000",
		BasePath: "/myUsers",
	}
)

func noRouteResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    "404",
		"message": "Not Found",
	})
}

func noMethodResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    "405",
		"message": "Method Not Allowed",
	})
}

func Init() {
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	addMiddleWare()

	initializeDao()
	addRoutes()

}

func addMiddleWare() {
	for _, m := range []gin.HandlerFunc{
		middleware.Logger(),
		middleware.Timeout(),
	} {
		engine.Use(m)
	}
}

func addRoutes() {
	engine.HandleMethodNotAllowed = true
	engine.NoRoute(noRouteResponse)
	engine.NoMethod(noMethodResponse)
	router := engine.Group(config.BasePath)
	controller.Init(router, middleware.PrometheusHandler())
}

func initializeDao() error {
	return controller.InitSampleDb("root", "Sai@1243")
}

func Run() error {
	Init()
	return engine.Run(config.Address)
}
