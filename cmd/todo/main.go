package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/merkata/todo"
)

var todoFileName = ".todo.json"

func main() {
	listItems := flag.Bool("list", false, "list items")
	addItem := flag.String("add", "", "item to be included in the to-do")
	completeItem := flag.Int("complete", 0, "complete an item off the list")

	flag.Parse()

	if envFile := os.Getenv("TODO_CONFIG"); envFile != "" {
		todoFileName = envFile
	}

	l := &todo.List{}

	if _, err := os.Stat(todoFileName); os.IsNotExist(err) {
		f, err := os.Create(todoFileName)
		if err != nil {
			os.Exit(1)
		}
		f.Close()
	}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *listItems:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	case *addItem != "":
		l.Add(*addItem)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *completeItem > 0:
		l.Complete(*completeItem)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Println("invalid option specified")
		os.Exit(1)
	}
}
