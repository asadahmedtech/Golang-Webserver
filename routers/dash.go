package routers

import(
	"net/http"
	"project/auth"
	"project/dash"
)

func DashRouter() http.Handler{
	router := http.NewServeMux()

	router.HandleFunc("/register", dash.RegisterHandler)
	router.HandleFunc("/login", dash.LoginPageHandler)
	router.HandleFunc("/logind", dash.LoginHandler)

	router.HandleFunc("/logout", dash.LogoutHandler)
	router.HandleFunc("/profile", auth.CookieMiddleware(dash.ProfileHandler))
	
	router.HandleFunc("/todopost", auth.CookieMiddleware(dash.TodoInsertHandler))
	router.HandleFunc("/todoview", auth.CookieMiddleware(dash.TodoFetchHandler))

	router.HandleFunc("/dostuff", auth.CookieMiddleware(dash.TaskHandler))
	
	return router
}