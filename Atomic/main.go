package main

import (
    "fmt"
    "sync/atomic"
)


/*

Atomic operations in Go can be used in several practical scenarios, including:

1. Incrementing or decrementing shared counters:
   When multiple goroutines need to update a shared counter (for example, tracking
   the number of requests, active connections, or events), atomic operations allow
   safe modifications without using locks.

2. Shared flag variables:
   Atomic boolean or integer flags can be used to signal state changes between
   goroutines, such as stopping workers or notifying that a condition has been met.

3. Building non-blocking algorithms:
   The CompareAndSwap (CAS) family of functions enables lock-free operations where
   goroutines try to update a variable only if it still matches an expected value.
   This is the foundation of many fast, non-blocking algorithms.

4. Pointer memory operations:
   With functions like StorePointer and LoadPointer, atomic operations can be used
   to safely update or read pointers. This is useful for advanced memory management
   techniques such as implementing custom garbage-collection mechanisms.

5. Safe memory visibility across goroutines:
   Atomic Store and Load ensure that writes made by one goroutine become visible to
   others safely, without requiring mutexes. This is useful when sharing simple
   configuration values or state indicators.

6. Building concurrent lock-free data structures:
   The atomic package allows implementing fast, concurrent queues, stacks, and other
   lock-free data structures that can be accessed by many goroutines at the same time
   without using traditional locking techniques.

*/

func main() {
    var count int64
    atomic.AddInt64(&count, 1)
    fmt.Println(count)
}