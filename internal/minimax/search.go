package minimax

import (
	"math"
	"time"

	"github.com/HunterBowie/GoChessEngine/internal/chess"
)

type SearchResults struct {
	BestMove *chess.Move
	Score    int
}

func Search(board chess.Board, timeMilliseconds int) SearchResults {
	startTime := time.Now()
	return search(board, 4, math.MinInt, math.MaxInt, startTime, int64(timeMilliseconds))
}

// Search performs a minimax search to the given depth
// Returns the bestmove and associated score
func search(board chess.Board, depth int, alpha int, beta int, timeStarted time.Time, totalMilliseconds int64) SearchResults {
	if depth == 0 {
		return SearchResults{nil, Evaluate(board)}
	}

	moves := chess.GetAllLegalMoves(board)

	if len(moves) == 0 {
		if chess.IsKingInCheck(board) {
			if board.ActiveColor == chess.White {
				return SearchResults{nil, math.MinInt}
			}
			return SearchResults{nil, math.MaxInt}
		} else {
			return SearchResults{nil, 0}
		}
	}

	if depth == 1 {
		moves = chess.GetMoves(board, false, true, true)
	}

	var bestResult *SearchResults

	maximizing := true

	if board.ActiveColor == chess.Black {
		maximizing = false
	}

	for _, move := range moves {
		boardCopy := board.Copy()
		boardCopy.PlayMove(move)
		result := search(boardCopy, depth-1, alpha, beta, timeStarted, totalMilliseconds)
		result.BestMove = &move

		if bestResult == nil {
			bestResult = &result
			continue
		}

		// if time.Since(timeStarted).Milliseconds() > totalMilliseconds {
		// 	break
		// }

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
