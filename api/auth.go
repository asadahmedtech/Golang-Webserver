package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"project/models"
	"project/auth"
	// "errors"
	// "time"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// var res models.ResponseResult

	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)

	fmt.Println(user)
	if err != nil {
		auth.ReturnJSONResp(w, err.Error(), 400)
		return
	}

	err = auth.Register(user)
	if err != nil {
		auth.ReturnJSONResp(w, err.Error(), 400)
		return
	}

	auth.ReturnJSONResp(w, "Registeration Successful", 200)
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	result, err := auth.Login(user.Username, user.Password)

	if err != nil {
		auth.ReturnJSONResp(w, err.Error(), 400)
		return
	}
	
	json.NewEncoder(w).Encode(result)
	return 
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

	user := r.Context().Value(0).(string)
	result := auth.Profile(user)

	json.NewEncoder(w).Encode(result)
	return
}

func LoginHandlerStatic(w http.ResponseWriter, r *http.Request) {
	var user models.User

	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	result, err := auth.LoginStatic(user.Username, user.Password)

	if err != nil {
		auth.ReturnJSONResp(w, err.Error(), 400)
		return
	}
	
	json.NewEncoder(w).Encode(result)
	return 
}

func ProfileHandlerStatic(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

	user := r.Context().Value(0).(string)
	result := auth.ProfileStatic(user)

	result.Password = ""
	json.NewEncoder(w).Encode(result)
	return
}

