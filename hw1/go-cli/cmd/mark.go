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
	// "fmt"
	"net/http"
	"errors"
	"io/ioutil"
	"log"
	"github.com/spf13/cobra"
)

// markCmd represents the mark command
var markCmd = &cobra.Command{
	Use:   "mark",
	Short: "mark command will mark todo with specified id as completed",
	Args: func(cmd *cobra.Command, args []string) error {
    if len(args) < 1 {
      return errors.New("mark command requires and id argument")
    }
		if len(args) > 1 {
      return errors.New("mark command accepts maximum 1 argument - id")
    }
		return nil
  },
	Run: func(cmd *cobra.Command, args []string) {
		markTodo(args[0])
	},
}

func init() {
	rootCmd.AddCommand(markCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func markTodo(id string) {
	// fmt.Printf("Todo with id: %v was marked as completed\n", id)

	client := &http.Client{}
	patch_url := API_URL + string(id)

	
	req, err := http.NewRequest(http.MethodPatch, patch_url, nil)

	resp, err := client.Do(req)
	if err != nil {
			log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
	log.Println(string(body))
}