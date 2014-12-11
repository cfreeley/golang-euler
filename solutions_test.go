package main

import (
	"testing"
)

func Test1 (t *testing.T) {
	CheckExpectInts(GaussSum(100), 5050,t)
	CheckExpectInts(ThreeOrFiveSum(10), 23,t)
	CheckExpectInts(P1(), 233168, t) // Actual Answer
}

func Test2 (t *testing.T) {
	if IsEven(3) || !IsEven(4) {
		t.Error("IsEven(3) =", IsEven(3), "; IsEven(4) =", IsEven(4))
	}
	CheckExpectInts(Fib(10), 89, t)
	CheckExpectInts(P2(), 4613732, t)
}

// Go doesn't like generics or overloading, so making specific
// equality testers

// CheckExpectInts Fails if the given ints are not equal
func CheckExpectInts(actual, expected int, t *testing.T) {
	t.Log("expected:", expected, "actual:", actual)
	if actual != expected {
		t.Fail()
	}
}