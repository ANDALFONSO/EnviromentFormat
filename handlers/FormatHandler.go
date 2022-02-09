package handlers

import (
	"format_enviroment/service"

	"github.com/gin-gonic/gin"
)

var (
	formatService service.IFormatService = service.NewPS()
)

func FormatHandler(server *gin.Engine) {
	server.POST("/env", func(c *gin.Context) {
		text, exist := c.GetPostForm("text")
		if !exist {
			c.JSON(404, "Error buscando etiqueta")
		}
		c.JSON(200, formatService.Format(text))
	})
}
