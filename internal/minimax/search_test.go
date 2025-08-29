package minimax

import (
	"fmt"
	"testing"
	"time"

	"github.com/HunterBowie/GoChessEngine/internal/chess"
)

// Test the time it takes to search at depth = 3
func TestTimeDepth3(t *testing.T) {
	board := chess.LoadBoardFromFEN("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1")

	now := time.Now()

	searchResults := Search(board, 1)

	passed := time.Since(now).Milliseconds()

	fmt.Println(passed)
	fmt.Println(searchResults.BestMove)

}
