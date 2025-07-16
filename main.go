package main

import (
	"fmt"
	"strings"

	"github.com/HunterBowie/GoChessEngine/internal/chess"
)

func main() {
	board := chess.LoadBoardFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PP3PPP/R3K2R w KQkq - 0 1")
	// board := chess.LoadBoardFromFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	// board.Add(chess.CreatePos(3, 5), chess.CreatePiece(chess.Rook|chess.Black))
	// board.PlayMove(chess.Move{Start: chess.CreatePos(1, 5), End: chess.CreatePos(2, 5), Flag: chess.BreaksCastlingRightsFlag})
	// board.PlayMove(chess.Move{Start: chess.CreatePos(7, 5), End: chess.CreatePos(6, 5), Flag: chess.NoFlag})
	// board.PlayMove(chess.Move{Start: chess.CreatePos(2, 5), End: chess.CreatePos(1, 5), Flag: chess.BreaksCastlingRightsFlag})
	// board.PlayMove(chess.Move{Start: chess.CreatePos(6, 5), End: chess.CreatePos(5, 5), Flag: chess.NoFlag})
	// printMoves(chess.GetMoves(board, false))
	fmt.Println("Starting...")
	fmt.Println()

	for _, move := range chess.GetMoves(board, false) {
		if move.Flag == chess.CastleKingsideFlag || move.Flag == chess.CastleQueensideFlag {
			fmt.Println("Contains castling move")
			board.PlayMove(move)
			break
		}
	}
	fmt.Println("end")
	printBoard(board)

	// for i := 0; i < 30; i++ {
	// 	time.Sleep(5 * time.Second)
	// 	moves := chess.GetMoves(board, false)
	// 	index := rand.Intn(len(moves))
	// 	board.PlayMove(moves[index])
	// 	printBoard(board)
	// }

	// time.Sleep(5 * time.Second)
	// board.PlayMove(chess.Move{Start: chess.CreatePos(7, 4), End: chess.CreatePos(8, 4), Flag: chess.PromoteToBishopFlag})
	// printBoard(board)
}

// printBoard prints a chess board to the console
func printBoard(board chess.Board) {
	for rank := 8; rank >= 1; rank-- {
		line := ""
		for file := 1; file <= 8; file++ {
			piece := board.Get(chess.CreatePos(rank, file))
			symbol := chess.PieceSymbol[piece.Type()]
			if piece.IsWhite() {
				symbol = strings.ToUpper(symbol)
			}
			line = line + symbol + " "
		}
		fmt.Println(line)
	}
	fmt.Println()
}

// printMoves prints a given list of moves to the console in readable form
func printMoves(moves []chess.Move) {
	for rank := 8; rank >= 1; rank-- {
		line := ""
		for file := 1; file <= 8; file++ {
			found := false
			for _, move := range moves {
				if move.End.Rank == rank && move.End.File == file {
					found = true
					break
				}
			}
			symbol := "-"
			if found {
				symbol = "X"
			}
			line = line + symbol + " "
		}
		fmt.Println(line)
	}
}
