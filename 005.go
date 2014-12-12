package main

func P5() int {
	return SmallestEvenlyDivisibleBy(20)
}

func SmallestEvenlyDivisibleBy(n int) int {
	i := n
	for {
		if allEvenlyDisibleBy(i, n) {
			return i
		}
		i += n
	}
}

func allEvenlyDisibleBy(candidate, goal int) bool {
	for i := goal; i > 1; i-- {
		if candidate % i > 0 {
			return false
		}
	}
	return true
}