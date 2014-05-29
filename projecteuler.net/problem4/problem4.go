package problem4

import (
	"strconv"
)

/*

Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.


*/

func IsPalindrome(n int) bool {
	s := strconv.Itoa(n)

	i := 0
	j := len(s) - 1

	for i <= j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

// Bruteforce is just a quick implementation
func Bruteforce(n int) int {
	maxNumber := 9
	minNumber := 1

	for i := 0; i < n-1; i++ {
		maxNumber = maxNumber*10 + 9
		minNumber = minNumber * 10
	}

	max := 0

	for a := maxNumber; a > minNumber-1; a-- {
		for b := maxNumber; b > minNumber-1; b-- {
			c := a * b

			if IsPalindrome(c) {
				if c > max {
					max = c
				}
			}
		}
	}

	return max
}

// FasterBruteforce is also just a quick implementation
func FasterBruteforce(n int) int {
	maxNumber := 9
	minNumber := 1

	for i := 0; i < n-1; i++ {
		maxNumber = maxNumber*10 + 9
		minNumber = minNumber * 10
	}

	max := 0

	for a := maxNumber; a > minNumber-1; a-- {
		for b := maxNumber; b > minNumber-1; b-- {
			c := a * b

			if IsPalindrome(c) {
				if c > max {
					max = c

					// only smaller numbers left with this b
					break
				}
			}
		}
	}

	return max
}

/*

	TODO find an optimal solution which exploits the attributes of a palindrom

*/
