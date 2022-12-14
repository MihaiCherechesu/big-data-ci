/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// pipelinesCmd represents the pipelines command
var pipelinesCmd = &cobra.Command{
	Use:   "pipelines",
	Short: "Lists the pipelines.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pipelines called")
		resp, err := http.Get("http://localhost:8081/pipelines/")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var indentedBody bytes.Buffer
		err = json.Indent(&indentedBody, []byte(string(body)), "", "\t")
		if err != nil {
			log.Fatalf("could not indent body json, %v", err)
		}

		log.Printf("%s\n", &indentedBody)
	},
}

func init() {
	lsCmd.AddCommand(pipelinesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pipelinesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pipelinesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
