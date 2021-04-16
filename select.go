package main

import (
	"os"
	"fmt"
	//"time"
)

func Select(c chan int, quits chan struct{}){
	//switch case
	// pero para canales y es async
	for{
		select {
		case <- c:
			fmt.Println("Received")
		case <- quits:
			fmt.Println("quit")
			os.Exit(0)
		}
	}
}

func main(){
	c := make(chan int, 2)
	quits := make(chan struct{})
	
	go func(){
		c <- 1
		quits <- struct{}{}
	}()
	
	go Select(c,quits) //corre en el background
	select {}

	fmt.Println()
}