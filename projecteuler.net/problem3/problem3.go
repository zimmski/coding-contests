package problem3

import (
	"math"
)

/*

Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/

// Bruteforce is just a quick implementation
func Bruteforce(n int) int {
	/*

		By simply dividing our number n with number i beginning from two
		and looking if n is divided exactly, we can say that the current
		number is made out of the prime factor i. Since n can be factored
		out of i more than one time, we have to check as long as n can be
		divided exactly and then move on to i + 1 for i if it does.

		We reached the largest prime factor when i is equal to n.

	*/

	i := 2

	for i != n {
		r := n % i

		if r == 0 {
			n /= i
		} else {
			i++
		}
	}

	return n
}

// FasterBruteforce is also just a quick implementation
func FasterBruteforce(n int) int {
	/*
		By ignoring all even numbers except two we can divide
		our number set for i by half.

		Also it is not possible that there is a prime factor greater
		than sqrt(n) (TODO explain in detail). This reduces our set
		of numbers again by (TODO how much? since it is recalculated).
	*/

	i := 2

	for i != n {
		r := n % i

		if r == 0 {
			n /= i
		} else {
			i++

			break
		}
	}

	nSqrt := int(math.Sqrt(float64(n)))

	for i != n {
		r := n % i

		if r == 0 {
			n /= i

			nSqrt = int(math.Sqrt(float64(n)))
		} else {
			i += 2

			if i > nSqrt {
				break
			}
		}
	}

	return n
}

func FasterBruteforceWithChannel(n int) int {
	/*
		The same as FasterBruteforce but with a channel to
		receive the next prime.
	*/

	next := make(chan int)

	go func() {
		next <- 2

		i := 3

		for {
			next <- i

			i += 2
		}
	}()

	i := <-next
	nSqrt := int(math.Sqrt(float64(n)))

	for i != n {
		r := n % i

		if r == 0 {
			n /= i

			nSqrt = int(math.Sqrt(float64(n)))
		} else {
			i = <-next

			if i > nSqrt {
				break
			}
		}
	}

	return n
}

/*

	TODO
	- implement better algorithms
	- document everything better

	After implementing the trivial bruteforce algorithms I could not think of better solutions.
	Wikipedia shows that prime factoring algorithms are called integer factorization algorithms
	https://en.wikipedia.org/wiki/Integer_factorization The brute force algorithms are called
	"Trivial division" https://en.wikipedia.org/wiki/Trial_division

	Better algorithms are
		Fermat's factorization method - which can be combined with trivial division
		Quadtratic sieve - fastest for number with <100 digits and second fastest for bigger numbers
		GNFS(general number field sieve) which is the fastest algorithm for bigger numbers

	Since these algorithms (except for fermat's factorization method) are kind of complicated to
	understand and implement and I do not think that prime factorization would give me any
	edge for the euler project, I do not plan to implement them. At least for now :-)

*/
