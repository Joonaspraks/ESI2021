// DISCLAIMER Based upon "github.com/ulno/esi/testing-intro/article"  //

package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
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
	IsCompleted bool   `json:"isCompleted"`
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
	//
	IDmax := "0"

	for _, action := range r.actions {
		if action.ID > IDmax {
			IDmax = action.ID
			action.IsCompleted = true
		}
	}

	i, err := strconv.Atoi(IDmax)
	if err != nil { //can I avoid error handling?
		fmt.Println(err)
	}
	action.ID = strconv.Itoa(i + 1)
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
	json.NewEncoder(buf).Encode(r.actions)
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
