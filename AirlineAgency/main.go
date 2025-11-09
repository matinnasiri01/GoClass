package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	countries := make(map[string]string)

	for i := 0; i < n; i++ {
		var name, code string
		fmt.Scan(&name, &code)
		countries[code] = name
	}

	var q int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var number string
		fmt.Scan(&number)

		found := false
		for code, country := range countries {
			if strings.HasPrefix(number, code) {
				fmt.Println(country)
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Invalid Number")
		}
	}
}
