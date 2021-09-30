// from https://www.hildeberto.com/2020/04/dealing-with-money.html
package main

import "fmt"

func main() {
	parties := split(100, 3)
	fmt.Println("split 100 with 3: %v\n", parties)

	parties = split(67, 4)
	fmt.Println("split 67 with 4: %v\n", parties)
}

func split(amount, n int) []int {
	division := amount / n
	toDistribute := amount - (division * n)

	parties := make([]int, n)

	// assign the equal division
	for i := 0; i < n; i++ {
		parties[i] = division
	}

	// assign whatever is left on top of the equal divivion
	j := 0
	for i := toDistribute; i > 0; i-- {
		parties[j] += 1
		j++
	}

	return parties
}