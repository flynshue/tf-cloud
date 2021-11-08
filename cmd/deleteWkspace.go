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
	"io/ioutil"
	"net/http"

	"github.com/flynshue/tf-cloud/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteWkspaceCmd represents the delete command
var deleteWkspaceCmd = &cobra.Command{
	Use:   "delete",
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
		return deleteWkspace(viper.GetString("organization"), args[0])
	},
}

func deleteWkspace(tfOrg, name string) error {
	params := map[string]string{"organization": tfOrg, "workspace": name}
	return TFAPI().Call("deleteWorkspace", params, nil)
}

func deleteWkspaceDefault(resp *http.Response) error {
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s from %s\nresponse: %s\n", resp.Status, resp.Request.URL.Path, string(b))
	return nil
}

func deleteWkspaceSuccess(resp *http.Response) error {
	fmt.Printf("\n%d Successfully deleted %s\n", resp.StatusCode, resp.Request.URL.Path)
	return nil
}

func deleteWkspaceResource() *api.RestResource {
	router := api.NewRouter()
	router.DefaultRouter = deleteWkspaceDefault
	router.AddFunc(200, deleteWkspaceSuccess)
	resource := api.NewRestResource("DELETE", "/organizations/{{.organization}}/workspaces/{{.workspace}}", router)
	return resource
}

func init() {
	workspaceCmd.AddCommand(deleteWkspaceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteWkspaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteWkspaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
