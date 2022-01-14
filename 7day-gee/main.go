package main

import (
	"gee-demo/gee"
	"log"
	"time"
)

var DB = make(map[string]string)

type User struct {
	Name string `json:"name" binding:"required"`
	Age  string `json:"age" binding:"required"`
}

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Status(200)
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.Default()

	r.Static("/static", "./static")
	r.LoadHTMLTemplates("templates/*")

	r.GET("/users", func(c *gee.Context) {
		arr := make([]*User, 0)
		for k, v := range DB {
			arr = append(arr, &User{Name: k, Age: v})
		}
		c.HTML(200, "arr.html", gee.H{"title": "gee", "stuArr": arr})
	})

	r.GET("/panic", func(c *gee.Context) {
		names := []string{"abc"}
		c.String(200, names[100])
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/user/:name", func(c *gee.Context) {
			name := c.Param("name")
			if age, ok := DB[name]; ok {
				c.JSON(200, gee.H{"name": name, "age": age})
			} else {
				c.JSON(201, gee.H{"name": name, "status": "404"})
			}
		})
	}

	v2 := r.Group("/v2")
	v2.Use(gee.BasicAuth([]string{"foo:bar", "test:123"}))
	// v2.Use(onlyForV2())
	{
		v2.POST("/admin", func(c *gee.Context) {
			var json User
			if c.EnsureBody(&json) {
				DB[json.Name] = json.Age
				c.JSON(200, gee.H{"status": "ok"})
			}
		})
	}

	r.Run(":8000")
}
