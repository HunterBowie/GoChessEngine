package chess

import "fmt"

var expectedPerft = map[int]int{
	0: 1,
	1: 20,
	2: 400,
	3: 8902,
	4: 197281,
	5: 4865609,
	6: 119060324,
}

// Prints the Perft results to the console for all depth values <= maxDepth
func PrintPerftResults(maxDepth int) {
	passed := true
	for depth := 0; depth <= maxDepth; depth++ {
		result := Perft(depth)
		expected := expectedPerft[depth]
		fmt.Printf("Depth: %d, Result: %d, Expected %d\n", depth, result, expected)
		if result != expected {
			passed = false
		}
	}

	if passed {
		fmt.Println("Perft Test Passed! :)")
	} else {
		fmt.Println("Perft Test Failed! :<")
	}

}

// Returns the number of nodes reached at a specified depth
func Perft(depth int) int {
	board := LoadBoardFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	return perft(board, depth)

}

func perft(board Board, depth int) int {
	nodes := 0

	if depth == 0 {
		return 1
	}

	for _, move := range GetAllLegalMoves(board) {
		boardCopy := board.Copy()
		boardCopy.PlayMove(move)
		nodes += perft(boardCopy, depth-1)
	}

	return nodes
}
