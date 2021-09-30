package main

import (
	"fmt"
)

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	for _, v := range evenVals {
		fmt.Println(v)
	}

	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	// ignoring the value
	for key := range uniqueNames {
		fmt.Println(key)
	}

	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	// for-range iterates over the runes in a string not bytes
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
	fmt.Println()

	// for-range value is a copy
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)

	outer:
		for _, sample := range samples {
			for i, r := range sample {
				fmt.Println(i, r, string(r))
				if r == 'l' {
					continue outer
				}
			}
			fmt.Println()
		}	

	}
}