package chess

import "strings"

// DATA DEFINITIONS

// Flag that is used to signal a special effect
const (
	NoFlag                   = 0
	EnPassantFlag            = 1
	CastleKingsideFlag       = 2
	CastleQueensideFlag      = 3
	PromoteToQueenFlag       = 4
	PromoteToRookFlag        = 5
	PromoteToBishopFlag      = 6
	PromoteToKnightFlag      = 7
	PawnDoublePushFlag       = 8
	BreaksCastlingRightsFlag = 9
)

// A chess move from one square to another with special effects recorded
type Move struct {
	Start Pos
	End   Pos
	Flag  int
}

// PUBLIC FUNCTION DEFINTIONS

func GetAllLegalMoves(board Board) []Move {
	return GetMoves(board, false, true, false)
}

// GetMoves returns all possible moves for the active color
func GetMoves(board Board, onlyAttacking bool, checkIllegal bool, onlyTaking bool) []Move {
	var moves []Move
	for rank := 1; rank <= 8; rank++ {
		for file := 1; file <= 8; file++ {
			pos := CreatePos(rank, file)
			piece := board.Get(pos)
			if piece.Color() != board.ActiveColor {
				continue
			}
			switch piece.Type() {
			case Pawn:
				moves = append(moves, getPawnMoves(board, pos, onlyAttacking)...)
			case Knight:
				moves = append(moves, getKnightMoves(board, pos, NoFlag)...)
			case Rook:
				moves = append(moves, getRookMoves(board, pos, BreaksCastlingRightsFlag)...)
			case Bishop:
				moves = append(moves, getBishopMoves(board, pos, NoFlag)...)
			case Queen:
				moves = append(moves, getRookMoves(board, pos, NoFlag)...)
				moves = append(moves, getBishopMoves(board, pos, NoFlag)...)
			case King:
				moves = append(moves, getKingMoves(board, pos, onlyAttacking)...)
			}

		}
	}

	if !checkIllegal && !onlyTaking {
		return moves
	}

	var filteredMoves []Move

	// check move legality and if take moves
	for _, move := range moves {
		boardCopy := board.Copy()
		boardCopy.PlayMove(move)
		boardCopy.changeActiveColor() // must be orginal color
		if onlyAttacking && boardCopy.Get(move.End) == None {
			continue
		}
		if checkIllegal && !IsKingInCheck(boardCopy) {
			filteredMoves = append(filteredMoves, move)
		}
	}

	return filteredMoves
}

// Generating pseudolegal moves

// Get pseudolegal pawn moves for a position on a given board
func getPawnMoves(board Board, pos Pos, onlyAttacking bool) []Move {
	var moves []Move

	direction := 1
	if board.ActiveColor == Black {
		direction = -1
	}

	blocked := false

	if !IsShiftIllegal(pos, direction, 0) {
		if board.Get(ShiftPos(pos, direction, 0)) != None {
			blocked = true
		}
	}

	flag := 0

	endRank := pos.Rank + direction

	if endRank == 8 || endRank == 1 {
		flag = PromoteToQueenFlag
	}

	// single pushes
	if !blocked {
		attachPawnMove(board, pos, &moves, direction, 0, false, flag, onlyAttacking)
	}
	attachPawnMove(board, pos, &moves, direction, 1, true, flag, onlyAttacking)
	attachPawnMove(board, pos, &moves, direction, -1, true, flag, onlyAttacking)

	// double push

	if !blocked {
		endRank = pos.Rank + direction*2
		if endRank == 5 && board.ActiveColor == Black ||
			endRank == 4 && board.ActiveColor == White {

			attachPawnMove(board, pos, &moves, direction*2, 0,
				false, PawnDoublePushFlag, onlyAttacking)
		}

	}

	// en passant

	if board.EnPassant != nil {
		enPassant := *board.EnPassant
		if enPassant.Rank == pos.Rank+direction {
			if enPassant.File == pos.File+1 || enPassant.File == pos.File-1 {
				moves = append(moves, Move{Start: pos, End: enPassant, Flag: EnPassantFlag})
			}
		}
	}

	return moves
}

// Attempts to add the given move to the list of moves if pseudolegal
func attachPawnMove(board Board, pos Pos, moves *[]Move, rankShift int, fileShift int, pawnAttack bool, flag int, onlyAttacking bool) {
	if !IsShiftIllegal(pos, rankShift, fileShift) {
		end := ShiftPos(pos, rankShift, fileShift)
		endPiece := board.Get(end)
		if endPiece == None && pawnAttack {
			return
		}
		if endPiece != None && !pawnAttack {
			return
		}
		if endPiece == None || (endPiece.Color() != board.ActiveColor) {
			if fileShift != 0 || !onlyAttacking {
				*moves = append(*moves, Move{Start: pos, End: end, Flag: flag})
			}
		}
	}
}

// Get pseudolegal king moves for a position on a given board
func getKingMoves(board Board, kingPos Pos, onlyAttacking bool) []Move {
	var moves []Move
	attachShiftedMove(board, kingPos, &moves, 0, 1, BreaksCastlingRightsFlag)
	attachShiftedMove(board, kingPos, &moves, 1, 0, BreaksCastlingRightsFlag)
	attachShiftedMove(board, kingPos, &moves, 1, 1, BreaksCastlingRightsFlag)
	attachShiftedMove(board, kingPos, &moves, 0, -1, BreaksCastlingRightsFlag)
	attachShiftedMove(board, kingPos, &moves, -1, 0, BreaksCastlingRightsFlag)
	attachShiftedMove(board, kingPos, &moves, -1, -1, BreaksCastlingRightsFlag)
	attachShiftedMove(board, kingPos, &moves, -1, 1, BreaksCastlingRightsFlag)
	attachShiftedMove(board, kingPos, &moves, 1, -1, BreaksCastlingRightsFlag)

	// castling
	if !onlyAttacking {
		kingside := "K"
		queenside := "Q"
		if board.ActiveColor == Black {
			kingside = "k"
			queenside = "q"
		}
		if strings.Contains(board.Castling, kingside) {
			attachKingsideCastle(board, kingPos, &moves)
		}
		if strings.Contains(board.Castling, queenside) {
			attachQueensideCastle(board, kingPos, &moves)
		}
	}

	return moves
}

// Attempts to add a kingside castle move if criteria is met
func attachKingsideCastle(board Board, kingPos Pos, moves *[]Move) {
	// we already know that the rook and king have not moved
	// check if pieces are in the way

	if IsShiftIllegal(kingPos, 0, 1) {
		return
	}

	if IsShiftIllegal(kingPos, 0, 2) {
		return
	}

	oneShift := ShiftPos(kingPos, 0, 1)
	twoShift := ShiftPos(kingPos, 0, 2)

	if board.Get(oneShift) != None {
		return
	}

	if board.Get(twoShift) != None {
		return
	}
	copiedBoard := board.Copy()
	copiedBoard.changeActiveColor()
	attackingMoves := GetMoves(copiedBoard, true, false, false)

	for _, move := range attackingMoves {
		switch move.End {
		case kingPos:
			return
		case oneShift:
			return
		case twoShift:
			return
		}
	}

	*moves = append(*moves, Move{Start: kingPos, End: twoShift, Flag: CastleKingsideFlag})

}

// Attempts to add a queenside castle move if criteria is met
func attachQueensideCastle(board Board, kingPos Pos, moves *[]Move) {
	// we already know that the rook and king have not moved
	// check if pieces are in the way

	if IsShiftIllegal(kingPos, 0, -1) {
		return
	}

	if IsShiftIllegal(kingPos, 0, -2) {
		return
	}

	if IsShiftIllegal(kingPos, 0, -3) {
		return
	}

	oneShift := ShiftPos(kingPos, 0, -1)
	twoShift := ShiftPos(kingPos, 0, -2)
	threeShift := ShiftPos(kingPos, 0, -3)

	if board.Get(oneShift) != None {
		return
	}

	if board.Get(twoShift) != None {
		return
	}

	if board.Get(threeShift) != None {
		return
	}

	copiedBoard := board.Copy()
	copiedBoard.changeActiveColor()
	attackingMoves := GetMoves(copiedBoard, true, false, false)

	for _, move := range attackingMoves {
		switch move.End {
		case kingPos:
			return
		case oneShift:
			return
		case twoShift:
			return
		case threeShift:
			return

		}
	}

	*moves = append(*moves, Move{Start: kingPos, End: twoShift, Flag: CastleQueensideFlag})

}

// Get pseudolegal knight moves for a position on a given board
func getKnightMoves(board Board, pos Pos, flag int) []Move {
	var moves []Move
	attachShiftedMove(board, pos, &moves, 1, 2, flag)
	attachShiftedMove(board, pos, &moves, -1, 2, flag)
	attachShiftedMove(board, pos, &moves, 1, -2, flag)
	attachShiftedMove(board, pos, &moves, -1, -2, flag)
	attachShiftedMove(board, pos, &moves, 2, 1, flag)
	attachShiftedMove(board, pos, &moves, -2, 1, flag)
	attachShiftedMove(board, pos, &moves, 2, -1, flag)
	attachShiftedMove(board, pos, &moves, -2, -1, flag)
	return moves
}

// Add a give move if it meets pseudolegal criteria and attacks enemy pieces
func attachShiftedMove(board Board, pos Pos, moves *[]Move, rankShift int, fileShift int, flag int) {
	if !IsShiftIllegal(pos, rankShift, fileShift) {
		end := ShiftPos(pos, rankShift, fileShift)
		endPiece := board.Get(end)
		if endPiece == None || (endPiece.Color() != board.ActiveColor) {
			*moves = append(*moves, Move{Start: pos, End: end, Flag: flag})
		}
	}
}

// Get pseudolegal bishop moves for a position on a given board
func getBishopMoves(board Board, pos Pos, flag int) []Move {
	var moves []Move
	moves = append(moves, getSlidingMovesInDirection(board, pos, 1, 1, flag)...)
	moves = append(moves, getSlidingMovesInDirection(board, pos, 1, -1, flag)...)
	moves = append(moves, getSlidingMovesInDirection(board, pos, -1, 1, flag)...)
	moves = append(moves, getSlidingMovesInDirection(board, pos, -1, -1, flag)...)
	return moves
}

// Get pseudolegal rook moves for a position on a given board
func getRookMoves(board Board, pos Pos, flag int) []Move {
	var moves []Move
	moves = append(moves, getSlidingMovesInDirection(board, pos, 1, 0, flag)...)
	moves = append(moves, getSlidingMovesInDirection(board, pos, 0, 1, flag)...)
	moves = append(moves, getSlidingMovesInDirection(board, pos, -1, 0, flag)...)
	moves = append(moves, getSlidingMovesInDirection(board, pos, 0, -1, flag)...)
	return moves
}

// Get pseudolegal sliding moves for the piece at the given position in given direction.
func getSlidingMovesInDirection(board Board, pos Pos, rankShift int, fileShift int, flag int) []Move {
	piece := board.Get(pos)
	var moves []Move

	if IsShiftIllegal(pos, rankShift, fileShift) {
		return moves
	}

	nextPos := ShiftPos(pos, rankShift, fileShift)
	for {
		nextPiece := board.Get(nextPos)
		if nextPiece != None && piece.HasSameColor(nextPiece) {
			break
		}
		moves = append(moves, Move{Start: pos, End: nextPos, Flag: flag})

		// pieces are different vision is blocked
		if nextPiece != None {
			break
		}
		// get next position in the sliding direction
		if IsShiftIllegal(nextPos, rankShift, fileShift) {
			return moves
		}
		nextPos = ShiftPos(nextPos, rankShift, fileShift)

	}
	return moves
}
