package todo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/merkata/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "Test Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("wrong task name, expected %q, got %q instead.", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "Test Task"
	l.Add(taskName)
	if l[0].Done {
		t.Errorf("task cannot be completed without invoking Complete on it first.")
	}
	l.Complete(1)
	if !l[0].Done {
		t.Errorf("task has to be marked as completed.")
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "Test Task"

	l1.Add(taskName)
	fi, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("error creating temp file %s\n", err)
	}
	defer os.Remove(fi.Name())

	if err := l1.Save(fi.Name()); err != nil {
		t.Errorf("could not save items list to temp file")
	}

	if err := l2.Get(fi.Name()); err != nil {
		t.Errorf("could not get items list from temp file")
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("items from saved and retrieved items list don't match")
	}

}
