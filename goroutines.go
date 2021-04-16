package main

import (
	"time"
	"fmt"
)

func heavy() {
	for {
		time.Sleep(time.Second*1)
		fmt.Println("heavy")
	}
}
func superheavy() {
	for {
		time.Sleep(time.Second*2)
		fmt.Println("super heavy")
	}
}
func main(){
	go heavy() //as go routine, corre en el background
	go superheavy()
	fmt.Println("Fin")
	time.Sleep(time.Second*5)
}