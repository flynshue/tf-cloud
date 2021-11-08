package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type Client struct {
	Client *http.Client
	Auth   Authorization
}

func NewClient() *Client {
	return &Client{Client: http.DefaultClient}
}

func (c *Client) ProcessRequest(baseURL string, resource *RestResource, params map[string]string, body interface{}) error {
	endpoint := strings.TrimLeft(resource.RenderEndpoint(params), "/")
	baseURL = strings.TrimRight(baseURL, "/")
	url := baseURL + "/" + endpoint
	req, err := buildRequest(resource.Method, url, body)
	if err != nil {
		return err
	}
	if c.Auth != nil {
		req.Header.Set("Authorization", c.Auth.AuthorizationHeader())
	}
	req.Header.Set("Content-Type", "application/vnd.api+json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	return resource.Router.CallFunc(resp)
}

func buildRequest(method, url string, body interface{}) (*http.Request, error) {
	if body == nil {
		return http.NewRequest(method, url, nil)
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)
	return http.NewRequest(method, url, buf)
}
