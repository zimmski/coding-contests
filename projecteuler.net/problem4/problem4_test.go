package problem4

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func assertProblem(t *testing.T, f func(n int) int) {
	Equal(t, 9009, f(2))   // 9009 = 91 * 99
	Equal(t, 906609, f(3)) // 906609 = 993 * 913
}

func TestBruteforce(t *testing.T) {
	assertProblem(t, Bruteforce)
}

func TestFasterBruteforce(t *testing.T) {
	assertProblem(t, FasterBruteforce)
}

func TestIsPalindrome(t *testing.T) {
	True(t, IsPalindrome(1))
	True(t, IsPalindrome(11))
	True(t, IsPalindrome(121))

	False(t, IsPalindrome(12))
	False(t, IsPalindrome(123))
}
