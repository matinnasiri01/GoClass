package main

import (
	"fmt"
	"time"
)
func sayHello(name string) {
    fmt.Println("Hello " + name)
}
func main() {
	
    people := []string{"Rose", "Erich", "Amelia"}
    for _, person := range people {
        go sayHello(person)
    }
    time.Sleep(1 * time.Second)
    


    /// anonymous functions & Goroutine:
    go func() {
        for i := 0; i < 3; i++ {
            fmt.Println("Goroutine:", i)
            time.Sleep(time.Millisecond * 300)
        }
    }()
    for i := 0; i < 3; i++ {
        fmt.Println("Main:", i)
        time.Sleep(time.Millisecond * 300)
    }
    time.Sleep(1 * time.Second)

}