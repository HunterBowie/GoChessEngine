package minimax

import (
	"math/bits"

	"github.com/HunterBowie/GoChessEngine/internal/chess"
)

const (
	PawnValue   = 100
	KnightValue = 300
	BishopValue = 300
	RookValue   = 500
	QueenValue  = 900
)

// Top of table is bottom of board
// values repersent the bonus score obtained from
// the given piece occupying that square
// All tables are for white
var QueenTable = [64]int{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 10, 10, 10, 10, 0, 0,
	0, 0, 10, 10, 10, 10, 0, 0,
	0, 0, 10, 10, 10, 10, 0, 0,
	0, 0, 10, 10, 10, 10, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}

// Returns a score the reflects how good the position is for white
func Evaluate(board chess.Board) int {
	score := 0

	// white pieces
	score += PawnValue * getPieceCount(board, chess.CreatePiece(chess.White|chess.Pawn))
	score += KnightValue * getPieceCount(board, chess.CreatePiece(chess.White|chess.Knight))
	score += BishopValue * getPieceCount(board, chess.CreatePiece(chess.White|chess.Bishop))
	score += RookValue * getPieceCount(board, chess.CreatePiece(chess.White|chess.Rook))
	score += QueenValue * getPieceCount(board, chess.CreatePiece(chess.White|chess.Queen))

	// black pieces
	score -= PawnValue * getPieceCount(board, chess.CreatePiece(chess.Black|chess.Pawn))
	score -= KnightValue * getPieceCount(board, chess.CreatePiece(chess.Black|chess.Knight))
	score -= BishopValue * getPieceCount(board, chess.CreatePiece(chess.Black|chess.Bishop))
	score -= RookValue * getPieceCount(board, chess.CreatePiece(chess.Black|chess.Rook))
	score -= QueenValue * getPieceCount(board, chess.CreatePiece(chess.Black|chess.Queen))

	return score
}

// Returns the number of pieces of a certain type on the board
func getPieceCount(board chess.Board, piece chess.Piece) int {
	bitboard := board.Bitboards[chess.GetBitboardIndex(piece)]
	return bits.OnesCount64(bitboard)
}
