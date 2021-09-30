package main

import (
	"errors"
	"fmt"
	"os"
)

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// func divAndRemainder(numerator int, denominator int) (int, int, error) {
// 	if denominator == 0 {
// 		return 0, 0, errors.New("Cannot divide by zero")
// 	}
// 	return numerator / denominator, numerator % denominator, nil
// }

// named return values
func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func makeMul(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
func main() {
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	// supplying a slice as an argument ... must be after the variable name
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

	twoBase := makeMul(2)
	threeBase := makeMul(3)

	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}