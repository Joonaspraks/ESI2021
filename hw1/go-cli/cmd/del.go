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

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del [id]",
	Short: "del command deletes todo with spesific id from list",
	Args: func(cmd *cobra.Command, args []string) error {
    if len(args) < 1 {
      return errors.New("del command requires and id argument")
    }
		if len(args) > 1 {
      return errors.New("dell command accepts maximum 1 argument - id")
    }
		return nil
  },
	Run: func(cmd *cobra.Command, args []string) {
		response := delTodo(args[0])
		log.Println(string(response))
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func delTodo(id string) []byte {
	// fmt.Printf("Todo  with id: %v was deleted succesfully!\n", id)
	API_URL:= "http://golang-be:8080/todos/"

	client := &http.Client{}
	patch_url := API_URL + string(id)

	
	req, err := http.NewRequest(http.MethodDelete, patch_url, nil)

	resp, err := client.Do(req)
	if err != nil {
			log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

	return body
}