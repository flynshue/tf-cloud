package api

import (
	"bytes"
	"html/template"
)

type RestResource struct {
	Method   string
	Endpoint string
	Router   *CBRouter
}

func NewRestResource(method, endpoint string, router *CBRouter) *RestResource {
	return &RestResource{
		Method:   method,
		Endpoint: endpoint,
		Router:   router,
	}
}

func (r *RestResource) RenderEndpoint(params map[string]string) string {
	if params == nil {
		return r.Endpoint
	}
	t, err := template.New("").Parse(r.Endpoint)
	if err != nil {
		panic(err)
	}
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, params); err != nil {
		panic(err)
	}
	return buf.String()
}
