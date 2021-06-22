package main

import (
    "time"
    "log"
    "gee-demo/gee"
)

func onlyForV2() gee.HandlerFunc {
    return func(c *gee.Context) {
        t := time.Now()
        c.Status(200)
        log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
    }
}

type student struct {
    Name string
    Age  int
}

func main() {
    r := gee.Default()
    r.LoadHTMLTemplates("templates/*")
    r.Static("/static", "./static")
    /*
    r.GET("/index", func (c *gee.Context) {
        c.HTML(200, "<h1>Index Page</h1>")
    })
    */
    r.GET("/students", func(c *gee.Context) {
        stu1 := &student{Name: "Abs", Age: 20}
        stu2 := &student{Name: "Cos", Age: 10}
        c.HTML(200, "arr.html", gee.H{
            "title": "gee",
            "stuArr": [2]*student{stu1, stu2},
        })
    })
    r.GET("/panic", func(c *gee.Context) {
        names := []string{"abc"}
        c.String(200, names[100])
    })
    v1 := r.Group("/v1")
    {
        /*
        v1.GET("/", func(c *gee.Context) {
            c.HTML(200, "<h1>Hello Gee</h1>")
        })
        */
        v1.GET("/hello", func(c *gee.Context) {
            c.String(200, "Hello %s you're at %s\n", c.Query("name"), c.Path)
        })
    }
    v2 := r.Group("/v2")
    // v2.Use(onlyForV2())
    {
        v2.GET("/hello/:name", func(c *gee.Context) {
            c.String(200, "Hello %s you're at %s\n", c.Param("name"), c.Path)
        })
        v2.POST("/login", func(c *gee.Context) {
            c.JSON(200, gee.H{
                "username": c.PostForm("username"),
                "password": c.PostForm("password"),
            })
        })
    }
    r.Run(":8000")
}
