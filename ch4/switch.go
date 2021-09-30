package main

import (
	"fmt"
	"math/rand"
)

func main() {
	words := []string{"a", "cow", "smile", "gopher","octopus", "anthropologist"}
	
	for _, word := range words {
		switch size := len(word) ; size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5: 
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}
	loop:
		for i := 0; i < 10; i++ {
			switch {
				case i%2 == 0:
					fmt.Println(i, "is even")
				case i%3 == 0:
					fmt.Println(i, "is divisible by 3 but not 2")
				case i%7 == 0:
					fmt.Println("exit the loop!")
					break loop
				default: 
					fmt.Println(i, "is boring")
			}
		}
	
		// use case of goto
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")

	done:
		fmt.Println("do complicated stuff no matter why we left the loop")
		fmt.Println(a)
}