package minimax

import (
	"math"
	"strconv"

	"github.com/HunterBowie/GoChessEngine/internal/chess"
)

type SearchResults struct {
	BestMove *chess.Move
	Score    int
}

func Search(board chess.Board, depth int) SearchResults {
	return search(board, depth, math.MinInt, math.MaxInt)
}

// Search performs a minimax search to the given depth
// Returns the bestmove and associated score
func search(board chess.Board, depth int, alpha int, beta int) SearchResults {
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

	moves := chess.GetAllLegalMoves(board)

	if depth == 1 {
		moves = chess.GetMoves(board, false, true, true)
	}

	for _, move := range moves {
		boardCopy := board.Copy()
		boardCopy.PlayMove(move)
		result := search(boardCopy, depth-1, alpha, beta)
		result.BestMove = &move

		if bestResult == nil {
			bestResult = &result
			continue
		}

		if maximizing && result.Score > bestResult.Score {
			bestResult = &result
			alpha = max(alpha, result.Score)
			if beta <= alpha {
				break
			}
		} else if !maximizing && result.Score < bestResult.Score {
			bestResult = &result
			beta = min(beta, result.Score)
			if beta <= alpha {
				break
			}
		}
	}

	return *bestResult

}
