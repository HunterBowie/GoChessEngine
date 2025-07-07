package chess

// DATA DEFINITIONS

// A chess piece with a type and color
// null piece is 0
type Piece = int

// first three bits (what piece is this)
const None int = 0
const Pawn int = 1
const Bishop int = 2
const Knight int = 3
const Rook int = 4
const Queen int = 5
const King int = 6

// fourth bit (what color is this piece)
const White int = 0
const Black int = 8

// A position on a chess board with bounds [1, 8]
type Pos struct {
	Rank int
	File int
} 

// A move on a chess board
type Move struct {
	Start Pos
	End Pos
}

// Stores a board state (equivalent to FEN data)
type Board struct {
	Mailbox [64]int
	ActiveColor int
	Castling string 
	EnPassant Pos
	HalfMoves int
	FullMoves int
}

/*
Board

- Mailbox: [64]int
Pieces on the board stored in mailbox format. Starts at the bottom left
of the board and goes up left to right with increasing indicies.

- ActiveColor: int
Either 0 or 8 depending on which color is to move next on the board.

- Castling: string
Empty string if no castling is possible for either color. Contains upper case
letters K, Q for kingside and queenside castling with white. Lower case
for black.

- EnPassant: Pos
The square over which a pawn has just passed while moving two squares that
might be vulnerable to attack by en passant. Empty position else.

- HalfMoves: int
The number of moves since the last capture or pawn advance, used for the 
fifty move rule.

- FullMoves: int
The number of the full moves. It starts at 1 and is incremented after
Black's move.

*/

// PUBLIC FUNCTION DEFINITIONS


// GetColors returns the color value of the piece
func GetColor(piece int) int {
	return piece & Black
}

// GetType returns the type value of the piece
func GetType(piece int) int {
	return piece & 7
}


// Set replaces the position on the board with given piece
func (board *Board) Set(pos Pos, piece int8) {
}

// Get gets the piece from the board at row, col
func (board *Board) Get(pos Pos) int8 {
	return board.squares[calcIndex(pos)]
}


// PRIVATE FUNCTION DEFINITIONS


// calcIndex returns the index from topleft row, col position
func calcIndex(pos Pos) int {
	return 8*(7-int(pos.Row)) + int(pos.Col)
}