package problem2

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func assertProblem(t *testing.T, f func(n int) int) {
	Equal(t, 10, f(10))
	Equal(t, 4613732, f(4000000))
}

func TestBruteforce(t *testing.T) {
	assertProblem(t, Bruteforce)
}

func TestFasterBruteforce(t *testing.T) {
	assertProblem(t, FasterBruteforce)
}

func TestOptimal(t *testing.T) {
	assertProblem(t, Optimal)
}
