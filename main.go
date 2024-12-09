package main

import (
	"encoding/json"
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

var commands map[string]func([]string, *todo) = map[string]func([]string, *todo){
	"add":    function_add,
	"list":   function_list,
	"delete": function_delete,
	// "edit":     function_edit,
	// "complete": function_complete,
	// "help":     function_help,
}

type task struct {
	Name      string
	Completed bool
}

type todo struct {
	Tasklist []task
}

func (t *todo) addTask(name string) {
	newTask := task{Name: name, Completed: false}
	t.Tasklist = append(t.Tasklist, newTask)
}

func (t *todo) toJson() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}

func (t todo) toString() string {
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

func saveFile(location string, t *todo) {
	if location == "" {
		location = "tasks.json"
	}
	if bytes, error := t.toJson(); error == nil {
		os.WriteFile(location, bytes, 0666)
	}
}

func importTasks(location string) todo {
	if location == "" {
		location = "tasks.json"
	}

	var jsonBytes []byte
	var err error

	if jsonBytes, err = os.ReadFile(location); err != nil {
		return todo{}
	}

	var importedTasks todo

	if err = json.Unmarshal(jsonBytes, &importedTasks); err != nil {
		return todo{}
	}
	return importedTasks
}

func main() {
	if len(os.Args) == 1 {
		print_help()
		return
	}

	todo := importTasks("")

	if command, ok := commands[os.Args[1]]; ok {
		command(os.Args[2:], &todo)
	} else {
		print_help()
	}
	function_list(nil, &todo)
	saveFile("", &todo)
}

func print_help() {
	fmt.Fprintln(os.Stdout, []any{help_message}...)
}

func function_list(args []string, t *todo) {
	fmt.Println(t.toString())
}

func function_add(args []string, t *todo) {
	for _, task := range args {
		t.addTask(task)
	}
}

func function_delete(args []string, t *todo) {
	newTasks := make([]task, 0)
	for _, arg := range args {
		for _, task := range t.Tasklist {
			if task.Name != arg {
				newTasks = append(newTasks, task)
			}
		}
	}
	t.Tasklist = newTasks
}

/*
Edit Task
Mark completed
List Completed tasks
*/
