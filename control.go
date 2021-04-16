package main

import "fmt"

func main(){
	//hay if, else, switch, for pero NO WHILE
	f:= true
	flag := &f //pointer es mejor por algo q no se
	if flag == nil{ // no poner ()
		fmt.Println("Nil")
	}else if *flag{
		fmt.Println("True")
	}else{
		fmt.Println("False")
	}

	for i := 0; i<10; i++{
		fmt.Println(i)
	}
	/* for {

	}
	loop infinito */
	arr := []int{4,5,6,7}
	for i,value := range arr{
		fmt.Println(i,value) //value solo tambien puede ir obvio
	}
	mymap :=make(map[string]interface{})
	mymap["name"]=01010171
	mymap["Hola"]=true
	mymap["ffff"]="BBBBB"
	for k, v := range mymap{
		fmt.Printf("Key: %s and Value: %v ",k,v)
	}

	//continue, break, switch
	flagX := true
	for i :=1; i<10; i++{
		if i%2 == 0{
			flagX = false
			break  // hace las veces de while (?
		}else if i==1{
			continue //retoma el loop
		}
		fmt.Println(i)
	}
	fmt.Println(flagX)

	day:= "Mon"
	switch day{ //util para catch errors
	case "Friday":
		fmt.Println("Frrriday")
	case "Mon", "Tue", "Wed":
		fmt.Println("Boring")
	default:
		fmt.Println("default")
	}
}