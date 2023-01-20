// Determine the minimal integer N such that L <= N <= D and sum of digits is X
// Determine the maximal integer M such that L < = M <= D and sum of digits is X

// Input looks like:
// L (1 - 10,000)
// D (1 - 10,000)
// X (1 - 36)

package main

import (
	"fmt"
	"math"
)

func main() {
	var L, D, X, N, M int
	fmt.Scan(&L, &D, &X)

	// Find minimal
	for i := L; i <= D; i++ {
		if sumDigits(i) == X {
			N = i
			break
		}
	}

	// Find maximal
	for i := D; i >= L; i-- {
		if sumDigits(i) == X {
			M = i
			break
		}
	}

	fmt.Println(N)
	fmt.Println(M)
}

func sumDigits(N int) int {
	// Num digits = round up log(N, 10)
	// for each digit, N % 10 is the final digit
	// N // 10 removes the last digit

	currentVal, sum := N, 0
	// Since log(1, 10) == 0, we need to add one to the result and floor so we get the proper number of digits for an input of 1.
	numDigits := int(math.Floor(math.Log10(float64(currentVal)) + 1.0))
	for i := 0; i < numDigits; i++ {
		sum += currentVal % 10
		currentVal = currentVal / 10
	}
	return sum
}
