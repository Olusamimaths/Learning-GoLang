package main

import "fmt"

// from https://yourbasic.org/golang/iota/
//STEP 1: create an integer
type Direction int

//STEP 2: list the values in iota
const (
	North Direction = iota
	East
	South
	West
)

//STEP 3: give the type a String method

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "Westie"}[d]
}

func main() {
	var d Direction = North
	fmt.Print(d)
	
	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}
}