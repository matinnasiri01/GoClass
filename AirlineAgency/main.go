/*
Airline Agency


*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	co := make(map[string]string,n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		inp := scanner.Text()
		con := strings.Split(inp, " ")
		co[con[0]] = con[1]
	}
	
	test(*scanner)
	
}

func test (scanner bufio.Scanner){
	scanner.Scan()
	pcc, _ := strconv.Atoi(scanner.Text())
	pcn := make([]string, pcc)

	for i := 0; i != pcc; i++ {
		scanner.Scan()
		pcn[i] = strings.TrimPrefix(scanner.Text(),"+")
	}

	fmt.Println(pcn)
}
