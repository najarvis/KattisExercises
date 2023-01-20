// Output all possible combinations of animals whose total number of legs adds up to the target value
// The output should be in lexigraphical order (sorted by first animal, then ties sorted by second animal, then third)
// If the number is impossible print "impossible" on a single line

// Input looks like:
// B C D L

// Where:
// 0 < B, C, D <= 100
// 0 <= L <= 250

package main

import (
	"fmt"
)

func main() {
	var B, C, D, L int
	found := false
	fmt.Scan(&B, &C, &D, &L)

	for first := 0; first <= L/B; first++ {
		legsLeftFirst := L - first*B
		for second := 0; second <= legsLeftFirst/C; second++ {
			legsLeftSecond := legsLeftFirst - second*C
			for third := 0; third <= legsLeftSecond; third++ {
				if legsLeftSecond-third*D == 0 {
					found = true
					fmt.Println(first, second, third)
				}
			}
		}
	}

	if !found {
		fmt.Println("impossible")
	}
}
