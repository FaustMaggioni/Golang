package main

import (
	"fmt"
	"net/http")

func main(){
	mux := http.NewServeMux() //router
	mux.HandleFunc("/getgoing/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received")
		fmt.Println(r.Method)
		w.Write([] byte("Hello Asshole "))
		w.Write([] byte("EPICO"))
	})
	http.ListenAndServe("localhost:3000", mux)  //(address,handler) creando server escuchando en localhost 

}