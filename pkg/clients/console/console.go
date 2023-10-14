package console

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/kertox662/game-site/pkg/clients/command"
	"github.com/kertox662/game-site/pkg/clients/command/tictactoe"
	"github.com/kertox662/game-site/proto/games/tictactoe/tictactoeconnect"
)

type Console struct {
	cmd command.Command
	in  io.Reader
	out io.Writer
}

func NewConsole(
	client tictactoeconnect.TicTacToeServiceClient,
	in io.Reader,
	out io.Writer,
) *Console {
	return &Console{
		cmd: consoleCommand(client),
		in:  in,
		out: out,
	}
}

func (c *Console) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Exiting...")
			return
		default:
		}
		fmt.Print(">> ")
		scanner := bufio.NewScanner(c.in)
		if scanner.Scan() {
			line := scanner.Text()
			if line == "exit" {
				return
			}
			c.handleCommand(ctx, line)
		} else {
			fmt.Println("Got EOF, exiting...")
			return
		}
	}
}

func (c *Console) handleCommand(ctx context.Context, line string) {
	parts := strings.Split(line, " ")
	msg, err := c.cmd.Execute(ctx, parts)
	for _, m := range msg {
		_, _ = c.out.Write([]byte(m))
		_, _ = c.out.Write([]byte("\n"))
	}
	if err != nil {
		_, _ = c.out.Write([]byte(err.Error()))
		_, _ = c.out.Write([]byte("\n"))
	}
}

func consoleCommand(client tictactoeconnect.TicTacToeServiceClient) command.Command {
	root := &command.InteriorCommand{
		Name:        "",
		BaseCommand: *command.NewBaseCommand(),
	}

	tttCmd := &command.InteriorCommand{
		Name: "ttt",
	}

	printCmd := &command.InteriorCommand{
		Name: "print",
	}

	printGames := &tictactoe.PrintGamesCommand{}

	printGame := &tictactoe.PrintGameCommand{}

	listGames := &tictactoe.ListGamesCommand{}
	listGames.WithProtoClient(client)

	createGame := &tictactoe.CreateGameCommand{}
	createGame.WithProtoClient(client)

	playMove := &tictactoe.PlayMoveCommand{}
	playMove.WithProtoClient(client)

	getGame := &tictactoe.GetGameDataCommand{}
	getGame.WithProtoClient(client)

	return root.
		WithCommand("ttt",
			tttCmd.
				WithCommand("print",
					printCmd.
						WithCommand("games", printGames).
						WithCommand("game", printGame),
				).
				WithCommand("list", listGames).
				WithCommand("create", createGame).
				WithCommand("play", playMove).
				WithCommand("get", getGame),
		)
}
