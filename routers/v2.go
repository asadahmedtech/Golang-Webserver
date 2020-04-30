package routers

import(
	"net/http"
	"project/api"
	"project/auth"
)

func V2Router() http.Handler{
	router := http.NewServeMux()

	router.HandleFunc("/api/auth/register", api.RegisterHandler)

	router.HandleFunc("/api/auth/login", api.LoginHandler)
	router.HandleFunc("/api/profile", auth.JWTMiddleware(api.ProfileHandler))

	router.HandleFunc("/api/todo/insert", auth.JWTMiddleware(api.TodoInsertHandler))
	router.HandleFunc("/api/todo/fetch", auth.JWTMiddleware(api.TodoFetchHandler))
	
	return router
}