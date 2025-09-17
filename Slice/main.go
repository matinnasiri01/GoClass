package main

import "fmt"

func main() {

	/* 
		Slice
		make()
		len()
		cap()
		append()
	*/

	var s []string
	fmt.Println(s, s == nil)


	nums := []int{10, 20, 30, 40, 50}
	a := nums[:3]   // [10 20 30]
	b := nums[2:]   // [30 40 50]
	c := nums[:]    // All

	fmt.Printf("[:3] -> %v\n[2:] ->%v\n[:] ->%v\n",a,b,c)



	// make() as new 

	s2 := make([]string, 3)
	fmt.Println(s2)


	// make([]datatype, length, capacity)
	// length vs capacity
	// length -> number of elements in the slice
	// capacity -> number of elements in the underlying array

	sw := make([]int, 3, 5) 
	// sw = append(sw, 60, 70)
	fmt.Println(sw)      // [0 0 0]
	fmt.Println(len(sw)) // 3
	fmt.Println(cap(sw)) // 5



	// marging two slices
	even := []int{2, 4, 6}
	odd := []int{1, 3, 5}

	// ... unpack the slice
	numbers := append(odd, even...)

	println(numbers) // [1 3 5 2 4 6]

}