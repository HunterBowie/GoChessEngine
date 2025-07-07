package chess

import (
	"errors"
	"strings"
    "strconv"
)


func loadBoardFromFEN(fen string) (Board, error) {
	// splitting the FEN into parts
	parts := strings.Split(fen, " ")
	if len(parts) != 6 {
	  return errors.New("FEN loading error: number of parts is not 6")
	}
	piecePlacement := parts[0];
	activeColor := White;
    if parts[1] == "b" {
        activeColor = Black
    }
	castling := parts[2];
    if castling == "-" {
        castling = ""
    }
	enPassant := parts[3];
    if enPassant == "-" {
        enPassant = ""
    }
	halfMoves, err := strconv.Atoi(parts[4]);
    if err != nil {
        return err
    }
	fullMoves, err := strconv.Atoi(parts[5]);
    if err != nil {
        return err
    }
  
	board := Board{
        ActiveColor: activeColor,
        Castling: castling,
        EnPassant: enPassant,
        HalfMoves: halfMoves,
        FullMoves: fullMoves,
    }
  
	// adding the pieces
	rankNum = 8
	fileNum = 1
	for i := 0; i < len(piecePlacement); i++ {
	  char := piecePlacement[i]
  
	  if char == "/" {
		rankNum--
		fileNum = 1
		continue
	  }
  
	  skips, err = strconv.Atoi(char)
	  if err == nil {
		fileNum += skips;
		continue;
	  }
	  color := White
      lowerCaseChar = strings.ToLower(char)
      if lowerCaseChar == char {
        color := Black
      }
	  pieceType = 0;
	  switch (lowerCaseChar) {
		case "p":
		  pieceType = Pawn;
		  break;
		case "b":
		  pieceType = Bishop;
		  break;
		case "n":
		  pieceType = Knight;
		  break;
		case "r":
		  pieceType = Rook;
		  break;
		case "q":
		  pieceType = Queen;
		  break;
		case "k":
		  pieceType = King;
		  break;
		default:
		  return errors.New("FEN loading error: character " + char + " is incorrect");
	  }
	  setPiece(makePos(rankNum, fileNum), (color | pieceType) as Piece, board);
	  fileNum++;
	}
  
	return board;
  }