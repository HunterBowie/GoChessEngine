package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/HunterBowie/GoChessEngine/internal/chess"
	"github.com/HunterBowie/GoChessEngine/internal/minimax"
	"github.com/gin-gonic/gin"
)

type BestMoveResponse struct {
	FEN      string `json:"fen"`
	BestMove string `json:"best_move"`
	MoveFlag int    `json:"move_flag"`
}

type EvalResponse struct {
	FEN  string `json:"fen"`
	Eval int    `json:"eval"`
}

// GetBotMove handles the bot best move generation requests
func GetBotMove(c *gin.Context) {
	fen := c.Query("fen")

	board := chess.LoadBoardFromFEN(fen)

	var bestMove string
	var flag int

	if board.FullMoves == 1 && board.ActiveColor == chess.White {
		move := minimax.GetOpeningWhiteMove()
		bestMove = chess.MoveToAlgebraic(move)
		flag = move.Flag
	} else {
		results := minimax.Search(board, 1)
		bestMove = chess.MoveToAlgebraic(*results.BestMove)
		flag = results.BestMove.Flag
	}

	output := BestMoveResponse{
		FEN:      fen,
		BestMove: bestMove,
		MoveFlag: flag,
	}

	c.IndentedJSON(http.StatusOK, output)
}

// GetBotEval handles the bot evaluation requests
func GetBotEval(c *gin.Context) {
	fen := c.Query("fen")

	board := chess.LoadBoardFromFEN(fen)

	score := minimax.Evaluate(board)

	output := EvalResponse{
		FEN:  fen,
		Eval: score,
	}
	c.IndentedJSON(http.StatusOK, output)

}

func main() {
	router := gin.Default()
	// CORS middleware
	// router.Use(func(c *gin.Context) {
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // allow all domains
	// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

	// 	if c.Request.Method == "OPTIONS" {
	// 		c.AbortWithStatus(200) // respond to preflight
	// 		return
	// 	}

	// 	c.Next()
	// })
	router.GET("/minimax/bestmove", GetBotMove)
	router.GET("/minimax/eval", GetBotEval)

	router.Run(":8080")
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
