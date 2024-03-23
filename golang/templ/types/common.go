package types

import (
	"context"
	"net/http"
)

type HelHandler func(*Context) error

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	ctx            context.Context
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		ResponseWriter: w,
		Request:        r,
		ctx:            context.Background(),
	}
}

type Data struct {
	Status  string      `json:"status"`
	Content interface{} `json:"content"`
}
