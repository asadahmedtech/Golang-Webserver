package routers

import(
	"net/http"
	"project/api"
	"project/auth"
)

func V1Router() http.Handler{
	router := http.NewServeMux()

	router.HandleFunc("/api/auth/register", api.RegisterHandler)

	router.HandleFunc("/api/auth/login", api.LoginHandlerStatic)
	router.HandleFunc("/api/profile", auth.JWTMiddleware(api.ProfileHandlerStatic))

	router.HandleFunc("/api/todo/insert", auth.JWTMiddleware(api.TodoInsertHandlerStatic))
	router.HandleFunc("/api/todo/fetch", auth.JWTMiddleware(api.TodoFetchHandlerStatic))
	
	return router
}