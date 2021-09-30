package main

import (
	"fmt"
	"math/rand"
)

type person struct {
	name string
	age int
	id string
	email string
}

func createPerson(name string, age int, email string) person {
	return person {
		name,
		age,
		generateRandomId(),
		email,
	}
}

func generateRandomId() string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	id := ""
	for i := 1; i < len(letters)/2; i++ {
		id += string(letters[rand.Intn(i)])
	}
	return id
}

func main() {
	var p1  = createPerson("Sam", 24, "solathecoder")
	var p2  = createPerson("Tobi", 20, "tobison")
	var p3  = createPerson("Ayo", 19, "AY@gmail.com")

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println()

	persons := []person {p1, p2, p3}
	personMap := map[string]person {
		"first": p1,
		"second": p2,
		"third": p3,
	}
	fmt.Println(persons)

	fmt.Println()
	fmt.Println("First Person is ")
	fmt.Println(personMap["first"])

}