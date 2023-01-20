package main

import (
	"fmt"
)

func main() {
	var N, sum, currentVal int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&currentVal)
		sum += currentVal
	}
	fmt.Println(sum)
}
