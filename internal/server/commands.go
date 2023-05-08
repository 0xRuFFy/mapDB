package server

import "fmt"

const (
	CMD_HELP   = "help"
	CMD_EXIT   = "exit"
	CMD_WHOAMI = "whoami"
)

type command interface {
	Name() string
	Description() string
	Handler(*User, []string)
}

func usage() string {
	return fmt.Sprintf(
		"Usage: <command> [arguments]\n\nCommands:\n%s ~ %s\n%s ~ %s\n",
		CMD_EXIT,
		commands[CMD_EXIT].Description(),
		CMD_HELP,
		commands[CMD_HELP].Description(),
	)
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

func (c *helpCommand) Handler(u *User, args []string) {
	u.Conn().Write([]byte("Commands:\n"))
	u.Conn().Write([]byte(usage()))
}

// ------------------------- WhoamiCommand -------------------------
func (c *whoamiCommand) Name() string {
	return CMD_WHOAMI
}

func (c *whoamiCommand) Description() string {
	return "Displays information about the current user."
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
