package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
	Keys       map[string]interface{}
	Params     map[string]string
	engine     *Engine
	handlers   []HandlerFunc
	index      int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		Keys:   make(map[string]interface{}),
		Params: make(map[string]string),
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err error) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err.Error()})
}

func (c *Context) Set(k string, v interface{}) {
	c.Keys[k] = v
}

func (c *Context) Get(k string) interface{} {
	if v, ok := c.Keys[k]; ok {
		return v
	} else {
		return nil
	}
}

func (c *Context) EnsureBody(item interface{}) bool {
	decoder := json.NewDecoder(c.Req.Body)
	if err := decoder.Decode(&item); err != nil {
		c.Fail(400, err)
		return false
	}
	return true
}

func (c *Context) Param(key string) string {
	if v, ok := c.Params[key]; ok {
		return v
	}
	return ""
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err)
	}
}
