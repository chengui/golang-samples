package main

import (
    "time"
    "log"
    "gee-demo/gee"
)

func onlyForV2() gee.HandlerFunc {
    return func(c *gee.Context) {
        t := time.Now()
        c.String(400, "Bad Request: %v", c.Path)
        log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
    }
}

func main() {
    r := gee.New()
    r.Use(gee.Logger())
    r.Static("/static", "./static")
    r.GET("/index", func (c *gee.Context) {
        c.HTML(200, "<h1>Index Page</h1>")
    })
    v1 := r.Group("/v1")
    {
        v1.GET("/", func (c *gee.Context) {
            c.HTML(200, "<h1>Hello Gee</h1>")
        })
        v1.GET("/hello", func (c *gee.Context) {
            c.String(200, "Hello %s you're at %s\n", c.Query("name"), c.Path)
        })
    }
    v2 := r.Group("/v2")
    v2.Use(onlyForV2())
    {
        v2.GET("/hello/:name", func (c *gee.Context) {
            c.String(200, "Hello %s you're at %s\n", c.Param("name"), c.Path)
        })
        v2.POST("/login", func (c *gee.Context) {
            c.JSON(200, gee.H{
                "username": c.PostForm("username"),
                "password": c.PostForm("password"),
            })
        })
    }
    r.Run(":8000")
}
