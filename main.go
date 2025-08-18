//go:build wasm

package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall/js"

	"github.com/HunterBowie/GoChessEngine/internal/chess"
	"github.com/HunterBowie/GoChessEngine/internal/minimax"
)

// GetBotMove returns the move for given chess bot and position
// GetBotMove(bot string, settings int, time int, fen string) string
func GetBotMove(this js.Value, args []js.Value) interface{} {
	// bot := args[0].String()
	// settings := args[1].Int()
	// timeLeft := args[2].Int()
	fen := args[3].String()

	board := chess.LoadBoardFromFEN(fen)

	results := minimax.Search(board, 3)
	moveRaw := chess.MoveToAlgebraic(*results.BestMove)

	output := moveRaw + "-" + strconv.Itoa(results.BestMove.Flag)

	return output
}

// GetBotEval returns the evaluation for given chess bot and position
// GetBotEval(bot string, settings int, fen string) string
func GetBotEval(this js.Value, args []js.Value) interface{} {
	// bot := args[0].String()
	// settings := args[1].Int()
	fen := args[2].String()

	board := chess.LoadBoardFromFEN(fen)

	score := minimax.Evaluate(board)

	output := strconv.Itoa(score)

	return output
}

// Runs a program that makes GetBotMove visible in a web assembly file
func runWebAssembly() {
	c := make(chan struct{}, 0)
	js.Global().Set("GetBotMove", js.FuncOf(GetBotMove))
	js.Global().Set("GetBotEval", js.FuncOf(GetBotEval))
	<-c
}

func main() {
	runWebAssembly()
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

// // printMoves prints a given list of moves to the console in readable form
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
