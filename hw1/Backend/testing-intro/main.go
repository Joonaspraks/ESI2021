// Example based on Rest and unit testing in https://golangdocs.com/

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	action "github.com/ulno/esi/testing-intro/article"

	"github.com/gorilla/mux"
)

var actions = []*action.Action{
	{
		ID:          "1",
		Message:     "Watch TV",
		Name:        "Mats",
		IsCompleted: "0",
	},
	{
		ID:          "2",
		Message:     "Make dinner",
		Name:        "Uku",
		IsCompleted: "1",
	},
	{
		ID:          "3",
		Message:     "Enjoy your weekend",
		Name:        "Ants",
		IsCompleted: "0",
	},
}

var actionRepository = action.NewActionRepository(actions)

const endPointHit = "Endpoint Hit:"

// genHomePage returns the content of the home page
func genHomePage() []byte {
	return []byte("Welocme to Group 11 todo-actions webpage!!!")
}

func homePage(w http.ResponseWriter, _ *http.Request) {
	log.Println(endPointHit, "home page")
	w.Write(genHomePage())
}

func returnSingleAction(w http.ResponseWriter, r *http.Request) {
	log.Println(endPointHit, "return single action")
	vars := mux.Vars(r)
	key := vars["id"]

	w.Write(actionRepository.GenSingleAction(key))
}

func createNewAction(w http.ResponseWriter, r *http.Request) {
	log.Println(endPointHit, "create new action")
	reqBody, _ := ioutil.ReadAll(r.Body)
	action := &action.Action{}
	json.Unmarshal(reqBody, action)
	actionRepository.AddNewAction(action)

	json.NewEncoder(w).Encode(action)
}

func deleteAction(w http.ResponseWriter, r *http.Request) {
	log.Println(endPointHit, "delete action")
	vars := mux.Vars(r)
	id := vars["id"]

	actionRepository.DeleteActionWithID(id)
}

func returnAllActions(w http.ResponseWriter, r *http.Request) {
	log.Println(endPointHit, "return all actions")
	w.Write(actionRepository.GenAllActions())
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods(http.MethodGet)
	myRouter.HandleFunc("/todos", returnAllActions).Methods(http.MethodGet)
	myRouter.HandleFunc("/todos", createNewAction).Methods(http.MethodPost)
	myRouter.HandleFunc("/todo/{id}", deleteAction).Methods(http.MethodDelete)
	myRouter.HandleFunc("/todo/{id}", returnSingleAction).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
