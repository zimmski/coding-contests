package problem1

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func assertProblem(t *testing.T, f func(n int) int) {
	Equal(t, 23, f(10))
	Equal(t, 233168, f(1000))
}

func TestBruteforce(t *testing.T) {
	assertProblem(t, Bruteforce)
}

func TestOptimal(t *testing.T) {
	assertProblem(t, Optimal)
}
