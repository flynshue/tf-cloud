package api

import "fmt"

type API struct {
	BaseURL       string
	Client        *Client
	DefaultRouter *CBRouter
	Resources     map[string]*RestResource
}

func NewAPI(baseURL string) *API {
	return &API{
		BaseURL:       baseURL,
		Client:        NewClient(),
		DefaultRouter: NewRouter(),
		Resources:     make(map[string]*RestResource),
	}
}

func (a *API) AddResource(name string, resource *RestResource) {
	a.Resources[name] = resource
}

func (a *API) SetAuth(auth Authorization) {
	a.Client.Auth = auth
}

func (a *API) Call(name string, params map[string]string, body interface{}) error {
	resource, ok := a.Resources[name]
	if !ok {
		return fmt.Errorf("%s resource not found", name)
	}
	return a.Client.ProcessRequest(a.BaseURL, resource, params, body)
}
