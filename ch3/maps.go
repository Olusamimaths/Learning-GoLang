package main

import "fmt"

func main() {
	var nilMap map[string]int // map[keyType]valueType
	fmt.Println(len(nilMap))

	teams := map[string][]string {
		"Chelsea": {"Kante", "Rudiger", "Mendy"},
		"PSG": [] string{"Messi", "Mbappe", "Neymar"},
	}
	fmt.Println(teams["PSG"])

	// Using a map as a set
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])

	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}