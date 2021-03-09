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
	"net/http"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"bytes"


	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [name of todo] [descriotion(optional)]",
	Short: "Add new todo to the list",
	Long: `You can use "go-cli new" command to create and add todos to the list.
If your todo's name or description is more than one word use quotes.`,
	Args: func(cmd *cobra.Command, args []string) error {
    if len(args) < 1 {
      return errors.New("new command requires a name argument")
    }
		if len(args) > 2 {
      return errors.New("new command accepts maximum 2 arguments - name and description")
    }
		return nil
  },
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)==2{
			// fmt.Printf("Argument one: %v, argument two: %v\n", args[0], args[1])
			response := addNewTodo(args[0], args[1])
			log.Println(string(response))
		} else {
			// fmt.Printf("Argument one: %v, argument two: none\n", args[0])
			response := addNewTodo(args[0],"")
			log.Println(string(response))

		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addNewTodo(name string, description string) []byte {
	// fmt.Printf("Todo with name %v was %v created succesfully!\n", name,description)



	requestBody, err := json.Marshal(map[string] string{
		"name": name,
		"message": description,
	})

	if err!=nil {
		log.Fatalln(err)
	}
	API_URL := "http://golang-be:8080/todos"
	resp, err := http.Post(API_URL, "application/json", bytes.NewBuffer(requestBody))	

	if err!=nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		log.Fatalln(err)
	}

	return body
	
}
