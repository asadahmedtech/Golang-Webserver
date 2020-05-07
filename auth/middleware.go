package auth

import(
	"strings"
	"net/http"
	"context"
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"		
	"encoding/json"
	"errors"
	"project/models"

)

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := ExtractJWT(r)
		user, err := ValidateToken(token)
		if err != nil{
			// ReturnErrorJSON(w,  errors.New("Unexpected Error"))
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		ctxWithUser := context.WithValue(r.Context(), 0, user)
	    rWithUser := r.WithContext(ctxWithUser)

	    next(w, rWithUser)
	}
}

func ExtractJWT(r *http.Request) string{
	tokenStrings := r.Header.Get("Authorization")
	tokenString := strings.Split(tokenStrings, " ")[1]
	fmt.Println(tokenString)

	return tokenString
}

func ValidateToken(tokenString string) (string, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte("secret"), nil
	})
	if err != nil{
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["username"].(string))
		return claims["username"].(string), nil
	}
	return "", nil
}

func CookieMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		fmt.Println(c, err)
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				http.Redirect(w, r, "/login", 307) 
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			http.Redirect(w, r, "/login", 307) 

			return
		}

		// Get the JWT string from the cookie
		token := c.Value
		if token == ""{
			http.Redirect(w, r, "/login", 307) 
			return
		}
		user, err := ValidateToken(token)

		fmt.Println(user)
		if err != nil{
			// ReturnErrorJSON(w,  errors.New("Unexpected Error"))
			http.Error(w, "Forbidden", http.StatusForbidden)
			http.Redirect(w, r, "/login", 307) 
			return
		}

		ctxWithUser := context.WithValue(r.Context(), 0, user)
	    rWithUser := r.WithContext(ctxWithUser)

	    next(w, rWithUser)
	}
}
func ReturnJSONResp(w http.ResponseWriter, resp string, status int){
	var res models.ResponseResult

	w.Header().Set("Content-Type", "application/json")
	res.Status = status
	res.Result = resp
	json.NewEncoder(w).Encode(res)
	return
}
