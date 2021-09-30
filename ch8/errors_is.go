package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		// return a new error wrapping the error from the os.Open() method
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

// custom error.Is comparator
type MyErr struct {
	Codes [] int // non-comparable with ==
}
func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}
func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok { // using type assertions
		return reflect.DeepEqual(me, me2)
	}
	return false
}

// custom error.Is to pattern match
// two ResourceErr instances should match when either field is set
type ResourceErr struct {
	Resource string
	Code int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s: %d", re.Resource, re.Code)
}

func (re ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code

		return matchResource && matchCode ||
				matchResource && ignoreCode ||
				matchCode && ignoreResource
	}
	return false
}

func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		// is os.ErrNotExist in the errorchaing of wrapped err?
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
		}
	}

	err = ResourceErr{
		Resource: "Database",
		Code:     123,
	}

	err2 := ResourceErr{
		Resource: "Network",
		Code:     456,
	}

	if errors.Is(err, ResourceErr{Resource: "Database"}) {
		fmt.Println("The database is broken:", err)
	}

	if errors.Is(err2, ResourceErr{Resource: "Database"}) {
		fmt.Println("The database is broken:", err2)
	}

	if errors.Is(err, ResourceErr{Code: 123}) {
		fmt.Println("Code 123 triggered:", err)
	}

	if errors.Is(err2, ResourceErr{Code: 123}) {
		fmt.Println("Code 123 triggered:", err2)
	}

	if errors.Is(err, ResourceErr{Resource: "Database", Code: 123}) {
		fmt.Println("Database is broken and Code 123 triggered:", err)
	}

	if errors.Is(err, ResourceErr{Resource: "Network", Code: 123}) {
		fmt.Println("Network is broken and Code 123 triggered:", err)
	}
}
