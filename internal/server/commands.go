package server

import (
	"fmt"
)

const (
	CMD_HELP   = "help"
	CMD_EXIT   = "exit"
	CMD_WHOAMI = "whoami"
)

type command interface {
	Name() string
	Description() string
	usage() (string, bool)
	Handler(*User, []string)
}

func usage() string {
	usage := "Usage: <command> [arguments]\n\nCommands:\n"
	for _, cmd := range commands {
		usage += fmt.Sprintf("  %s ~ %s\n", cmd.Name(), cmd.Description())
		cmdUsage, ok := cmd.usage()
		if ok {
			usage += fmt.Sprintf("    Usage: %s\n", cmdUsage)
		}
	}
	return usage
}

type helpCommand struct{}
type exitCommand struct{}
type whoamiCommand struct{}

// ------------------------- ExitCommand -------------------------
func (c *exitCommand) Name() string {
	return CMD_EXIT
}

func (c *exitCommand) Description() string {
	return "Exits the server."
}

func (c *exitCommand) usage() (string, bool) {
	return "", false
}

func (c *exitCommand) Handler(u *User, args []string) {
	u.Conn().Write([]byte("Bye!\n"))
	u.Disconnect()
}

// ------------------------- HelpCommand -------------------------
func (c *helpCommand) Name() string {
	return CMD_HELP
}

func (c *helpCommand) Description() string {
	return "Displays help information."
}

func (c *helpCommand) usage() (string, bool) {
	return "", false
}

func (c *helpCommand) Handler(u *User, args []string) {
	u.Conn().Write([]byte(usage()))
}

// ------------------------- WhoamiCommand -------------------------
func (c *whoamiCommand) Name() string {
	return CMD_WHOAMI
}

func (c *whoamiCommand) Description() string {
	return "Displays information about the current user."
}

func (c *whoamiCommand) usage() (string, bool) {
	return "", false
}

func (c *whoamiCommand) Handler(u *User, args []string) {
	u.Conn().Write([]byte(fmt.Sprintf("You are connected from %s\n", u.Addr())))
}

// ------------------------- Commands -------------------------
var commands = map[string]command{
	CMD_HELP:   &helpCommand{},
	CMD_EXIT:   &exitCommand{},
	CMD_WHOAMI: &whoamiCommand{},
}
