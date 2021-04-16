package main

import "fmt"
//Tipo de dato simil genérico
func Anything(anything interface{}){
	fmt.Println(anything)
}
func main(){
	Anything(2)
	Anything("Hola")
	Anything(3>2) //cualquier tipo de dato
	//otro tipo de uso
	mymap := make(map[string]interface{}) 
	//quiere decir q el mapeo tendrá como clave strings
	//pero como value cualquier tipo de dato
	mymap["name"]=01010101
	mymap["Hola"]=true
	mymap["ffff"]="BBBBB"
	fmt.Println(mymap)
}