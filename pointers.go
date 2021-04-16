//una variable que guarda la direccion a otra variable

package main
import (
	"fmt"
)

func swap(m1,m2 *int){
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}
func main(){
	m1 := 2
	m2 := 10
	poin := &m1 //get addres de m1
	p2:= &m2
	fmt.Println(m1,m2)
	swap(poin,p2)
	fmt.Println(m1,m2) //* da el valor de esa dirección de memoria

}