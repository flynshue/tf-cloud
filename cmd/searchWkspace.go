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
	"fmt"

	"github.com/flynshue/tf-cloud/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// searchWkspaceCmd represents the searchWkspace command
var searchWkspaceCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("must supply search keyword")
		}
		return searchWkspace(args[0], viper.GetString("organization"))
	},
}

func searchWkspace(name, tfOrg string) error {
	params := map[string]string{"organization": tfOrg, "workspace": name}
	return TFAPI().Call("searchWorkspace", params, nil)
}

func searchWkspaceResource() *api.RestResource {
	router := api.NewRouter()
	router.AddFunc(200, getWkspaceSuccess)
	resource := api.NewRestResource("GET", "/organizations/{{.organization}}/workspaces?search[name]={{.workspace}}", router)
	return resource
}

func init() {
	workspaceCmd.AddCommand(searchWkspaceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchWkspaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchWkspaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
