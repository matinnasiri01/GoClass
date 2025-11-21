package main

import (
    "fmt"
    "time"
)

func main() {
    
    unbufferedChannel := make(chan string)
    go func() {
        time.Sleep(5 * time.Second)
        unbufferedChannel <- "Hello"
    }()

    fmt.Println("Started")
    // Blocking Operation
    // in this part code is stop and wait for '<-unbufferedChannel' anser.
    fmt.Println(<-unbufferedChannel)
    fmt.Println("Finished")

    select {
    case msg := <-unbufferedChannel:
    fmt.Println("after 5 seconds")
        fmt.Println(msg)
    }
}
