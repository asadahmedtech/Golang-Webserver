package dash

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	"net/http"
	"time"
	"html/template"
	"project/models"
	"project/auth"

	// jwt "github.com/dgrijalva/jwt-go"

)

func LoadFile(fileName string) (string, error) {
    bytes, err := ioutil.ReadFile(fileName)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	t, _  := LoadFile("templates/register.html")
	fmt.Fprintf(w, t)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var res models.ResponseResult

	r.ParseForm()
	user.Username = r.FormValue("username")
	user.FirstName = r.FormValue("firstname")
	user.LastName = r.FormValue("lastname")
	user.Password = r.FormValue("password")

	err := auth.Register(user)

	t, _  := template.ParseFiles("templates/register.html")
	if err != nil {
		res.Result = err.Error()
		t.Execute(w, res)
		return
	}
	res.Result = "Registration Successful"
	t.Execute(w, res)
	return
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	t, _  := LoadFile("templates/login.html")
	fmt.Fprintf(w, t)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// var user models.User
	var res models.ResponseResult

	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	result, err := auth.Login(username, password)

	fmt.Println(username, password, result, err)
	if err != nil {
		t, _  := template.ParseFiles("templates/login.html")
		res.Result = err.Error()
		t.Execute(w, res)
		return
	}

	tokenString := result.Token
	expiration := time.Now().Add(365 * 24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expiration,
	})
	fmt.Println(tokenString)
	c, err := r.Cookie("token")
	fmt.Println(c, err)

	http.Redirect(w, r, "/profile", 307) 
	return
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	expiration := time.Now().Add(365 * 24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   "",
				Expires: expiration,
	})

	http.Redirect(w, r, "/login", 307) 
	return
}
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// c, err := r.Cookie("token")
	// fmt.Println(c, err)
	// if err != nil {
	// 	if err == http.ErrNoCookie {
	// 		// If the cookie is not set, return an unauthorized status
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}
	// 	// For any other type of error, return a bad request status
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// // Get the JWT string from the cookie
	// tknStr := c.Value

	// // Initialize a new instance of `Claims`
	// claims := &models.Claims{}

	// // Parse the JWT string and store the result in `claims`.
	// // Note that we are passing the key in this method as well. This method will return an error
	// // if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// // or if the signature does not match
	// tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte("secret"), nil
	// })
	// if err != nil {
	// 	if err == jwt.ErrSignatureInvalid {
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// if !tkn.Valid {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }
	// user := claims.Username


	user := r.Context().Value(0).(string)
	result := auth.Profile(user)

	t, _  := template.ParseFiles("templates/profile.html")
	t.Execute(w, result)

}
