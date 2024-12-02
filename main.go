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

Run 'todo COMMAND help' for more information on a command.`

const add_help_message string = `Usage: todo add <number> <task_1> ... <task_number>`

var commands map[string]func([]string, tasks) = map[string]func([]string, tasks){
	"add": function_add,
	"list":     function_list,
	// "delete":   function_delete,
	// "edit":     function_edit,
	// "complete": function_complete,
	// "help":     function_help,
}

type task struct {
	name      string
	completed bool
}

type tasks struct {
	tasklist []task
}

func (t tasks) addTask(name string) {
	newTask := task{name: name, completed: false}
	t.tasklist = append(t.tasklist, newTask)
}

func (t tasks) toString() string {
	var representation string
	for _, task := range t.tasklist {
		if task.completed {
			representation += " [X] "
		} else {
			representation += " [X] "
		}
		representation += task.name
		representation += "\n"
	}
	representation += "\n"
	return representation
}

func main() {
	if len(os.Args) == 1 {
		print_help()
		return
	}

	tasks := tasks{}

	if command, ok := commands[os.Args[1]]; ok {
		command(os.Args[2:], tasks)
	} else {
		print_help()
	}
}

func print_help() {
	fmt.Fprintln(os.Stdout, []any{help_message}...)
}

func function_list(args []string, t tasks)  {
	fmt.Println(tasks.toString())
}

func function_add(args []string, t tasks) {
	if args[0] == "help" {
		fmt.Println(add_help_message)
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
