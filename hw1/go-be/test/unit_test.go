package test

import (
	action "hw1/action"
	"testing"
)

var actions = []*action.Action{
	{
		ID:          "1",
		Message:     "Watch TV",
		Name:        "Mats",
		IsCompleted: false,
	},
	{
		ID:          "2",
		Message:     "Make dinner",
		Name:        "Uku",
		IsCompleted: true,
	},
	{
		ID:          "3",
		Message:     "Enjoy your weekend",
		Name:        "Ants",
		IsCompleted: false,
	},
}

var actionRepository = action.NewActionRepository(actions)

func TestAddNewTodo(t *testing.T) {
	repoActions := actionRepository.GetActions()
	startLength := len(repoActions)
	t.Log(startLength)
	actionRepository.AddNewAction(&action.Action{
		ID:          "999",
		Message:     "Test",
		Name:        "Test",
		IsCompleted: true,
	})

	repoActions = actionRepository.GetActions()
	endLength := len(actionRepository.GetActions())

	if endLength != startLength+1 {
		t.Error("New object was not added!")
	}

}

func TestCompleteTodo(t *testing.T) {
	id := "3"
	actionRepository.UpdateActionwithID(id)
	repoActions := actionRepository.GetActions()
	for _, action := range repoActions {
		if action.ID == id && action.IsCompleted == true {
			return
		}
	}
	t.Error("Todo was not updated!")
}
