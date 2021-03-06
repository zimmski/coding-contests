package problem2

import (
	"math"
)

/*

Even Fibonacci numbers

Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

*/

// Bruteforce is just a quick implementation
func Bruteforce(n int) int {
	/*

		This solution validates every Fibonacci term

	*/

	a := 1
	b := 2

	sum := 0

	for b < n {
		if b%2 == 0 {
			sum += b
		}

		a, b = b, a+b
	}

	return sum
}

// FasterBruteforce is also just a quick implementation
func FasterBruteforce(n int) int {
	/*

		Since every third term of the Fibonacci sequence is even
		we can reduce the steps to one third and since we
		can guarantee that every third term is even (TODO show proof
		adding two even make even two odd make even one even one odd make odd)
		we can omit the validation for even.

	*/

	a := 1
	b := 2

	sum := 0

	for b < n {
		sum += b

		/*
			Take three terms at one time

			1: a, b = b, a + b
			2: a, b = a + b, b + (a + b)
			3: a, b = b + (a + b), (a + b) + (b + (a + b))
			        = a + 2 * b, 2 * a, 3 * b
		*/
		a, b = a+2*b, 2*a+3*b
	}

	return sum
}

// Optimal is the best solution I made so fare
func Optimal(n int) int {
	/*

		While reading the Fibonacci wikipedia articel https://en.wikipedia.org/wiki/Fibonacci_number
		I stumbled upon the relation of the Fibonacci sequence and the golden ratio which allows
		us to directly calculate a Fibonacci term without needing all preceding terms. https://en.wikipedia.org/wiki/Fibonacci_number#Limit_of_consecutive_quotients

			φ (golden ratio) = (1 + sqrt(5)) / 2

			F(n) = (φ^n - (-φ)^(-n)) / sqrt(5)

		Although this is indeed neat, it does not justify replacing a few simple additions and multiplications with two enormous exponentiations.

		Luckily there is another relation https://en.wikipedia.org/wiki/Fibonacci_number#Limit_of_consecutive_quotients
		The ratio of two consecutive Fibonacci terms converge towards the golden ratio.

			lim[n -> ∞](F(n + 1) / F(n)) = φ

		Furthermore this property holds for two terms which have terms between them

			lim[n -> ∞](F(n + α) / F(n)) = φ^α

		We can therefore take three terms at one time by simply multiplying φ^α to our current term. Furthermore we can save φ^α
		which reduces the whole calculation for three term-steps to one multiplication.

		Since I do not see any other properties to obtain the set of even terms of the Fibonacci sequence or even
		the sum of even terms I do not think that there is a better solution than one multiplication and summing up the resulting terms.

	*/

	i := 2
	sum := 0

	phiThree := math.Pow(math.Phi, 3.0)

	for i < n {
		sum += i

		// round to the nearest positive integer with Floor(i + 0.5)
		i = int(math.Floor(float64(i)*phiThree + 0.5))
	}

	return sum
}
