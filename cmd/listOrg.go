/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flynshue/tf-cloud/pkg/api"
	"github.com/spf13/cobra"
)

// listOrgCmd represents the list command
var listOrgCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return listTFOrgs()
	},
}

type listOrgSuccessResponse struct {
	Data []tfOrganizations `json:"data"`
}

type tfOrganizations struct {
	Attributes struct {
		ExternalID string `json:"external-id"`
		Created    string `json:"created-at"`
		Name       string `json:"name"`
	} `json:"attributes"`
}

func listTFOrgs() error {
	return TFAPI().Call("listOrgs", nil, nil)
}

func listOrgSuccess(resp *http.Response) error {
	defer resp.Body.Close()
	response := &listOrgSuccessResponse{}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, response); err != nil {
		return err
	}
	for _, tfOrg := range response.Data {
		fmt.Println("Name:", tfOrg.Attributes.Name)
		fmt.Println("External ID:", tfOrg.Attributes.ExternalID)
		fmt.Println("Created:", tfOrg.Attributes.Created)
	}
	return nil
}

func listOrgResource() *api.RestResource {
	router := api.NewRouter()
	router.AddFunc(200, listOrgSuccess)
	resource := api.NewRestResource("GET", "/organizations", router)
	return resource
}

func init() {
	organizationsCmd.AddCommand(listOrgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listOrgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listOrgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
