package api

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"project/models"
	"project/auth"
	tododb "project/todo"
	"time"
	)

func TodoInsertHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo
	var res models.ResponseResult

	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &todo)

	todo.Username = r.Context().Value(0).(string)
	todo.CreatedAt = time.Now()

	fmt.Println(todo)
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res, err = tododb.Insert(todo)

	if err != nil {
		auth.ReturnErrorJSON(w, err)
		return
	}
	json.NewEncoder(w).Encode(res)
	return
}

func TodoFetchHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(0).(string)
	res, err := tododb.Fetch(user)

	if err != nil {
		auth.ReturnErrorJSON(w, err)
		return
	}
	json.NewEncoder(w).Encode(res)
	return
}

func TodoInsertHandlerStatic(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo
	var res models.ResponseResult

	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &todo)

	todo.Username = r.Context().Value(0).(string)
	todo.CreatedAt = time.Now()

	fmt.Println(todo)
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}


	res.Result = "Success"
	err = nil

	if err != nil {
		auth.ReturnErrorJSON(w, err)
		return
	}
	json.NewEncoder(w).Encode(res)
	return
}

func TodoFetchHandlerStatic(w http.ResponseWriter, r *http.Request) {
	res := models.ToDo{
		Task:"Oh okay static thing",
		CreatedAt: time.Now(),
	}
	// err er

	// if err != nil {
	// 	auth.ReturnErrorJSON(w, err)
	// 	return
	// }
	json.NewEncoder(w).Encode(res)
	return
}