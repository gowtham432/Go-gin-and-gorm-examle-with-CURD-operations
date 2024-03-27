package main

import (
	"gogin-basics/internal/logger"
	"gogin-basics/server"
)

func main() {
	if err := server.Run(); err!=nil{
		logger.NewLogger("main").Warn(err.Error())
	}
}
