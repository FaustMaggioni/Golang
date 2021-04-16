package main

import "fmt"
// abstract data type

type Car struct {
	Name string
	Age int
	ModelNumber int
}
//methods
func (c Car)Print(){
	fmt.Println(c)
}
func (c Car) GetName() string {
	return c.Name
}
func (c Car)Drive(){
	fmt.Println("driving...")
}
func main(){
	//c:= Car{"chev",1,2}
	c := Car{
		Name: "chevy",
		Age: 1,
		ModelNumber: 20,
	}
	c.Print()
	c.Drive()
	fmt.Println(c.Name)
} 