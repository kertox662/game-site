package tictactoe

import "github.com/kertox662/game-site/proto/games/tictactoe"

type board [][]int

// newBoard returns a new blank board, populated withs 0s.
// Requires positive size.
func newBoard(size int) board {
	b := make([][]int, 0, size)
	for i := 0; i < size; i++ {
		b = append(b, make([]int, size))
	}
	return b
}

func (b *board) makeMove(row, col, player int) {
	(*b)[row][col] = player
}

func (b board) isMoveInBounds(row, col int) bool {
	s := b.size()
	if row < 0 || col < 0 || row >= s || col >= s {
		return false
	}
	return true
}

func (b *board) size() int {
	return len(*b)
}

func (b board) findConnect(length int) (int, bool) {
	// checking is the player number currently currently found
	// found is the number of their symbols found
	var checking, found int

	// Check horizontal
	for r := 0; r < b.size(); r++ {
		checking = 0
		for c := 0; c < b.size(); c++ {
			if length-found > b.size()-c {
				break
			}

			if b[r][c] != checking {
				found = 1
				checking = b[r][c]
			} else {
				found++
			}

			if found == length && checking != 0 {
				return checking, true
			}
		}
	}

	// Check vertical
	for c := 0; c < b.size(); c++ {
		checking = 0
		for r := 0; r < b.size(); r++ {
			if length-found > b.size()-c {
				break
			}

			if b[r][c] != checking {
				found = 1
				checking = b[r][c]
			} else {
				found++
			}

			if found == length && checking != 0 {
				return checking, true
			}
		}
	}

	// Check TD,LR diagonal
	for ind := 0; ind <= b.size()-length; ind++ {
		checking = 0
		checking2 := 0
		found2 := 0
		for i, j := ind, 0; i < b.size(); i++ {
			if b[i][j] != checking {
				found = 1
				checking = b[i][j]
			} else {
				found++
			}

			if found == length && checking != 0 {
				return checking, true
			}

			if b[j][i] != checking2 {
				found2 = 1
				checking2 = b[j][i]
			} else {
				found2++
			}

			if found2 == length && checking2 != 0 {
				return checking2, true
			}

			j++
		}
	}

	// Check BU,LR diagonal
	for ind := 0; ind <= b.size()-length; ind++ {
		checking = 0
		checking2 := 0
		found2 := 0

		lRowInd := b.size() - 1 - ind
		lColInd := 0
		bRowInd := b.size() - 1
		bColInd := ind
		for {
			if lRowInd < 0 || bColInd >= b.size() {
				break
			}

			if b[lRowInd][lColInd] != checking {
				found = 1
				checking = b[lRowInd][lColInd]
			} else {
				found++
			}

			if found == length && checking != 0 {
				return checking, true
			}

			if b[bRowInd][bColInd] != checking2 {
				found2 = 1
				checking2 = b[bRowInd][bColInd]
			} else {
				found2++
			}

			if found2 == length && checking2 != 0 {
				return checking2, true
			}

			lRowInd--
			bRowInd--
			lColInd++
			bColInd++
		}
	}

	return 0, false
}

func (b *board) toProto() *tictactoe.Board {
	proto := &tictactoe.Board{
		Dimension: int32(b.size()),
		Cells:     make([]*tictactoe.BoardRow, 0, b.size()),
	}

	// Maps game board to proto board.
	for i := 0; i < b.size(); i++ {
		row := new(tictactoe.BoardRow)
		row.Cells = make([]int32, b.size())
		for j := 0; j < b.size(); j++ {
			row.Cells[j] = int32((*b)[i][j])
		}
		proto.Cells = append(proto.Cells, row)
	}

	return proto
}
