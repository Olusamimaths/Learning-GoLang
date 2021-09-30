package main

import "fmt"


func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]

	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(d)
	fmt.Println(e)

	// copying the last 3 values of x on top of the first 3
	num := copy(x[:3], x[1:])
	fmt.Println(x, num)
}