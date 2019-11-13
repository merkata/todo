package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"testing"
)

const (
	todoFileName = ".todo.json"
	binName      = "todo"
)

func TestMain(m *testing.M) {
	if _, err := os.Stat(binName); os.IsNotExist(err) {
		build := exec.Command("go", "build", "-o", binName)
		if err := build.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Cannot build %s: %s", binName, err)
			os.Exit(1)
		}
	}

	result := m.Run()

	//	os.Remove(binName)
	//	os.Remove(todoFileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "testTask"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := path.Join(dir, binName)

	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := task + "\n"

		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead", expected, string(out))
		}

	})
}
