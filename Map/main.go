package main

import "fmt"

func main() {


	/*
		We use make() for reference type variables:
		slice, map, and channel to allocate space for them.
		If we don't do this, we will get an error.
	*/


	// Map
	var m = make(map[int]string)

	// also 
	var w = map[int]string {
		0:"zero",
		1:"one",
	}

	// add
	m[0] = "nasiri"

	fmt.Println(m[0], m)
	fmt.Println(w)


	// check null or not
	// m => map[0:nasiri]
	// m[0] => "nasiri"
	// m[1] => ""
	
	value0, exist0 := m[0]
	fmt.Printf("value:%v exist:%v\n",value0,exist0)

	
	value1, exist1 := m[1]
	fmt.Printf("value:%v exist:%v\n",value1,exist1)



	// delete()

	for i:=1; i<3; i++ {
		m[i] = fmt.Sprintf("valueNum%v",i)
	}

	for k,v := range m {
		fmt.Printf("key:%v value:%v\n",k,v)
	}

	// one by one
	delete(m,1)
	fmt.Print(exist1)

	// all
	for key:= range m {
		delete(m, key)
	}

}