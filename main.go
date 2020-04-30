package main

import (
    // "github.com/gorilla/mux"
	"net/http"
	"log"
	
	"project/routers"

)

func main(){

	v1router := routers.V1Router()
	v2router := routers.V2Router()
	dashrouter := routers.DashRouter()

	balancer := configBalancer()

	go runServer(v1router, "8001")
	go runServer(v2router, "8002")
	go func(){
		if err := balancer.ListenAndServe(); err != nil{
			log.Println(err)
		}
	}()
	runServer(dashrouter, "8000")
}

func runServer(router http.Handler, port string){
	if err := http.ListenAndServe(":"+port, router); err != nil{
		log.Println(err)
	}
}

func configBalancer() *http.Server{

	v1 := routers.NewProxy("http://127.0.0.1:8001")
	v2 := routers.NewProxy("http://127.0.0.1:8002")

	handler := routers.NewBalancer(v1, v2)

	server := &http.Server{
		Handler: handler,
		Addr: "127.0.0.1:9000",
	}
	return server
}
