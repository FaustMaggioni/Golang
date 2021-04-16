package main

import (
    //"fmt"
	"net/http"
	"encoding/json"
	"./structs"
    )

// JSON API


func main(){
	const ping="/ping"
	mux := http.NewServeMux()
	mux.HandleFunc(ping,func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			data:= views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			
			JSON.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe(":3000", mux)
}