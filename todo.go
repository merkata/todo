//Package todo implements a simple To Do library
/*
API
type item struct
type List []item
func (l *List) Add(task string)
func (l *List) Complete(task string) error
func (l *List) Save(filename string) error
func (l *List) Get(filename string) error
*/
package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List defines all items and exposes
// a well defined API over the protected item(s)
type List []item

// Add creates a new item and appends it to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

// Complete sets an item as done and the CompletedAt time
// to the time (time.Now) the call was invoked
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}

// Save tries to write the list of items as JSON in a filename
func (l *List) Save(filename string) error {
	content, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, content, 0644)
}

// Get tries to get a list of items as JSON from a filename
func (l *List) Get(filename string) error {
	fi, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if len(fi) == 0 {
		return nil
	}
	return json.Unmarshal(fi, l)
}
