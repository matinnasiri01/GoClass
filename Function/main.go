package main

import "fmt"

func main() {

	defer fmt.Printf("%d",sum(12,21))


	func (i int){
		fmt.Printf("Inner func %d",i)
	}(12)



}

func sum(p, q int) int {
	return p + q
}
// use type to keep clean code
type Opp func (int,int) int


func high(f Opp){ }

