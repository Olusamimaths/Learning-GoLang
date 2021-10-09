package adder

import (
	"testing"
	"time"
	"fmt"
	"os"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("Set up stuff for tests here")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("Clean up stuff after tests here")
	os.Exit(exitVal) // you must call os.Exit with exitVal
}

func TestFirst(t *testing.T) {
	fmt.Println("TestFirst uses stuff set up in TestMain", testTime)
}
func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got ", result)
	}
}