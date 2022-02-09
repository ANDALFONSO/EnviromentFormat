package main

import (
	"format_enviroment/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	handlers.FormatHandler(server)
	server.Run(":9090")
}
