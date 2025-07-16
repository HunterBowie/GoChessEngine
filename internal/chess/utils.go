package chess

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// DATA DEFINITIONS

var FileIndexes = [8]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}

// PUBLIC FUNCTION DEFINITIONS

// Returns the shifted given position by the rank and file numbers
func ShiftPos(pos Pos, rankShift int, fileShift int) Pos {
	rankNum := pos.Rank + rankShift
	fileNum := pos.File + fileShift
	if rankNum > 8 || fileNum > 8 || rankNum < 1 || fileNum < 1 {
		panic(fmt.Sprintf("The given rank or file shift is resulting in a shifted position of %d, %d ", rankNum, fileNum))
	}
	return CreatePos(rankNum, fileNum)
}

// returns true if the shift to the pos results in an illegal pos
func IsShiftIllegal(pos Pos, rankShift int, fileShift int) bool {
	newRank := pos.Rank + rankShift
	newFile := pos.File + fileShift
	if newRank <= 0 || newRank > 8 {
		return true
	}
	if newFile <= 0 || newFile > 8 {
		return true
	}
	return false
}

// Load an algebraic chess position into a Pos
func LoadPos(algebraic string) Pos {
	callPanic := func() {
		panic("Cannot load an algebraic position " + algebraic)
	}

	if len(algebraic) != 2 {
		callPanic()
	}
	fileChar := rune(algebraic[0])
	var file int
	found := false
	for index, char := range FileIndexes {
		if fileChar == char {
			file = index + 1
			found = true
			break
		}
	}
	if !found {
		callPanic()
	}

	rank := int(algebraic[1] - '0')

	return CreatePos(rank, file)
}

// Loads a board from a FEN string
func LoadBoardFromFEN(fen string) Board {
	callPanic := func(message string) {
		panic("Cannot load FEN position due to " + message)
	}

	// splitting the FEN into parts
	parts := strings.Split(fen, " ")
	if len(parts) != 6 {
		callPanic("the parts length")
	}
	piecePlacement := parts[0]
	activeColor := White
	if parts[1] == "b" {
		activeColor = Black
	}
	castling := parts[2]
	if castling == "-" {
		castling = ""
	}
	var enPassant *Pos
	if parts[3] != "-" {
		pos := LoadPos(parts[3])
		enPassant = &pos
	}
	halfMoves, err := strconv.Atoi(parts[4])
	if err != nil {
		callPanic(err.Error())
	}
	fullMoves, err := strconv.Atoi(parts[5])
	if err != nil {
		callPanic(err.Error())
	}

	board := Board{
		ActiveColor: activeColor,
		Castling:    castling,
		EnPassant:   enPassant,
		HalfMoves:   halfMoves,
		FullMoves:   fullMoves,
	}

	// adding the pieces
	rank := 8
	file := 1
	for i := 0; i < len(piecePlacement); i++ {
		char := piecePlacement[i]

		if char == '/' {
			rank--
			file = 1
			continue
		}

		if char > '0' && char < '9' {
			skips := int(char - '0')
			file += skips
			continue
		}

		color := White
		runeChar := rune(char)
		lowerCaseChar := unicode.ToLower(runeChar)
		if lowerCaseChar == runeChar {
			color = Black
		}
		pieceType := 0
		switch lowerCaseChar {
		case 'p':
			pieceType = Pawn
		case 'b':
			pieceType = Bishop
		case 'n':
			pieceType = Knight
		case 'r':
			pieceType = Rook
		case 'q':
			pieceType = Queen
		case 'k':
			pieceType = King
		default:
			callPanic(string(char) + " is an incorrect character")
		}
		board.Add(CreatePos(rank, file), CreatePiece(color|pieceType))
		file++
	}

	return board
}

// PRIVATE FUNCTION DEFINITIONS
