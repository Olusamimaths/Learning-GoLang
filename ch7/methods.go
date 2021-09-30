package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName string
	Age int
}
// method has to be declared in package level
func (p Person) String() string {
		return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// pointer and value receivers
type Counter struct {
	total int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++;
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {
	person := Person {
		"Sam",
		"Tobi",
		22,
	}
	fmt.Println(person) // String called automatically

	var c Counter
	fmt.Println(c.String())
	c.Increment() // automatically converted to (&c).Increment()
	fmt.Println(c.String())
}