package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAPICall(t *testing.T) {
	testApi := NewAPI("https://httpbin.org")
	router := NewRouter()
	router.AddFunc(200, func(resp *http.Response) error {
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	})
	getResource := NewRestResource("GET", "/get", router)
	anyResource := NewRestResource("GET", "/anything/{{.user}}", router)
	testApi.AddResource("get", getResource)
	testApi.AddResource("anything", anyResource)

	if err := testApi.Call("get", nil, nil); err != nil {
		t.Error(err)
	}
	err := testApi.Call("anything", map[string]string{"user": "johndoe"}, nil)
	if err != nil {
		t.Error(err)
	}
}
