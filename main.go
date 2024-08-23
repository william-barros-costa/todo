package main

import (
	"fmt"
	"os"
)

// Uses docopt format
const help_message string = `CLI tool used to keep track of to do list

Usage: todo COMMAND [OPTIONS]

Commands:
  add            Adds a task to the todo list
  list           Lists tasks in the todo list
  delete         Deletes a task from the todo list
  edit           Edits a task
  complete       Marks a task as completed 

Global Options:
  -h --help      Show this screen.
  --version      Show version.

Run 'docker COMMAND --help' for more information on a command.`

func main() {
	if len(os.Args) == 1 {
		print_help()
		os.Exit(0)
	}
	fmt.Println(os.Args)
}

func print_help() {
	fmt.Fprintln(os.Stdout, []any{help_message}...)
}

/*
Add Task
List Tasks
Remove Task
Edit Task
Mark completed
List Completed tasks
*/
