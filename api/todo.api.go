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

	// "errors"
	)

func TodoInsertHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo
	// var res models.ResponseResult

	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &todo)

	todo.Username = r.Context().Value(0).(string)
	todo.CreatedAt = time.Now()

	fmt.Println(todo)
	if err != nil {
		auth.ReturnJSONResp(w, err.Error(), 400)
		return
	}

	err = tododb.Insert(todo)

	if err != nil {
		auth.ReturnJSONResp(w, err.Error(), 400)
		return
	}
	auth.ReturnJSONResp(w, "Todo Inserted", 200)
	return
}

func TodoFetchHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(0).(string)
	res, err := tododb.Fetch(user)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		auth.ReturnJSONResp(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(res)
	return
}

func TodoInsertHandlerStatic(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo
	// var res models.ResponseResult

	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &todo)

	todo.Username = r.Context().Value(0).(string)
	todo.CreatedAt = time.Now()

	fmt.Println(todo)
	if err != nil {
		auth.ReturnJSONResp(w, err.Error(), 400)
		return
	}


	// res.Result = "Success"
	err = nil

	if err != nil {
		auth.ReturnJSONResp(w, err.Error(), 400)
		return
	}
	auth.ReturnJSONResp(w, "Todo Inserted", 200)
	return
}

func TodoFetchHandlerStatic(w http.ResponseWriter, r *http.Request) {
	res := models.ToDo{
		Task:"Oh okay static thing",
		CreatedAt: time.Now(),
	}
	// err er

	// if err != nil {
	// 	auth.ReturnJSONResp(w, err.Error(), 400)
	// 	return
	// }
	json.NewEncoder(w).Encode(res)
	return
}