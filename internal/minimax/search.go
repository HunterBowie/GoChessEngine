package minimax

import (
	"strconv"

	"github.com/HunterBowie/GoChessEngine/internal/chess"
)

type SearchResults struct {
	BestMove *chess.Move
	Score int
}

// Search performs a minimax search to the given depth
// Returns the bestmove and associated score
func Search(board chess.Board, depth int) SearchResults {
	if depth < 0 {
		panic("Depth: " + strconv.Itoa(depth) + "is invalid for Search")
	}

	if depth == 0 || board.GetGameState() != chess.GamePlayState {
		return SearchResults{nil, Evaluate(board)}
	}


	maximizing := true

	if board.ActiveColor == chess.Black {
		maximizing = false
	}

	var bestResult *SearchResults

	for _,move := range chess.GetAllLegalMoves(board) {
		boardCopy := board.Copy()
		boardCopy.PlayMove(move)
		result := Search(boardCopy, depth - 1)
		result.BestMove = &move

		if bestResult == nil {
			bestResult = &result
			continue
		}

		if maximizing && result.Score > bestResult.Score {
			bestResult = &result
		} else if (!maximizing && result.Score < bestResult.Score) {
			bestResult = &result
		}
	}

	return *bestResult
	
}
