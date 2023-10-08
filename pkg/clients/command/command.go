package command

import (
	"context"
	"fmt"
	"strings"
)

type Command interface {
	Execute(ctx context.Context, args []string) ([]string, error)
	Help() string
	State() CommandState // command name -> []game Ids
	SetParent(parent Command)
}

type BaseCommand struct {
	parent Command
	state  CommandState
}

func NewBaseCommand() *BaseCommand {
	return &BaseCommand{
		state: make(CommandState),
	}
}

func (bc *BaseCommand) State() CommandState {
	if bc.parent != nil {
		return bc.parent.State()
	}
	return bc.state
}

func (bc *BaseCommand) SetParent(parent Command) {
	bc.parent = parent
}

type InteriorCommand struct {
	BaseCommand
	Name     string
	Commands map[string]Command
}

func (ic *InteriorCommand) Execute(ctx context.Context, args []string) ([]string, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf(ic.Help())
	}

	cmd, ok := ic.Commands[args[0]]
	if !ok {
		return []string{ic.Help()}, fmt.Errorf("command %s not found", args[0])
	}

	return cmd.Execute(ctx, args[1:])
}

func (ic *InteriorCommand) Help() string {
	subCommands := ""
	spacing := "  "
	for name, cmd := range ic.Commands {
		subParts := strings.Split(cmd.Help(), "\n")
		sub := ""
		for _, s := range subParts {
			sub += fmt.Sprintf("%s%s\n", spacing, s)
		}
		subCommands += fmt.Sprintf("%s: %s", name, sub)
	}
	return fmt.Sprintf("%s\n%s", ic.Name, subCommands)
}

func (ic *InteriorCommand) WithCommand(name string, cmd Command) *InteriorCommand {
	if ic.Commands == nil {
		ic.Commands = make(map[string]Command)
	}
	ic.Commands[name] = cmd
	cmd.SetParent(ic)
	return ic
}

func NewInteriorCommand() *InteriorCommand {
	return &InteriorCommand{
		Commands: make(map[string]Command),
	}
}
