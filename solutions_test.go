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

func Test3 (t *testing.T) {
	if IsPrime(927) || !IsPrime(103) {
		t.Error("IsPrime(927) =", IsPrime(927), "; IsPrime(103)", IsPrime(103))
	}
	CheckExpectInts(P3(), 6857, t)
}

func Test4 (t *testing.T) {
	if !IsPalindromeNumber(47574) || IsPalindromeNumber(4454) {
		t.Error("IsPalindromeNumber(47574) =", IsPalindromeNumber(47574), 
			"; IsPalindromeNumber(4454) =", IsPalindromeNumber(4454))
	}
	CheckExpectInts(NumberSlice(1234)[0], 1, t)
	CheckExpectInts(NumberSlice(1234)[2], 3, t)
	CheckExpectInts(P4(), 906609, t)
}

// Current slowest solution!
func Test5 (t *testing.T) {
	CheckExpectInts(SmallestEvenlyDivisibleBy(10), 2520, t)
	CheckExpectInts(P5(), 232792560, t)
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