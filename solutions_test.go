package main

import (
	"testing"
)

func Test1 (t *testing.T) {
	CheckExpectInts(GaussSum(100), 5050,t)
	CheckExpectInts(ThreeOrFiveSum(10), 23,t)
	CheckExpectInts(ThreeOrFiveSum(1000), 233168,t) // Actual Answer
}

// Don't know if I can take parameters of unknown types
// So for now I'm making my tester on a type by type basis
func CheckExpectInts(actual, expected int, t *testing.T) {
	t.Log("expected:", expected, "actual:", actual)
	if actual != expected {
		t.Fail()
	}
}