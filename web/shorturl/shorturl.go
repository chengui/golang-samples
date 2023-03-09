package shorturl

import (
	"github.com/gin-gonic/gin"
)

func Run(addr string) error {
	app := gin.Default()
	hdlr := NewHandler()
	app.GET("/shorten", hdlr.Shorten)
	app.GET("/s/:code", hdlr.Expand)
	return app.Run(addr)
}
