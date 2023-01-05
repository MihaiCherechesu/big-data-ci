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
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// pipelinesCmd represents the pipelines command
var pipelinesCmd = &cobra.Command{
	Use:   "pipelines",
	Short: "Lists the pipelines.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pipelines called")

		db := initConn()

		rows, err := db.Query("SELECT * FROM pipelines WHERE user_id = $1", "172.28.0.1")
		if err != nil {
			log.Fatalf("Error executing query: %q", err)
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var name string
			err = rows.Scan(&id, &name)
			if err != nil {
				log.Fatalf("Error scanning rows: %q", err)
			}
			fmt.Printf("ID: %d, Name: %s\n", id, name)
		}

		err = rows.Err()
		if err != nil {
			log.Fatalf("Error: %q", err)
		}
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
