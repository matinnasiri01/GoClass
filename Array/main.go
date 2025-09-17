package main

import "fmt"

func main() {

	// Array
	// [size]datatype{values}
	var zero [3]int
	var init = [3]int{1, 2, 3}

	fmt.Printf("zero: %v\ninit %v\n", zero,init)


	// if you don't know the size of the array
	var arr = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("arr: %v\n", arr)


	// add values to specific index
	var sarr [3]string
	sarr[0] = "Hello"
	sarr[1] = "World"
	sarr[2] = "!"
	fmt.Printf("sarr: %v\n", sarr)

	// %v -> value & %#v -> value with type
	// %d -> integer


	// reading values from array
	// range based for loop
	var name = [...]string{"Matin", "John", "Doe"}
	for index, element := range name {
     	fmt.Printf("Element %v At Index %d\n", element, index)
  	}


	// for index, element := range arr {} Normal usage of range
	// for _, element := range arr {} Omit index with _ and use element
	// for index := range arr {} Use index only
	// for range arr {} Simply loop over the array


	// multi dimensional array
	var matrix = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}

	
}