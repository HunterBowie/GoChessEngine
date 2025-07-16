package chess

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// DATA DEFINITIONS

// A chess piece with a type and color
// null piece is 0
type Piece int

const None Piece = 0

// first three bits (what piece is this)
const Pawn int = 1
const Bishop int = 2
const Knight int = 3
const Rook int = 4
const Queen int = 5
const King int = 6

// fourth bit (what color is this piece)
const White int = 0
const Black int = 8

var PieceSymbol = map[int]string{Pawn: "p", Bishop: "b", Knight: "n", Rook: "r", Queen: "q", King: "k", 0: "-"}

// A position on a chess board with bounds [1, 8]
type Pos struct {
	Rank int
	File int
}

// Stores a board state (equivalent to FEN data)
type Board struct {
	Bitboards   [12]uint64
	ActiveColor int
	Castling    string
	EnPassant   *Pos
	HalfMoves   int
	FullMoves   int
}

/*
Board

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

// BOARD FUNCTIONS

// Replaces the position on the board with given piece
func (board *Board) Add(pos Pos, piece Piece) {
	if board.Get(pos) != None {
		board.Remove(pos)
	}

	board.Bitboards[GetBitboardIndex(piece)] |= CalcBitboard(pos)
}

// Gets the piece from the board at row, col
func (board *Board) Get(pos Pos) Piece {
	bitboard := CalcBitboard(pos)
	for index := 0; index < 12; index++ {
		if board.Bitboards[index]&bitboard != 0 {
			return GetPieceFromIndex(index)
		}
	}
	return None
}

// Removes the piece from the board at row, col
func (board *Board) Remove(pos Pos) Piece {
	bitboard := CalcBitboard(pos)
	for index := 0; index < 12; index++ {
		if board.Bitboards[index]&bitboard != 0 {
			piece := GetPieceFromIndex(index)
			board.Bitboards[index] &= ^bitboard
			return piece
		}
	}
	return None
}

// Play the given move on the board
func (board *Board) PlayMove(move Move) {
	piece := board.Remove(move.Start)

	if piece.Color() != board.ActiveColor {
		panic("Attempting to play a move with the wrong color piece")
	}

	rank := 1
	if board.ActiveColor == Black {
		rank = 8
	}	

	switch move.Flag {
	case BreaksCastlingRightsFlag:
		if piece.Type() == Rook {
			
			leftPos := CreatePos(rank, 1)
			rightPos := CreatePos(rank, 8)

			if move.Start == leftPos {
				board.removeQueensideCastlingRights()
			} else if move.Start == rightPos {
				board.removeKingsideCastlingRights()
			}
		} else {
			board.removeAllCastlingRights()
		}
		board.Add(move.End, piece)
	case PromoteToKnightFlag:
		board.Add(move.End, CreatePiece(Knight|board.ActiveColor))
	case PromoteToBishopFlag:
		board.Add(move.End, CreatePiece(Bishop|board.ActiveColor))
	case PromoteToRookFlag:
		board.Add(move.End, CreatePiece(Rook|board.ActiveColor))
	case PromoteToQueenFlag:
		board.Add(move.End, CreatePiece(Queen|board.ActiveColor))
	case CastleKingsideFlag:
		board.Add(move.End, piece)
		rook := board.Remove(CreatePos(rank, 8))
		board.Add(ShiftPos(move.Start, 0, 1), rook)
		board.removeAllCastlingRights()
	case CastleQueensideFlag:
		board.Add(move.End, piece)
		rook := board.Remove(CreatePos(rank, 1))
		board.Add(ShiftPos(move.Start, 0, -1), rook)
		board.removeAllCastlingRights()
	default:
		board.Add(move.End, piece)

	}
	board.changeActiveColor()

}

func (board *Board) Copy() Board {
	newBoard := Board{
		Bitboards:   board.Bitboards,
		ActiveColor: board.ActiveColor,
		Castling:    board.Castling,
		EnPassant:   board.EnPassant, // TODO: FIX FUTURE BUG
		HalfMoves:   board.HalfMoves,
		FullMoves:   board.FullMoves,
	}
	return newBoard
}

// Flips the current active color
func (board *Board) changeActiveColor() {
	if board.ActiveColor == White {
		board.ActiveColor = Black
	} else {
		board.ActiveColor = White
	}
}

func (board *Board) removeAllCastlingRights() {
	board.removeKingsideCastlingRights()
	board.removeQueensideCastlingRights()
}

func (board *Board) removeKingsideCastlingRights() {
	if board.ActiveColor == White {
		board.Castling = strings.ReplaceAll(board.Castling, "K", "")
	} else {
		board.Castling = strings.ReplaceAll(board.Castling, "k", "")
	}
}

func (board *Board) removeQueensideCastlingRights() {
	if board.ActiveColor == White {
		board.Castling = strings.ReplaceAll(board.Castling, "Q", "")
	} else {
		board.Castling = strings.ReplaceAll(board.Castling, "q", "")
	}
}

// PIECE FUNCTIONS

// Create a new Piece with proper error handling
func CreatePiece(pieceValue int) Piece {
	validPiece := false
	if 0 <= pieceValue && pieceValue <= 6 {
		validPiece = true
	} else if 9 <= pieceValue && pieceValue <= 14 {
		validPiece = true
	}
	if validPiece {
		return Piece(pieceValue)
	}
	panic("New Piece cannot be made from value " + strconv.Itoa(pieceValue))
}

// Returns true if the piece is white
func (piece *Piece) IsWhite() bool {
	return int(*piece)&0b1000 == 0
}

// Returns true if the piece is black
func (piece *Piece) IsBlack() bool {
	return !piece.IsWhite()
}

// Returns piece color value
func (piece *Piece) Color() int {
	return int(*piece) & 0b1000
}

// Returns piece integer value
func (piece *Piece) Type() int {
	return int(*piece) & 0b111
}

// Returns true if the given piece has the same color as this one
func (piece *Piece) HasSameColor(otherPiece Piece) bool {
	return piece.Color() == otherPiece.Color()
}

// Pos functions

// Create a new Pos with proper error handling
func CreatePos(rank int, file int) Pos {
	callPanic := func() {
		panic("New Pos cannot be made from rank: " + strconv.Itoa(rank) + " file: " + strconv.Itoa(file))
	}

	if rank < 1 || rank > 8 {
		callPanic()
	}

	if file < 1 || file > 8 {
		callPanic()
	}

	return Pos{Rank: rank, File: file}

}

// Bitboard functions

// calcBitboardPos returns a bitboard containing the given pos
func CalcBitboard(pos Pos) uint64 {
	answer := uint64(1) << (8*(pos.Rank-1) + pos.File - 1)

	callPanic := func() {
		fmt.Println(pos.Rank)
		fmt.Println(pos.File)
		panic(fmt.Sprintf("The calculated bitbord position: %d, is invalid", answer))
	}

	if answer > uint64(math.Ceil(math.Pow(2, 64)-1)-(math.Pow(2, 63)-1)) {
		callPanic()
	}
	return uint64(answer)
}

// Returns the index of the bitboard containing board information for that piece
func GetBitboardIndex(piece Piece) int {
	colorIndex := 0
	if piece.IsBlack() {
		colorIndex = 6
	}
	return (piece.Type() - 1) + colorIndex
}

// Returns the piece associated with the bitboard index
func GetPieceFromIndex(index int) Piece {
	if index >= 12 || index < 0 {
		panic(fmt.Sprintf("Cannot calculate piece from invalid index: %d", index))
	}
	pieceColor := White
	pieceType := index + 1
	if index > 5 {
		pieceColor = Black
		pieceType = index - 5
	}
	return CreatePiece(pieceType | pieceColor)
}


