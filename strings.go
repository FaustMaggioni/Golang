package main
import (
	"fmt"
	"strings"
)

func main(){
	var m1 string = "Hola mundo"
	m2 := "mundo"
	fmt.Println(strings.Contains(m2,m1))
	fmt.Println(strings.ReplaceAll(m1,"o","NO"))
	fmt.Println(strings.Split(m1," "))
}