package main

import (
	"fmt"
)

func main() {
	var N int
	var word string
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&word)
		if i%2 == 1 {
			continue
		}
		fmt.Println(word)
	}
}
