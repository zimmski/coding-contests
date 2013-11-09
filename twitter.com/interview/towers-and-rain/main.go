package main

/*

	Problem described by http://qandwhat.apps.runkite.com/i-failed-a-twitter-interview/

	I quote (with own example)

	> "Consider the following picture:"

	|     x
	|     x
	|     x
	| x   xx x
	| x   xx x
	| xx  xx x
	|xxx  xx x
	|xxx  xx x
	|xxxxxxxxx
	----------

	> "In this picture we have walls of different heights. This picture is represented by an array of integers, where the value at each index is the height of the wall. The picture above is represented with an array as [3 6 4 1 1 9 6 1 6]."

	> "Now imagine it rains. How much water is going to be accumulated in puddles between walls?"

	|     x
	|     x
	|     x
	| x---xx-x
	| x---xx-x
	| xx--xx-x
	|xxx--xx-x
	|xxx--xx-x
	|xxxxxxxxx
	+---------

	> "We count volume in square blocks of 1X1. So in the picture above, everything to the left of index 1 spills out. Water to the right of index 7 also spills out. We are left with a puddle between 1 and 6 and the volume is 10."


*/

import (
	"fmt"
	"math/rand"
	"time"
)

const HEIGHT = 10
const WIDTH = 9

func generate(width int) []int {
	var a = make([]int, width)

	for i, _ := range a {
		a[i] = rand.Intn(HEIGHT)
	}

	return a
}

func display(a []int) {
	for y := HEIGHT - 1; y > 0; y-- {
		fmt.Print("|")

		for x := 0; x < len(a); {
			if y <= a[x] {
				for x < len(a) {
					fmt.Print("x")

					var found = false
					var n = x + 1

					for ; n < len(a); n++ {
						if y <= a[n] {
							found = true

							break
						}
					}

					x++

					if found {
						for ; x < n; x++ {
							fmt.Print("-")
						}
					} else {
						for ; x < n; x++ {
							fmt.Print(" ")
						}
					}
				}

				break
			} else {
				fmt.Print(" ")

				x++
			}
		}

		fmt.Print("\n")
	}

	fmt.Print("+")

	for i := 0; i < len(a); i++ {
		fmt.Print("-")
	}

	fmt.Print("\n ")

	for _, v := range a {
		fmt.Print(v)
	}

	fmt.Print("\n")
}

func calc(a []int) int {
	var water = 0

	var l, r = 0, len(a) - 1
	var maxL, maxR = 0, 0

	for l <= r {
		if maxL <= maxR {
			var i = a[l]

			if maxL < i {
				maxL = i
			} else {
				water += maxL - i
			}

			l++
		} else {
			var i = a[r]

			if maxR < i {
				maxR = i
			} else {
				water += maxR - i
			}

			r--
		}
	}

	return water
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var a = generate(WIDTH)

	display(a)

	fmt.Printf("\nWater: %d\n", calc(a))
}
