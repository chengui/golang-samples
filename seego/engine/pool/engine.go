package pool

import (
	"strings"
	"time"
)

type Engine struct {
	ID int
}

func NewEngine(id int) *Engine {
	time.Sleep(200 * time.Millisecond)
	return &Engine{ID: id}
}

func (e *Engine) Predict(str string) (string, error) {
	return strings.ToUpper(str), nil
}
