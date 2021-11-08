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
	"github.com/spf13/viper"
)

// createWkspaceCmd represents the create command
var createWkspaceCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("must supply workspace name")
		}
		return createWkspace(args[0], viper.GetString("organization"))
	},
}

func newTfworkspace(name string) *tfworkspacePayload {
	return &tfworkspacePayload{Data: tfworkspaceData{
		Type: "workspaces",
		Attributes: tfworkspaceAttributes{
			Name:          name,
			ExecutionMode: "local",
		},
	}}
}

func createWkspace(name, tfOrg string) error {
	workspace := newTfworkspace(name)
	params := map[string]string{"organization": tfOrg}
	return TFAPI().Call("createWkspace", params, workspace)
}

func createWkspaceSuccess(resp *http.Response) error {
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	wkspace := &tfworkspacePayload{}
	if err := json.Unmarshal(b, wkspace); err != nil {
		return err
	}
	fmt.Printf("\nSuccessfully created workspace\n")
	fmt.Printf("id: %s, name: %s, created: %s, locked: %t, execution-mode: %s\n", wkspace.Data.ID,
		wkspace.Data.Attributes.Name,
		wkspace.Data.Attributes.Created,
		wkspace.Data.Attributes.Locked,
		wkspace.Data.Attributes.ExecutionMode)
	return nil
}

func createWkspaceDefaultRouter(resp *http.Response) error {
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return fmt.Errorf("%s from %s\nresponse: %s", resp.Status, resp.Request.URL.Path, string(b))
}

func createWkspaceResource() *api.RestResource {
	router := api.NewRouter()
	router.AddFunc(201, createWkspaceSuccess)
	router.DefaultRouter = createWkspaceDefaultRouter
	resource := api.NewRestResource("POST", "/organizations/{{.organization}}/workspaces", router)
	return resource
}

func init() {
	workspaceCmd.AddCommand(createWkspaceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createWkspaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createWkspaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
