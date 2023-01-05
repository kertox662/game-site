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

// makeMove allows the player to make a move
func (b *board) makeMove(row, col, player int) {
	(*b)[row][col] = player
}

// isMoveInBounds checks if the the move is within the board
func (b board) isMoveInBounds(row, col int) bool {
	s := b.size()

	// Check the move to see if row/col are negative or above the maximum index
	if row < 0 || col < 0 || row >= s || col >= s {
		return false
	}
	// If the move is within the board, return true
	return true
}

// size returns the size of the board
func (b *board) size() int {
	return len(*b)
}

// findConnect identifies whether a winning connection has been made
func (b board) findConnect(length int) (int, bool) {
	// Checking is the player number currently currently found
	// Found is the number of their symbols found
	var checking, found int

	// Check horizontal
	for r := 0; r < b.size(); r++ {

		//Checking resets to 0 to once again look for an empty spot first when switched to a new row
		checking = 0
		for c := 0; c < b.size(); c++ {

			// If a full connection isn't possible, stop checking that row
			if length-found > b.size()-c {
				break
			}

			// If the space on the board is different from what was checked previously, set found as one
			// Also set checking as that players number where all future checks will look for that number
			// If not, increase the number of the player-numbers found
			if b[r][c] != checking {
				found = 1
				checking = b[r][c]
			} else {
				found++
			}

			// If enough non-empty consecutive symbols have been found,
			// Return the player number and that a connection has been made
			if found == length && checking != 0 {
				return checking, true
			}
		}
	}

	// Check vertical
	for c := 0; c < b.size(); c++ {
		checking = 0
		for r := 0; r < b.size(); r++ {
			// If a full connection isn't possible, stop checking that column
			if length-found > b.size()-r {
				break
			}

			if b[r][c] != checking {
				found = 1
				checking = b[r][c]
			} else {
				found++
			}

			// Returning the player number and that a connection has been made
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

			// If the space on the board is different from what was checked previously, set found as one
			// Also set checking as that players number where all future checks will look for that number
			// If not, increase the number of the player-numbers found
			if b[i][j] != checking {
				found = 1
				checking = b[i][j]
			} else {
				found++
			}

			// If enough non-empty consecutive symbols have been found,
			// Return the player number and that a connection has been made
			if found == length && checking != 0 {
				return checking, true
			}

			// If the space on the board is different from what was checked previously in check2, set found2 as one
			// Also set checking2 as that players number where all future checks will look for that number
			// If not, increase the number of the player-numbers found(2)
			if b[j][i] != checking2 {
				found2 = 1
				checking2 = b[j][i]
			} else {
				found2++
			}

			// Return the player number and that a connection has been made
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
			// If a connection is not possible
			if lRowInd < 0 || bColInd >= b.size() {
				break
			}

			if b[lRowInd][lColInd] != checking {
				found = 1
				checking = b[lRowInd][lColInd]
			} else {
				found++
			}

			// Return the player number and that a connection has been made
			if found == length && checking != 0 {
				return checking, true
			}

			if b[bRowInd][bColInd] != checking2 {
				found2 = 1
				checking2 = b[bRowInd][bColInd]
			} else {
				found2++
			}

			// Return the player number and that a connection has been made
			if found2 == length && checking2 != 0 {
				return checking2, true
			}

			// Increment statements for each row/column index
			lRowInd--
			bRowInd--
			lColInd++
			bColInd++
		}
	}

	// If no connection has been made, return blank value and that no connection is possible
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
		row.Length = int32(b.size())
		for j := 0; j < b.size(); j++ {
			row.Cells[j] = int32((*b)[i][j])
		}
		proto.Cells = append(proto.Cells, row)
	}

	return proto
}
