package api

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestCallFunc(t *testing.T) {
	router := NewRouter()
	router.AddFunc(200, func(resp *http.Response) error {
		fmt.Printf("%d from %s\n", resp.StatusCode, resp.Request.URL.Path)
		return nil
	})
	resp := &http.Response{
		StatusCode: 200,
		Request:    &http.Request{URL: &url.URL{Path: "/Get"}},
	}
	if err := router.CallFunc(resp); err != nil {
		t.Error(err)
	}
}
