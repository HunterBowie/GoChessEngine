package minimax

import (
	"math"
	"math/bits"

	"github.com/HunterBowie/GoChessEngine/internal/chess"
)

const (
	PawnValue   = 100
	KnightValue = 300
	BishopValue = 300
	RookValue   = 500
	QueenValue  = 900
	KingValue   = 0
)

var pieceValues = map[int]int{
	chess.Pawn:   PawnValue,
	chess.Knight: KnightValue,
	chess.Bishop: BishopValue,
	chess.Rook:   RookValue,
	chess.Queen:  QueenValue,
	chess.King:   KingValue,
}

// Returns a score the reflects how good the position is for white
func Evaluate(board chess.Board) int {

	state := board.GetGameState()
	if state != chess.GamePlayState {
		if state == chess.GameWonState {
			if board.ActiveColor == chess.Black {
				// White Won
				return math.MaxInt
			}
			// Black Won
			return math.MinInt
		} else {
			// Tied
			return 0
		}
	}

	score := 0

	for _, color := range [2]int{chess.White, chess.Black} {
		for pieceType := chess.Pawn; pieceType <= chess.King; pieceType++ {
			score += evaluatePositions(board, chess.CreatePiece(color|pieceType))
		}
	}

	return score
}

// 101000

// Returns the number of pieces of a certain type on the board
func getPieceCount(board chess.Board, piece chess.Piece) int {
	bitboard := board.Bitboards[chess.GetBitboardIndex(piece)]
	return bits.OnesCount64(bitboard)
}

func getPieceSquareTableBonus(index int, piece chess.Piece) int {
	typeIndex := piece.Type() - 1
	if piece.IsBlack() {
		return pieceSquareTables[typeIndex][index]
	}
	index = 63 - index
	return pieceSquareTables[typeIndex][index]
}

// Returns total evaluation of a piece including table bonuses
func evaluatePositions(board chess.Board, piece chess.Piece) int {
	sign := 1
	if piece.IsBlack() {
		sign = -1
	}

	bitboard := board.Bitboards[chess.GetBitboardIndex(piece)]
	score := 0
	totalShifts := 0
	for bitboard != 0 {
		zeros := bits.TrailingZeros64(bitboard)
		totalShifts += zeros
		score += pieceValues[piece.Type()] * sign
		score += getPieceSquareTableBonus(totalShifts, piece) * sign
		bitboard = bitboard >> (zeros + 1)
		totalShifts += 1
	}
	return score
}
