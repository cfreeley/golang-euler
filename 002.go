package main

func P2() int {
	return evenFibSumUnder(4000000)
}

// Fib calculates the nth fibonacci number
// Contract: n >= 0
func Fib(n int) int {
	return fibAcc(n, 0, 1) 
}
// FibAcc is a helper accumulator for Fib
func fibAcc (n, prev, curr int) int {
	if n <= 0 {
		return curr
	}
	return fibAcc(n-1, curr, curr + prev)
}

// The above functions were a fun excercise, but it would be
// slow if we used them for summation. We really should be using
// a fold, but there isn't a generic one in Go, so I'm writing
// one for this case.

func evenFibSumUnder(limit int) int {
	return evenFibSumUnderAcc(limit, 0, 1, 0);
}

func evenFibSumUnderAcc(limit, prev, curr, evenSum int) int {
	if curr > limit {
		return evenSum
	} else if IsEven(curr) {
		evenSum += curr
	}
	return evenFibSumUnderAcc(limit, curr, curr + prev, evenSum)
}

// IsEven returns whether an int is an even number
// % is modulo, for the record. So this returns true if it can be divided
// by 2 and leave no remainder
func IsEven(n int) bool {
	return n % 2 == 0
}

