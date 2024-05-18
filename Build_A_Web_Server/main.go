package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Todo struct represents a single todo item
type Todo struct {
	Task string
}

// TodoList represents the list of todo items
type TodoList struct {
	Todos []Todo
}

var todos TodoList

func main() {
	// Initialize an empty list of todos
	todos = TodoList{}

	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/delete", deleteHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the template file
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the todo list data
	err = tmpl.Execute(w, todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	r.ParseForm()
	task := r.Form.Get("task")

	// Add the new todo to the list
	todos.Todos = append(todos.Todos, Todo{Task: task})

	// Redirect back to the home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	r.ParseForm()
	index, err := strconv.Atoi(r.Form.Get("index"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Remove the todo at the specified index
	if index >= 0 && index < len(todos.Todos) {
		todos.Todos = append(todos.Todos[:index], todos.Todos[index+1:]...)
	}

	// Redirect back to the home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
