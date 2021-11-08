package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestProcessRequest(t *testing.T) {
	client := NewClient()
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
	resource := NewRestResource("GET", "/anything/{{.user}}", router)
	params := map[string]string{"user": "johndoe"}
	err := client.ProcessRequest("https://httpbin.org", resource, params, nil)
	if err != nil {
		t.Error(err)
	}
}
