package gyn

import (
	"log"
	"time"
)

var DB = make(map[string]string)

type User struct {
	Name string `json:"name" binding:"required"`
	Age  string `json:"age" binding:"required"`
}

func onlyForV2() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Status(200)
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func Example() {
	r := Default()

	r.Static("/static", "./static")
	r.LoadHTMLTemplates("templates/*")

	r.GET("/users", func(c *Context) {
		arr := make([]*User, 0)
		for k, v := range DB {
			arr = append(arr, &User{Name: k, Age: v})
		}
		c.HTML(200, "arr.html", H{"title": "gyn", "stuArr": arr})
	})

	r.GET("/panic", func(c *Context) {
		names := []string{"abc"}
		c.String(200, names[100])
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/user/:name", func(c *Context) {
			name := c.Param("name")
			if age, ok := DB[name]; ok {
				c.JSON(200, H{"name": name, "age": age})
			} else {
				c.JSON(201, H{"name": name, "status": "404"})
			}
		})
	}

	v2 := r.Group("/v2")
	v2.Use(BasicAuth([]string{"foo:bar", "test:123"}))
	v2.Use(onlyForV2())
	{
		v2.POST("/admin", func(c *Context) {
			var json User
			if c.EnsureBody(&json) {
				DB[json.Name] = json.Age
				c.JSON(200, H{"status": "ok"})
			}
		})
	}

	r.Run(":8000")
}
