package cmd

import (
	"github.com/flynshue/tf-cloud/pkg/api"
	"github.com/spf13/viper"
)

var tfAPI *api.API

func TFAPI() *api.API {
	if tfAPI != nil {
		return tfAPI
	}
	tfAPI = api.NewAPI("https://app.terraform.io/api/v2/")
	tfAPI.SetAuth(api.AuthToken{Token: viper.GetString("token")})
	tfAPI.AddResource("listOrgs", listOrgResource())
	tfAPI.AddResource("getWorkspaces", getWkspaceResource())
	tfAPI.AddResource("createWkspace", createWkspaceResource())
	tfAPI.AddResource("searchWorkspace", searchWkspaceResource())
	tfAPI.AddResource("deleteWorkspace", deleteWkspaceResource())
	return tfAPI
}
