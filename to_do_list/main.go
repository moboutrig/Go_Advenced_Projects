package main

import (
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

type Todo struct {
	Item   string
	Status string
}

var (
	todos []Todo
	tmpl  *template.Template
	mutex sync.Mutex
)

func init() {
	tmpl = template.Must(template.ParseFiles("templates/template.html"))
	todos = []Todo{
		{Item: "js piscine ", Status: "Not Completed"},
		{Item: "ai piscine ", Status: "Not Completed"},
		{Item: "ai daily doose", Status: "Not Completed"},
		{Item: "ai project", Status: "Not Completed"},
		{Item: "zone project ", Status: "Not Completed"},
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/complete", completeHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	tmpl.Execute(w, todos)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		item := r.FormValue("item")
		if item != "" {
			mutex.Lock()
			todos = append(todos, Todo{Item: item, Status: "Not Completed"})
			mutex.Unlock()
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func completeHandler(w http.ResponseWriter, r *http.Request) {
	index := r.URL.Query().Get("index")
	if index != "" {
		i := atoi(index)
		if i >= 0 && i < len(todos) {
			mutex.Lock()
			todos[i].Status = "Completed"
			mutex.Unlock()
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	index := r.URL.Query().Get("index")
	if index != "" {
		i := atoi(index)
		if i >= 0 && i < len(todos) {
			mutex.Lock()
			todos = append(todos[:i], todos[i+1:]...)
			mutex.Unlock()
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}
