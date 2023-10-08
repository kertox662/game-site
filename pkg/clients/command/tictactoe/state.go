package tictactoe

import "github.com/kertox662/game-site/pkg/clients/command"

const (
	knownGamesKey = "knownGames"
	gameDataKey   = "gameData"
)

func SetInitialGameState(cs command.CommandState) {
	knownGames := make(map[string]gameMetadata)
	command.SetStateData(cs, knownGamesKey, knownGames)

}

type gameMetadata struct {
	CurrentPlayer int32
	PlayerCount   int32
	MaxPlayers    int32
	BoardSize     int32
	ConnectTarget int32
}

type gameData struct {
	Board [][]int32
	Chat  []string
}
