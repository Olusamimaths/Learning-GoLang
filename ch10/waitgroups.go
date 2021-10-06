package main

import (
	"fmt"
	"sync"
)

// making sure that a channel being written to by multiple goroutines is only closed once
func processAndGather(in <- chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func ()  {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}

	go func ()  {
		wg.Wait()
		close(out)
	}()
	
	var result []int
	for v := range out {
		result = append(result, v)
	}
	return result
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		doThing1()
	}()
	go func() {
		defer wg.Done()
		doThing2()
	}()
	go func() {
		defer wg.Done()
		doThing3()
	}()
	wg.Wait()
}

func doThing1() {
	fmt.Println("Thing 1 done!")
}

func doThing2() {
	fmt.Println("Thing 2 done!")
}

func doThing3() {
	fmt.Println("Thing 3 done!")
}
