// https://open.kattis.com/problems/2048
// Input: 4 lines of 4 ints, representing the board, followed by a single int 0-4 representing the direction
// 0, 1, 2, 3 represents left, up, right, down respectively

// Conceptually, if we are shifting everything TO THE LEFT:
// Start with the leftmost element in each row (each row is independent of the others)
// If it is zero, shift all the elements over by however many leading zeros there are (if there are 4 we skip the row)
// If it is a value, check if the next nonzero element to its right is the same
// If it is, double the value and zero out the other value
// Otherwise, keep the value where it is

package main

import (
	"fmt"
)

type Board [16]int

func main() {
	// Setup
	var inputVar int
	var board Board

	// Input
	for i := 0; i < 16; i++ {
		fmt.Scan(&inputVar)
		board[i] = inputVar
	}

	fmt.Scan(&inputVar)
	direction := inputVar

	// Output
	if direction == 0 {
		// Left (Seems to work)

		for row := 0; row < 4; row++ {
			for startCol := 0; startCol < 4; startCol++ {

				currVal := getBoardVal(board, row, startCol)
				if currVal == 0 {
					for col := startCol + 1; col < 4; col++ {
						nextVal := getBoardVal(board, row, col)
						if nextVal != 0 {
							setBoardVal(&board, row, startCol, nextVal)
							setBoardVal(&board, row, col, 0)
							break
						}
					}
				}

				for col := startCol + 1; col < 4; col++ {
					checkVal := getBoardVal(board, row, col)
					if checkVal == 0 {
						continue
					} else if checkVal == currVal {
						setBoardVal(&board, row, startCol, currVal*2)
						setBoardVal(&board, row, col, 0)
						break
					}
					break
				}
				fmt.Println("\nRow:", row)
				printBoard(board)
			}
		}
		fmt.Println("\nFinal\n")

		// for row := 0; row < 4; row++ {
		// 	firstEmptyIndex := 0
		// 	lastVal := 0
		// 	lastValIndex := 0

		// 	for col := 0; col < 4; col++ {
		// 		val := getBoardVal(board, row, col)

		// 		if val == 0 {
		// 			continue
		// 		} else if lastVal == 0 {
		// 			lastVal = val
		// 			lastValIndex = col
		// 			firstEmptyIndex++
		// 		} else if lastVal == val {
		// 			setBoardVal(&board, row, col, 0)
		// 			setBoardVal(&board, row, lastValIndex, val*2)
		// 		} else {
		// 			setBoardVal(&board, row, col, 0)
		// 			setBoardVal(&board, row, firstEmptyIndex, val)
		// 			firstEmptyIndex++
		// 		}
		// 	}
		// }
	} else if direction == 1 {
		// Up (Does not work, last row is all zeros?)
		for col := 0; col < 4; col++ {
			firstEmptyIndex := 0
			lastVal := 0
			lastValIndex := 0

			for row := 0; row < 4; row++ {
				val := getBoardVal(board, row, col)

				if val == 0 {
					continue
				} else if lastVal == 0 {
					lastVal = val
					lastValIndex = row
					firstEmptyIndex++
				} else if lastVal == val {
					setBoardVal(&board, row, col, 0)
					setBoardVal(&board, lastValIndex, col, val*2)
				} else {
					setBoardVal(&board, row, col, 0)
					setBoardVal(&board, firstEmptyIndex, col, val)
					lastVal = val
					firstEmptyIndex++
				}
			}

		}
	}
	printBoard(board)

}

// Helper methods
func getBoardVal(board Board, row int, col int) int {
	return board[row*4+col]
}

func setBoardVal(board *Board, row int, col int, newVal int) {
	board[row*4+col] = newVal
}

func printBoard(board Board) {
	for row := 0; row < 4; row++ {
		fmt.Println(board[row*4], board[row*4+1], board[row*4+2], board[row*4+3])
	}
}
