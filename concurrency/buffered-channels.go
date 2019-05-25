/*
Channels can be buffered. Provide the buffer length as the second
argument to make to initialize a buffered channel:
ch := make(chan int, 100)
Sends to a buffered channel block only when the buffer is full.
Receives block when the buffer is empty.
 */
package main

import "fmt"

func main()  {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//!ch <- 3 //fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
