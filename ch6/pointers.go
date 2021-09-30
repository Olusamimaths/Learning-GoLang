package main

import "fmt"

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func update(g *int) {
	*g = 20
}

func main() {
	x := 10
	pointerToX := &x

	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)

	var f *int // f is nil
	failedUpdate(f)
	fmt.Println(f) 

	update(pointerToX)
	fmt.Println(x)

}