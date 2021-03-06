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
	"io/ioutil"
	"errors"
	"net/http"
	// "net/url"
	// "encoding/json"
	"log"
	"github.com/spf13/cobra"
)

var API_URL = "https://reqres.in/api/users/"

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list command lists all todos from the list and does not require any arguments.",
	Args: func(cmd *cobra.Command, args []string) error {
    if len(args) > 0 {
      return errors.New("list command does not require any arguments")
    }
		return nil
  },

	Run: func(cmd *cobra.Command, args []string) {
		listTodos()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listTodos(){
	// fmt.Println("List of todos:\n\n Name: todo\n Description: something here\n Status: not completed\n")

	resp, err := http.Get(API_URL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

}