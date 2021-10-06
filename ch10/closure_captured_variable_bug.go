package main

import "fmt"

func searchData(s string, searchers []func(string) [] string) []string {
	done := make(chan struct{}) // empty struct because the value is unimportant, we don't want to read from the channel, just want to close it
	result := make(chan []string)
	for _, searcher := range searchers {
		go func(searcher func(string)[]string)  {
			select {
			case result <- searcher(s):
			case <-done:
			}
		}(searcher)
	}

	r := <-result // the first read proceeds
	close(done)
	return r
}

func main() {
	nums := []int{1, 2, 4, 6, 10}
	ch := make(chan int, len(nums))

	// for _, v := range nums {
	// 	go func() {
	// 		ch <- v * 2
	// 	}()
	// }

	// for i := 0; i < len(nums)-1; i++ { // the behaviour when i < len(nums-1)
	// 	fmt.Println(<-ch)
	// }

	for _, v := range nums {
		go func(val int) {
			ch <- val * 2
		}(v)
	}

	for i := 0; i < len(nums); i++ { // the behaviour when i < len(nums-1)
		fmt.Println(<-ch)
	}
}