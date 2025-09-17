// Hop Game
// ----------------------------
// In this game, two numbers p and q are given.
// Players count from 1 up to q in order.
// - If the current number is not a multiple of p, print the number.
// - If it is a multiple of p, print "Hop" repeated (number / p) times.
// The game ends once the number q is reached.
//
// Example: p = 3, q = 10
// Output: 1, 2, Hop, 4, 5, Hop Hop, 7, 8, Hop, 10

package main


import (
	"fmt"
	"strings"
)

func main() {

	var p, q int
	var word = "Hope"
	fmt.Print("Enter the p and q:")
	fmt.Scanf("%d %d", &p, &q)

	for i := 1; i <= q; i++ {
		if i%p == 0 {
			result := ""
			for o := 0; o < i/p; o++ {
				result += word + " "
			}
			fmt.Println(strings.TrimSpace(result))
		} else {
			fmt.Println(i)
		}
	}
}