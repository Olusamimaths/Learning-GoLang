package main

import "fmt"

func main() {
	// unbuffered channels: every write causes goroutine to pause until another goroutine reads from same channel
	// same with read
	ch1 := make(chan int) // making any or both of the channels buffered  fixes deadlock i.e make(chan int,1)
	ch2 := make(chan int)
	go func ()  {
		v := 1
		ch1 <- v // written to ch1, this goroutine cannot proceed until ch1 is read
		v2 := <-ch2
		/**
		// to see the randomness of the select statement
		var v2 int
		select {
		case ch1 <- v:
		case v2 = <-ch2:
		}
		**/
		fmt.Println(v, v2)
	} ()

	v := 2
	
	var v2 int
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println(v, v2)
}