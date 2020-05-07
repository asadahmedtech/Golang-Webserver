package dash

import(
	"fmt"
	"net/http"
	// "io/ioutil"
	// "encoding/json"
	"html/template"

	"project/models"
	// "project/auth"
	tododb "project/todo"
	"strings"
	"time"
	)

func TodoInsertHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo
	var res models.ResponseResult
	t, _  := template.ParseFiles("templates/todopost.html")

	r.ParseForm()
	task := r.FormValue("task")
	task = strings.TrimSpace(task)
	if task==""{
		res.Status = 400
		res.Result = "Enter something"
		t.Execute(w, res)
		return
	}
	todo.Task = task
	todo.Username = r.Context().Value(0).(string)
	todo.CreatedAt = time.Now()

	fmt.Println(todo)
	err := tododb.Insert(todo)
	
	if err != nil {
		res.Status = 400
		res.Result = err.Error()
		t.Execute(w, res)
		return
	}

	t.Execute(w, res)
	return
}

func TodoFetchHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(0).(string)
	result, err := tododb.Fetch(user)

	fmt.Println(result, err)
	t, _  := template.ParseFiles("templates/todoview.html")

	t.Execute(w, result)
	return
}