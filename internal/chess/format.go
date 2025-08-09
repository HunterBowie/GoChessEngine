package chess

import "strconv"


func MoveToAlgebraic(move Move) string {
	return PosToAlgebraic(move.Start) + PosToAlgebraic(move.End)
}

func PosToAlgebraic(pos Pos) string {
	rankString := strconv.Itoa(pos.Rank)
	fileString := string(FileIndexes[pos.File - 1])
	return fileString + rankString
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
