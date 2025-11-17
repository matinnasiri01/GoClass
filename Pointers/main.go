package main

import "fmt"

func main() {


    var a  int  = 24
    var p *int
    fmt.Println(p)
    p = &a
    fmt.Println(p)
    fmt.Println(dereferencing(p))

    s := new(string)
    *s = "hellow"
    fmt.Printf("address: %d\nvalue: %s",s,*s)


    pls := 23
    callByValue(pls) // afret this func pls still 23
    callByReference(&pls)

}


func dereferencing(p *int) int{
	return *p
}


func callByValue(i int){
	i += 1
}
func callByReference(i *int){
	*i += 1
}