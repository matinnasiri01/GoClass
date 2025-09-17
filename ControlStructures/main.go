package main

import "fmt"

func main(){

	/* 
		if
		else if
		else
	*/ 

	x := 2
	if x > 1 {
		fmt.Println(x)
	}


	if x := 10; x > 5 {
    		fmt.Println("x is greater than 5")
	}
//---------------------------------------------------------------------



	/*
		switch
		case
		fallthrough 
		switch true
	*/


	name := "mohammad"

	switch name {
	case "ahmad":
	    fmt.Println("Hello Ahmad")
	case "mohammad":
	    fmt.Println("Hello Mohammad")
	default:
	    fmt.Println("Greetings")
	}

	z := 3
	switch {
		case z>2 :
			fmt.Println("z > 2")
			fallthrough
			// run next case even wrong!
		case z<2 :
			fmt.Println("z < 2")
		default: 
			fmt.Println(z)
	}


	y := 3
	switch {
		case y>2 :
			fmt.Println("y > 2")
		default: 
			fmt.Println(y)
	}
//---------------------------------------------------------------------

	/*
		for
	*/

	var target int = 5
	sum := 0
	for i := 1; i <= target; i++ {
	    sum += i
	}
	fmt.Println(sum)

	fmt.Println("//=============")


	for i:=1 ; ; i++{
		fmt.Println(i)
		if i > 12 {
			break
		}
	}
	
}