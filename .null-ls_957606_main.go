package main

import (
	"fmt"
	"os"
	"strconv"
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

Run 'todo COMMAND help' for more information on a command.`

const add_help_message string = `Usage: todo add <number> <task_1> ... <task_number>`

var commands [6]string = [6]string{"add", "list", "delete", "edit", "complete", "help"}

func main() {
	if len(os.Args) == 1 {
		print_help()
		return
	}
	var command string = os.Args[1]
	if is_valid_command(command) {
		switch command {
		case "add":
			println("activated add")
			break
		}
	} else {
		print_help()
	}
}

func print_help() {
	fmt.Fprintln(os.Stdout, []any{help_message}...)
}

func is_valid_command(command string) bool {
	for _, c := range commands {
		if command == c {
			return true
		}
	}
	return false
}

func add(args []string) {
	int, error := strconv.Atoi(args[0])
	if error != nil {
		print_help()
		return
	}
}

/*
Add Task
List Tasks
Remove Task
Edit Task
Mark completed
List Completed tasks
*/
