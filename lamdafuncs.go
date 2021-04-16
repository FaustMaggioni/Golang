package main

import (
	"fmt"
	"sync"
)

func main(){
	//wait groups
	wg := &sync.WaitGroup{}
	//add, done, wait funcionallity
	//mutex
	mut := &sync.Mutex()
	// mut.Lock() mut.Unlock()
	wg.Add(2) //quiere decir que esperará a 2 subrutinas
	go func (){ //función lamda
		fmt.Println("Hello")
		wg.Done()
	}()
	go func (){ //función lamda
		fmt.Println("world")
		wg.Done()
	}()
	fmt.Println("FIN")
	wg.Wait()
	fmt.Println("Exit")
}	