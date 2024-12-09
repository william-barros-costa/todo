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

var commands map[string]func([]string, *tasks) = map[string]func([]string, *tasks){
	"add":  function_add,
	"list": function_list,
	// "delete":   function_delete,
	// "edit":     function_edit,
	// "complete": function_complete,
	// "help":     function_help,
}

type task struct {
	Name      string
	Completed bool
}

type tasks struct {
	Tasklist []task
}

func (t *tasks) addTask(name string) {
	newTask := task{Name: name, Completed: false}
	t.Tasklist = append(t.Tasklist, newTask)
}

func (t tasks) toString() string {
	var representation string
	for _, task := range t.Tasklist {
		if task.Completed {
			representation += " [X] "
		} else {
			representation += " [] "
		}
		representation += task.Name
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
		command(os.Args[2:], &tasks)
	} else {
		print_help()
	}
	function_list(nil, &tasks)
}

func print_help() {
	fmt.Fprintln(os.Stdout, []any{help_message}...)
}

func function_list(args []string, t *tasks) {
	fmt.Println(t.toString())
}

func function_add(args []string, t *tasks) {
	for _, task := range args {
		t.addTask(task)
	}
}

func function_delete(args []string, t *tasks) {
	newTasks := make([]task, 0)
	for _, arg := range args {
		for _, task := range t.Tasklist {
			if task.Name != arg {
				newTasks = append(newTasks, task)
			}
		}
	}
}

/*
Remove Task
Edit Task
Mark completed
List Completed tasks
*/
