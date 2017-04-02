package data

import (
	"time"
)

type Task struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Desc         string    `json:"desc"`
	DueDate      time.Time `json:"dueDate"`
	RelatedTasks *[]int    `json:"relatedTasks"`
	Completed    bool      `json:"completed"`

	creationDate time.Time `json:"creationDate"`
	valid        bool      `json:"valid"`
}

func CreateNewTask(id int, dueDate time.Time, relatedTasks *[]int, name, desc string) *Task {
	t := new(Task)
	t.creationDate = time.Now().Local()
	t.Id = id
	t.Name = name
	t.Desc = desc
	t.RelatedTasks = relatedTasks
	t.Completed = false
	t.DueDate = dueDate
	if dueDate.Before(t.creationDate) {
		t.valid = true
	} else {
		t.valid = false
	}
	return t
}

//used to get the creation date of the task
func (t *Task) CreationDate() time.Time {
	return t.creationDate
}

// used to see if the task is visible
func (t *Task) Valid() bool {
	return t.valid
}

// function to add an id to the related task
// of a task
func (t *Task) AddRelatedTask(id int) {
	length := len(*t.RelatedTasks)
	capacity := cap(*t.RelatedTasks)
	//dont reply on append to increase capacity
	//we always double when full
	if length == capacity {
		newList := make([]int, length, (capacity+1)*2)
		copy(newList, *t.RelatedTasks)
		t.RelatedTasks = &newList
	}
	//append data
	(*t.RelatedTasks)[length] = id
}
