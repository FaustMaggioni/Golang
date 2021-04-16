package main

import (
    //"fmt"
	"net/http"
	"encoding/json"
	"github.com/FaustMaggioni/Golang/tree/main/API_Intro/API_todo/structs"
    )

// JSON API



func main(){
	const ping="/ping"
	mux := http.NewServeMux()
	mux.HandleFunc(ping,func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			data:= structs.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			
			JSON.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe(":3000", nil)
}