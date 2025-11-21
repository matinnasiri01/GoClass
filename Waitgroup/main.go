package main

import (
    "fmt"
    "sync"
    "time"
)

// waitgroup is struct and for use it first make it
var wg sync.WaitGroup

func mygoroutine() {
   // defer: after all task done this line lunch
   // wg.Done() allow the wg to wait for this goroutine
    defer wg.Done()
    time.Sleep(1 * time.Second)
    fmt.Println("Hello")
}

func main() {
   // add count of your goroutine you have. Int value!
    wg.Add(1)
    // lunch:
    go mygoroutine()
    // call this to wait and goroutine do the tasks!
    wg.Wait()
    fmt.Println("finished")
}