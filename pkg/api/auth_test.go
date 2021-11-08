package api

import "testing"

func TestAuth(t *testing.T) {
	authToken := AuthToken{Token: "secretToken"}
	want := "Bearer secretToken"
	got := authToken.AuthorizationHeader()
	if got != want {
		t.Errorf("authToken.AuthorizationHeader() = %s; wanted %s", got, want)
	}
}
