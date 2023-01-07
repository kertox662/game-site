package tictactoe

import (
	"errors"
	"strconv"

	"github.com/kertox662/game-site/proto/games/tictactoe"
)

const (
	EMPTY_PLAYER = 0 // Value for a blank space on the board
)

// Error values for broken rules
var (
	ErrIncorrectPlayerTurn = errors.New("not the given player's turn")
	ErrGameConcluded       = errors.New("game is already done")
	ErrOutOfBounds         = errors.New("move is out of bounds")
	ErrNonEmptySpace       = errors.New("move space in not empty")
)

// game is the data and configuration of the tic tac toe game
type game struct {
	board         board
	boardSize     int
	connectTarget int
	playerCount   int
	maxPlayers    int
	currentTurn   int
	winner        int
}

// newGame returns a new game.
// Requires positive boardSize, maxPlayers and connectTarget.
func newGame(boardSize, maxPlayers, connectTarget int) *game {
	return &game{
		board:         newBoard(boardSize),
		boardSize:     boardSize,
		maxPlayers:    maxPlayers,
		playerCount:   maxPlayers,
		connectTarget: connectTarget,
		currentTurn:   1,
		winner:        EMPTY_PLAYER,
	}
}

// toProto converts the game to a proto Game data response.
func (g *game) toProto() *tictactoe.GetGameDataResponse {
	// Temporarily will use player number for players
	players := make([]string, g.playerCount)
	for i := 0; i < g.playerCount; i++ {
		players[i] = strconv.Itoa(i + 1)
	}

	return &tictactoe.GetGameDataResponse{
		Data:        g.board.toProto(),
		PlayerCount: int32(g.playerCount),
		Players:     players,
		CurrentTurn: int32(g.currentTurn),
		Winner:      int32(g.winner),
	}
}

// makeMove makes the move and checks to see if the move breaks any rules
// If it does, it returns an error
func (g *game) makeMove(player, row, col int) error {
	if g.winner != EMPTY_PLAYER { // Check if winner exists
		return ErrGameConcluded
	} else if g.currentTurn != player { // Check if given player's turn
		return ErrIncorrectPlayerTurn
	} else if !g.board.isMoveInBounds(row, col) { // Check if move in bounds
		return ErrOutOfBounds
	} else if g.board[row][col] != EMPTY_PLAYER { // Check empty space
		return ErrNonEmptySpace
	}

	g.board.makeMove(row, col, player)
	g.winner = g.checkWinner()
	g.currentTurn = (g.currentTurn % g.playerCount) + 1 // Increment turn
	return nil
}

// checkWinner checks if the game has a winner, and returns
// their player number. If no winner, return the empty player.
func (g *game) checkWinner() int {
	if player, ok := g.board.findConnect(g.connectTarget); ok {
		return player
	}
	return EMPTY_PLAYER
}
