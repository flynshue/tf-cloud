package api

import (
	"fmt"
	"net/http"
)

type RouterFunc func(resp *http.Response) error

type CBRouter struct {
	Routers       map[int]RouterFunc
	DefaultRouter RouterFunc
}

func NewRouter() *CBRouter {
	return &CBRouter{
		Routers: make(map[int]RouterFunc),
		DefaultRouter: func(resp *http.Response) error {
			return fmt.Errorf("%d unknown status from %s", resp.StatusCode, resp.Request.URL.Path)
		},
	}
}

func (c *CBRouter) AddFunc(status int, handler RouterFunc) {
	c.Routers[status] = handler
}

func (c *CBRouter) CallFunc(resp *http.Response) error {
	f, ok := c.Routers[resp.StatusCode]
	if !ok {
		f = c.DefaultRouter
	}
	return f(resp)
}
