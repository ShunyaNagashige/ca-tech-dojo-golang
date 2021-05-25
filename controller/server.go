package controller

import (
	"net/http"
)

func createUserHandler(w http.ResponseWriter, r *http.Request){
	
}

func StartWebServer() {
	http.HandleFunc("/user/create", createUserHandler)
	http.ListenAndServe(":8080", nil)
}
