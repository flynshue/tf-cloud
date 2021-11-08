package api

import "fmt"

type Authorization interface {
	AuthorizationHeader() string
}

type AuthToken struct {
	Token string
}

func (at AuthToken) AuthorizationHeader() string {
	return fmt.Sprintf("Bearer %s", at.Token)
}
