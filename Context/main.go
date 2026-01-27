package main

import (
	"context"
	"fmt"
	"time"
)

func processRequest(ctx context.Context) {
	userID := ctx.Value("UserID").(int)
	fmt.Printf("%d\n", userID)
}

func performTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("task cancelled")
			return
		default:
			fmt.Println("performing task...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {

	// create simple context:
	// context.Background()

	// Context With Data:					  key:	 value:
	ctx := context.WithValue(context.Background(), "UserID", 123)
	processRequest(ctx)

	// How to cancel a Context:
	ct, cancel := context.WithCancel(context.Background())
	go performTask(ct)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	// With Deadline:
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	// With Timeout:
	context.WithTimeout(context.Background(), 2*time.Second)

}
