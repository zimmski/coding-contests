package problem1

/*

https://projecteuler.net/problem=1

Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

*/

// Bruteforce is just a quick implementation
// This is an O(n) solution with 2n modulos
func Bruteforce(n int) int {
	sum := 0

	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

// Optimal is the best solutions I made so fare
// This is an O(1) solution
func Optimal(n int) int {
	/*

		We can find out how many mupltiplies by 3 and 5 are in n by simply dividing N with 3 and 5

			count_3 = n / 3
			count_5 = n / 5

		To get the sum we have to multiply from 1 to count with their corresponding number.

			sum_3 = Σ[i = 1..count_3](i * 3)
			sum_5 = Σ[i = 1..count_5](i * 5)

		This can be reduced from the O(n) sum to a simple O(1) operation with https://en.wikipedia.org/wiki/Summation

			Σ[i = 1..n] = (n * (n + 1)) / 2

		and by applying the distributive property https://en.wikipedia.org/wiki/Distributive_property to factor the corresponding number out

			short_sum(n, k) = k * ((n * (n + 1)) / 2)
			multiples_sum(n, k) = short_sum(n / k, k)

			sum_3(n) = multiples_sum(n, 3)
			sum_5(n) = multiples_sum(n, 5)

		By summing up count_3 and count_5 we run into the problem that they overlap for all multiples of 3*5 = 15.
		Since the multiples of 3 and 5 represent two sets that overlap. We can remove the overlapping with the
		inclusion-exclusion principle https://en.wikipedia.org/wiki/Inclusion%E2%80%93exclusion_principle

			|A ∪ B| = |A| + |B| - |A ∩ B|

		Applied to our problem we can find the sum of all multiples of 3 or 5 below n with

			sum_3_or_5(n) = multiples_sum(n, 3) + multiples_sum(n, 5) - multiples_sum(n, 15)

		But we have to use 1000 - 1 as n since it is below 1000.
	*/

	short_sum := func(n, k int) int {
		return k * ((n * (n + 1)) / 2)
	}
	multiples_sum := func(n, k int) int {
		return short_sum(n/k, k)
	}

	return multiples_sum(n-1, 3) + multiples_sum(n-1, 5) - multiples_sum(n-1, 15)
}
