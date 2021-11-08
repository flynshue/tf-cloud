package api

import "testing"

func TestRenderEndpoint(t *testing.T) {
	router := NewRouter()
	resource := NewRestResource("GET", "/user/{{.user}}", router)
	params := map[string]string{"user": "johndoe"}
	got := resource.RenderEndpoint(params)
	want := "/user/johndoe"
	if got != want {
		t.Errorf("resource.RenderEndpoint() = %s; wanted %s", got, want)
	}
}
