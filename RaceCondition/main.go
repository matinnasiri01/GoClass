package main

import (
    "fmt"
    "time"
)

// Each time this code is executed, it has a different result.
func main() {
    raceVariable := 0
    // 1000 Goroutine want to add 1 to raceVariable
    // Some Goroutines cannot read the actual value and make errors.
    for i := 0; i < 1000; i++ {
        go func() {
            raceVariable += 1
        }()
    }

    time.Sleep(2 * time.Second)
    fmt.Println(raceVariable)
}