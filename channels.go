package main

import (
	"fmt"
	"time")

func main(){
	canal := make(chan int ) //unbuffered o tamaño buffer 1/0
	// send 
	go func(){
		canal <-3 // pone 3 en el canal
	}()
	//sniff
	val := <- canal
	fmt.Println(val)
	go func(){
		canal <-5 // pone 3 en el canal
	}()
	time.Sleep(time.Second * 2)
	val = <- canal
	fmt.Println(val)
	
}