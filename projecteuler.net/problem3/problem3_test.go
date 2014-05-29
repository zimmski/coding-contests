package problem3

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func assertProblem(t *testing.T, f func(n int) int) {
	Equal(t, 2, f(4))
	Equal(t, 3, f(12))
	Equal(t, 5, f(1200))
	Equal(t, 29, f(13195))
	Equal(t, 13, f(455))
	Equal(t, 6857, f(600851475143))
}

func TestBruteforce(t *testing.T) {
	assertProblem(t, Bruteforce)
}

func TestFasterBruteforce(t *testing.T) {
	assertProblem(t, FasterBruteforce)
}

func TestFasterBruteforceWithChannel(t *testing.T) {
	assertProblem(t, FasterBruteforceWithChannel)
}
