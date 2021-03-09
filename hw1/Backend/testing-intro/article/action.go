// DISCLAIMER Based upon "github.com/ulno/esi/testing-intro/article"  //

package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
)

// Repository to store actions
type Repository struct {
	actions []*Action
}

// NewActionRepository returns action repository
func NewActionRepository(actions []*Action) *Repository {
	return &Repository{
		actions: actions,
	}
}

// Action ...
type Action struct {
	ID          string `json:"ID"`
	Message     string `json:"message"`
	Name        string `json:"name"`
	IsCompleted bool   `json:"link"`
}

// GenSingleAction returns all actions matching the given id
func (r *Repository) GenSingleAction(ID string) []byte {
	buf := &bytes.Buffer{}
	for _, action := range r.actions {
		if action.ID == ID {

			json.NewEncoder(buf).Encode(action)
		}
	}
	return buf.Bytes()
}

// AddNewAction add an action to the internal actions list
func (r *Repository) AddNewAction(action *Action) {
	r.actions = append(r.actions, action)
}

// DeleteAction deletes all actions that have the given id from teh internal actions list
func (r *Repository) DeleteActionWithID(ID string) {
	for index, action := range r.actions {
		if action.ID == ID {
			r.actions = append(r.actions[:index], r.actions[index+1:]...)
		}
	}
}

// genAllActions returns a json list of all actions in the internal action list
func (r *Repository) GenAllActions() []byte {

	buf := &bytes.Buffer{}
	json.NewEncoder(buf).Encode(
		sortTodosByCaps(r.actions))
	return buf.Bytes()
}

// GenSingleAction returns all actions matching the given id
func (r *Repository) UpdateActionwithID(ID string) []byte {
	buf := &bytes.Buffer{}
	for _, action := range r.actions {
		if action.ID == ID {
			action.IsCompleted = true
		}
	}
	return buf.Bytes()
}

func checkStringAlphabet(str string) int {
	a := 0
	for _, charVariable := range str {
		if charVariable >= 'A' && charVariable <= 'Z' {
			a++
		}
	}
	return a
}

func sortTodosByCaps(actions []*Action) []*Action {
	sort.Slice(actions, func(i, j int) bool {
		return checkStringAlphabet(actions[i].Message) > checkStringAlphabet(actions[j].Message)
	})

	for _, action := range actions {
		fmt.Printf("%+v", action)
	}

	return actions
}

func testChecker() {
	fmt.Println(checkStringAlphabet("ActioZ"))
	fmt.Println(checkStringAlphabet("tegeVus"))
	fmt.Println(checkStringAlphabet("HELLLOOO!!"))
}
