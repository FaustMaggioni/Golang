package main
import (
	"fmt"
)
// tienen que ser de un mismo tipo

func todo(){
	arr := []int {1,2,3,4}
	arr2 := []string{"a","b","c"}
	arr2 = append(arr2,"d")
	fmt.Println(arr, arr2)
}
func main(){
		todo()
	}