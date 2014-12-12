package main

import (
	"testing"
	"fmt"
)

func Test1 (t *testing.T) {
	CheckExpectInts(GaussSum(100), 5050,t)
	CheckExpectInts(ThreeOrFiveSum(10), 23,t)
	CheckExpectInts(P1(), 233168, t) // Actual Answer
}

func Benchmark1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintln(P1())
    }
}

func Test2 (t *testing.T) {
	if IsEven(3) || !IsEven(4) {
		t.Error("IsEven(3) =", IsEven(3), "; IsEven(4) =", IsEven(4))
	}
	CheckExpectInts(Fib(10), 89, t)
	CheckExpectInts(P2(), 4613732, t)
}

func Benchmark2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintln(P2())
    }
}

func Test3 (t *testing.T) {
	if IsPrime(927) || !IsPrime(103) {
		t.Error("IsPrime(927) =", IsPrime(927), "; IsPrime(103)", IsPrime(103))
	}
	CheckExpectInts(P3(), 6857, t)
}

func Benchmark3(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintln(P3())
    }
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

func Benchmark4(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintln(P4())
    }
}

// Second slowest solution!
func Test5 (t *testing.T) {
	CheckExpectInts(SmallestEvenlyDivisibleBy(10), 2520, t)
	CheckExpectInts(P5(), 232792560, t)
}

func Benchmark5(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintln(P5())
    }
}

// Current slowest solution!
func Test92(t *testing.T) {
	CheckExpectInts(SquareDigitSum(16), 37, t)
	CheckExpectInts(SquareDigitChain(145, make([]int, 10000000)), 89, t)
	CheckExpectInts(P92(), 8581146, t)
}

func Benchmark92(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintln(P92())
    }
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