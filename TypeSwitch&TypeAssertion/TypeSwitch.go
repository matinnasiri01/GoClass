
package main

import "fmt"

func myfunc(a interface{}) {

    switch a.(type) {

    case int:

        fmt.Println("Type: int, Value: ", a.(int))

    case string:

        fmt.Println("Type: string, Value: ", a.(string))

    case float64:

        fmt.Println("Type: float64, Value: ", a.(float64))

    default:

        fmt.Println("Type not Found")

    }

}

func main() {

    myfunc("Quera")

    myfunc(12.3)

    myfunc(true)

}
