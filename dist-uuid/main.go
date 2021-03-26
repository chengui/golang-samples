package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dist-uuid/uuid"
)

func main() {
	options := &uuid.Options{
		EpochTime: "2021-03-01",
		LocalFile: "./machineID",
		RedisAddr: "127.0.0.1:6379",
		RedisPass: "",
		RedisDB:   0,
	}
	uuidGen := uuid.NewUUID(options)

	srv := gin.Default()
	srv.GET("/uuid", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"uuid": uuidGen.Generate(),
		})
	})
	srv.Run(":8001")
}
