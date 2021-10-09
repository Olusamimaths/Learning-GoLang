package person

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name: "Sam",
		Age : 20,
	}

	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})

	result := CreatePerson("Sam", 20)
	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}
}