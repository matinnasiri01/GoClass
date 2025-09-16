package main

import "fmt"

func main() {


	/* 
	rune  ♬
	byte  (0 to 255)
	int8  (-128 to 127)
	int16 (-32,768 to 32,767)
	int32 (-2,147,483,648 to 2,147,483,647)
	int64 (-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807)
	uint8  (0 to 255)
	uint16 (0 to 65,535)
	uint32 (0 to 4,294,967,295)
	uint64 (0 to 18,446,744,073,709,551,615)
	float32 (1.18E-38 to 3.4E+38)
	float64 (2.23E-308 to 1.8E+308)
	bool (true or false)
	string ("matin nasiri")
	*/
	

	// var can change value
	name := "matin nasiri"
     fmt.Println(name)
	
	// const can not change value
	const age int8 = 18
	fmt.Println(age)
}