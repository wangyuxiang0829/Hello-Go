/*
A 'goroutine' is a lightweight thread managed by the Go runtime.
go f(x, y, z) starts a new goroutine running.The evaluation of
f, x, y, and z happens in the current goroutine and the execution
of f happens in the new goroutine.
 */
package main

import (
	"fmt"
	"time"
)

func say(s string)  {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main()  {
	go say("world")
	say("hello")
}
