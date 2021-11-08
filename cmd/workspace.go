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

	"github.com/spf13/cobra"
)

// workspaceCmd represents the workspace command
var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("workspace called")
	},
}

type tfworkspacePayload struct {
	Data tfworkspaceData `json:"data"`
}

type tfworkspaceData struct {
	ID         string                `json:"id,omitempty"`
	Type       string                `json:"type"`
	Attributes tfworkspaceAttributes `json:"attributes"`
}

type getWkspaceResponse struct {
	Data []tfworkspaceData `json:"data"`
}

type tfworkspaceAttributes struct {
	Created       string `json:"created-at,omitempty"`
	ExecutionMode string `json:"execution-mode"`
	Name          string `json:"name"`
	Locked        bool   `json:"locked,omitempty"`
}

func init() {
	rootCmd.AddCommand(workspaceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workspaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workspaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
