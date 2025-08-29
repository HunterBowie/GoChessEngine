package minimax

import (
	"math/rand"

	"github.com/HunterBowie/GoChessEngine/internal/chess"
)

var openingWhiteMoves = [4]chess.Move{
	{Start: chess.LoadPos("d2"), End: chess.LoadPos("d4"), Flag: chess.PawnDoublePushFlag},
	{Start: chess.LoadPos("e2"), End: chess.LoadPos("e4"), Flag: chess.PawnDoublePushFlag},
	{Start: chess.LoadPos("c2"), End: chess.LoadPos("c4"), Flag: chess.PawnDoublePushFlag},
	{Start: chess.LoadPos("g1"), End: chess.LoadPos("f3"), Flag: chess.NoFlag},
}

// Returns a random opening move for white
func GetOpeningWhiteMove() chess.Move {
	index := rand.Intn(len(openingWhiteMoves))
	return openingWhiteMoves[index]
}
