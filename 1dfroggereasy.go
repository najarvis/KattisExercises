// The first line contains 3 integer, N, S, and M
// N is the number of board squares (1 <= N <= 200)
// S is the index of the starting square (1 <= S <= M)
// M is the magic number

// The next line contains n space separated non-zero integers representing the board (each with a value bound by [-200, 200])

package main

import (
	"fmt"
)

func main() {
	var N, S, M, boardSquare int
	fmt.Scan(&N, &S, &M)

	jumps := 0
	// This is just a set
	seen := make(map[int]struct{})

	// Construct the board
	board := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&boardSquare)
		board[i] = boardSquare
	}

	for {
		board_value := board[S-1]
		if board_value == M {
			fmt.Println("magic")
			fmt.Println(jumps)
			break
		}

		seen[S] = struct{}{}
		S += board_value
		jumps++
		if S < 1 {
			fmt.Println("left")
			fmt.Println(jumps)
			break
		}
		if S > N {
			fmt.Println("right")
			fmt.Println(jumps)
			break
		}
		_, ok := seen[S]
		if ok {
			fmt.Println("cycle")
			fmt.Println(jumps)
			break
		}
	}
}
