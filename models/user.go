package models

import(	
	jwt "github.com/dgrijalva/jwt-go"

)
type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}