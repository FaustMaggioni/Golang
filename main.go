package main

import "fmt"

func main(){
	var m1 int
	m1 = 2
	m7 := 10
	var (
		m2= 3
		m3= 4
	)
	var m4 int32 = 5
	var m5 int64 = 6
	fmt.Println(m1,m2,m3)
	fmt.Println(int64(m4)+m5)
}
