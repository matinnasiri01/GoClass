package main

import (
	"fmt"
)

func main() {


	//  blank identifier
	// also can use in import 
	// import _ "fmt"

	var _ int = 12

	// initial value if not set always zero value of that type

	var test int 
	fmt.Printf("int -> %v\n", test)


	// Grouping variables
	// var (
	// 	name string = "John"
	// 	age  int    = 30
	// )


	// shadow variables

	x := 10
	fmt.Println("x before shadowing:", x)
	for x:= 1 ; ; x++{
		fmt.Println("x inside loop:", x)
	}


}