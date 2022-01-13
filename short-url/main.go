package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"short-url/handler"
)

func main() {
	addr := ":8000"
	app := gin.Default()
	hdlr := handler.NewHandler()
	app.GET("/shorten", hdlr.Shorten)
	app.GET("/s/:code", hdlr.Expand)
	log.Fatal(app.Run(addr))
}
