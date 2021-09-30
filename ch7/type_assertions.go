package main

import (
	"fmt"
)

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt) // i2 is of type MyInt
	// i2 := i.(string) // causes panic
	fmt.Println(i2)
}
