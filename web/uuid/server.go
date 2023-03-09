package uuid

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type server struct {
	uuid    *UUID
	options *Options
}

func NewServer(options *Options) *server {
	if options == nil {
		options = &Options{
			EpochTime: "2021-03-01",
			LocalFile: "./machineID",
			RedisAddr: "127.0.0.1:6379",
			RedisPass: "",
			RedisDB:   0,
		}
	}
	uuid := NewUUID(options)
	return &server{uuid, options}
}

func (s *server) Run(addr string) error {
	srv := gin.Default()
	srv.GET("/uuid", func(c *gin.Context) {
		uid, err := s.uuid.Generate()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"uuid": uid,
			})
		}
	})
	return srv.Run(addr)
}
