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
	"time"
	)

func TodoInsertHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo
	var res models.ResponseResult

	r.ParseForm()
	todo.Task = r.FormValue("task")

	todo.Username = r.Context().Value(0).(string)
	todo.CreatedAt = time.Now()

	fmt.Println(todo)
	res, err := tododb.Insert(todo)

	t, _  := template.ParseFiles("templates/todopost.html")
	if err != nil {
		res.Error = err.Error()
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