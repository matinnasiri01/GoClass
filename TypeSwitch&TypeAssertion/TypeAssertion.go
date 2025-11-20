
package main

import "fmt"

func myfunc1(a interface{}) {


    val,ok := a.(string)

    fmt.Printf("Value: $s\n,status: %t", val,ok)

}

func main1() {

    var val interface{} = "Quera"

    myfunc(val)

}
