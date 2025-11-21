package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a buffered channel with capacity 3
	ch := make(chan int, 3)

	// Sending three values without any receiver (allowed because of buffer)
	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println("Three values sent to buffered channel")

	// Receiving values (FIFO order)
	fmt.Println("Reading from channel:")
	fmt.Println(<-ch) // 10
	fmt.Println(<-ch) // 20

	// Add a new value after freeing buffer space
	ch <- 40

	// Closing the channel
	close(ch)

	fmt.Println("\nReading remaining values using for-range:")

	// Iterate over remaining values in the closed channel
	for v := range ch {
		fmt.Println(v)
	}

	// ⚠ Sending to a closed channel causes panic
	// ch <- 50 // -> panic: send on closed channel


	// make readonly or ...
	name := make(chan string , 1)
	sendData(name,"Matin")
	fmt.Printf("\nHellow %s",receiverData(name))



	/// Select
	ch1 := make(chan string,1)
	ch2 := make(chan string,1)
	go func ()  {
		time.Sleep(1 * time.Second)
		ch1 <- "Ch1"	
	}()
	go func ()  {
		time.Sleep(2 * time.Second)
		ch2 <- "Ch2"	
	}()
	for i := 0; i < 2; i++{
		select{
		case msg1 := <- ch1:
			fmt.Println(msg1)
		
		case msg2 := <- ch1:
			fmt.Println(msg2)
		}
	} 
}
func sendData (ch chan<- string,text string){
	ch <- text
}

func receiverData(ch <-chan string) string{
	return <-ch 
}