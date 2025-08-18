package minimax

import (
	"testing"

	"github.com/HunterBowie/GoChessEngine/internal/chess"
)

// TestHelloName calls minimax.Evaluate with the starting board postiion,
// checking for the correct score value.
func TestEvaluateStartingPos(t *testing.T) {
	// Parameters
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	expectedScore := 0

	// Test
	board := chess.LoadBoardFromFEN(fen)
	score := Evaluate(board)
	if score != expectedScore {
		t.Errorf(`Evaluate("%s") = %d want match for %d`, fen, score, expectedScore)
	}
}

// TestHelloName calls minimax.Evaluate with the starting position with
// only the white pawns, checking for the correct score value.
func TestEvaluateOnlyWhitePawns(t *testing.T) {
	// Parameters
	fen := "8/8/8/8/8/8/PPPPPPPP/4K3 w - - 0 1"
	whitePieces := PawnValue*8 + 10 + KingValue
	blackPieces := 0
	expectedScore := whitePieces + blackPieces

	// Test
	board := chess.LoadBoardFromFEN(fen)
	score := Evaluate(board)
	if score != expectedScore {
		t.Errorf(`Evaluate("%s") = %d want match for %d`, fen, score, expectedScore)
	}
}

// TestHelloName calls minimax.Evaluate with the starting position with
// only the white pieces, checking for the correct score value.
func TestEvaluateOnlyWhite(t *testing.T) {
	// Parameters
	fen := "4k3/8/8/8/8/8/PPPPPPPP/RNBQKBNR w - - 0 1"
	whitePieces := PawnValue*8 + 10 + KnightValue*2 - 80 + BishopValue*2 - 20 +
		RookValue*2 + QueenValue - 5 + KingValue
	blackPieces := 0
	expectedScore := whitePieces + blackPieces

	// Test
	board := chess.LoadBoardFromFEN(fen)
	score := Evaluate(board)
	if score != expectedScore {
		t.Errorf(`Evaluate("%s") = %d want match for %d`, fen, score, expectedScore)
	}
}

// TestHelloName calls minimax.Evaluate with the starting position with
// only the black pieces, checking for the correct score value.
func TestEvaluateOnlyBlack(t *testing.T) {
	// Parameters
	fen := "rnbqkbnr/pppppppp/8/8/8/8/8/4K3 w - - 0 1"
	whitePieces := 0
	blackPieces := -PawnValue*8 - 10 - KnightValue*2 + 80 - BishopValue*2 + 20 -
		RookValue*2 - QueenValue + 5 - KingValue
	expectedScore := whitePieces + blackPieces

	// Test
	board := chess.LoadBoardFromFEN(fen)
	score := Evaluate(board)
	if score != expectedScore {
		t.Errorf(`Evaluate("%s") = %d want match for %d`, fen, score, expectedScore)
	}
}

// TestHelloName calls minimax.Evaluate with the an advatages white position,
// checking for a valid return value.
func TestEvaluateWhiteAdvantage(t *testing.T) {
	// Parameters
	fen := "3rqrk1/ppp2ppp/2n2n2/2bppb2/8/8/PPPPPPPP/RNBQKBNR w KQ - 0 1"
	whitePieces := PawnValue*8 + 10 + KnightValue*2 - 80 + BishopValue*2 - 20 +
		RookValue*2 + QueenValue - 5 + KingValue
	blackPieces := -8*PawnValue - 90 - KnightValue*2 - 20 - BishopValue*2 - 20 -
		RookValue*2 - 5 - QueenValue + 5 - KingValue - 30
	expectedScore := whitePieces + blackPieces

	// Test
	board := chess.LoadBoardFromFEN(fen)
	score := Evaluate(board)
	if score != expectedScore {
		t.Errorf(`Evaluate("%s") = %d want match for %d`, fen, score, expectedScore)
	}
}
