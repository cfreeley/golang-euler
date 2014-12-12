package main

/* Problem:
A palindromic number reads the same both ways. 
The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
*/

func P4() int {
	return palindromeProductSearchAndDestroy(999)
}

/*
Let's say we have a multiplication table that's 999x999
We can iterate the highest products by cutting it in half (since it's symmetrical)
across the diagonal and starting from the 999x999 corner to make a triangle. In other words,
if we call 999x999 point (0,0), all coordinates above the line y=x are removed from our table,
making it a triangle. 

The closer the number is to our origin corner of the triangle, the higher it will be.
Optimally we would travel along the closest hyperbola to the corner, but this up right path
will hit the first palindrome eventually, and once it does, we have a lower bound and
can just search through the numbers between them.
*/

// Finds the highest palindrome product backwards from the given start
func palindromeProductSearchAndDestroy(start int) int {
	row, col, pal, product, max := 0, 0, 0, start * start, start * start
	for pal < max { // If we have a pal greater than any products we're gonna find, we're done
		if IsPalindromeNumber(product) {
			pal = product
		}
		if row > col || product < pal { // Moves to the next column if this one can be skipped
			row = 0
			col++
			max = (start - row) * (start - col)
		} else {
			row++
		}
		product = (start - row) * (start - col)
	}
	return pal
}

// IsPalindromeNumber returns true if the number is a 'palindrome'
// invariant: n >= 0
func IsPalindromeNumber(n int) bool {
	slice := NumberSlice(n)
	// Iterates forward and backwards through the slice to find any discrepancies 
	for i, last := 0, len(slice) - 1; i <= len(slice) / 2; i++ {
		if slice[i] != slice[last-i] {
			return false
		}
	}
	return true
}

// NumberSlice makes a slice of all the individual digits of a Natural number
// invariant: n >= 0
func NumberSlice(n int) []int {
	if n < 10 {
		return []int{n}
	}
	return append(NumberSlice(n / 10), n % 10) // % 10 to get the 'one's place
}