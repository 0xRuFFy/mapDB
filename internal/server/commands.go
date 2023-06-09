package server

import (
	"fmt"
	"sort"
)

const (
	CMD_HELP   = "help"
	CMD_EXIT   = "exit"
	CMD_WHOAMI = "whoami"
	CMD_GET    = "get"
	CMD_SET    = "set"
	CMD_DEL    = "del"
	CMD_KEYS   = "keys"
)

type command interface {
	Name() string
	Description() string
	Access() AccessPrivilege
	usage() (string, bool)
	Handler(*User, []string)
}

func usage(user User) string {
	usage := "Usage: <command> [arguments]\n\nCommands:\n"
	keys := make([]string, 0, len(commands))
	for k := range commands {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		cmd := commands[key]
		usage += fmt.Sprintf("  %s", cmd.Name())
		if cmd.Access().CheckAccess(user) {
			usage += "    [insufficient privilege]"
		}
		usage += fmt.Sprintf("\n    ~ %s\n", cmd.Description())
		cmdUsage, ok := cmd.usage()
		if ok {
			usage += fmt.Sprintf("        Usage: %s\n", cmdUsage)
		}
	}
	return usage
}

type helpCommand struct{}
type exitCommand struct{}
type whoamiCommand struct{}
type getCommand struct{}
type setCommand struct{}
type delCommand struct{}
type keysCommand struct{}

// ------------------------- ExitCommand -------------------------
func (c *exitCommand) Name() string {
	return CMD_EXIT
}

func (c *exitCommand) Description() string {
	return "Exits the server."
}

func (c *exitCommand) Access() AccessPrivilege {
	return NoAccess
}

func (c *exitCommand) usage() (string, bool) {
	return "", false
}

func (c *exitCommand) Handler(u *User, args []string) {
	if c.Access().CheckAndFeedbackAccess(*u) {
		return
	}

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

func (c *helpCommand) Access() AccessPrivilege {
	return NoAccess
}

func (c *helpCommand) usage() (string, bool) {
	return "", false
}

func (c *helpCommand) Handler(u *User, args []string) {
	if c.Access().CheckAndFeedbackAccess(*u) {
		return
	}

	u.Conn().Write([]byte(usage(*u)))
}

// ------------------------- WhoamiCommand -------------------------
func (c *whoamiCommand) Name() string {
	return CMD_WHOAMI
}

func (c *whoamiCommand) Description() string {
	return "Displays information about the current user."
}

func (c *whoamiCommand) Access() AccessPrivilege {
	return NoAccess
}

func (c *whoamiCommand) usage() (string, bool) {
	return "", false
}

func (c *whoamiCommand) Handler(u *User, args []string) {
	if c.Access().CheckAndFeedbackAccess(*u) {
		return
	}

	u.Conn().Write([]byte(fmt.Sprintf("You are connected from %s\n", u.Addr())))
}

// ------------------------- GetCommand -------------------------
func (c *getCommand) Name() string {
	return "get"
}

func (c *getCommand) Description() string {
	return "Gets a key from the database."
}

func (c *getCommand) Access() AccessPrivilege {
	return ReadAccess
}

func (c *getCommand) usage() (string, bool) {
	return "get <key>", true
}

func (c *getCommand) Handler(u *User, args []string) {
	if c.Access().CheckAndFeedbackAccess(*u) {
		return
	}

	if len(args) != 1 {
		u.Conn().Write([]byte("Invalid arguments.\n"))
		return
	}

	key := args[0]
	value, err := u.db.Get(key)
	if err != nil {
		u.Conn().Write([]byte(err.Error() + "\n"))
		return
	}

	u.Conn().Write([]byte(fmt.Sprintf("%s\n", value)))
}

// ------------------------- SetCommand -------------------------
func (c *setCommand) Name() string {
	return "set"
}

func (c *setCommand) Description() string {
	return "Sets a key in the database."
}

func (c *setCommand) Access() AccessPrivilege {
	return ReadWriteAccess
}

func (c *setCommand) usage() (string, bool) {
	return "set <key> <value>", true
}

func (c *setCommand) Handler(u *User, args []string) {
	if c.Access().CheckAndFeedbackAccess(*u) {
		return
	}

	if len(args) != 2 {
		u.Conn().Write([]byte("Invalid arguments.\n"))
		return
	}

	key := args[0]
	value := args[1]
	err := u.db.Set(key, value)
	if err != nil {
		u.Conn().Write([]byte(err.Error() + "\n"))
		return
	}

	u.Conn().Write([]byte("OK\n"))
}

// ------------------------- DelCommand -------------------------
func (c *delCommand) Name() string {
	return "del"
}

func (c *delCommand) Description() string {
	return "Deletes a key from the database."
}

func (c *delCommand) Access() AccessPrivilege {
	return ReadWriteAccess
}

func (c *delCommand) usage() (string, bool) {
	return "del <key>", true
}

func (c *delCommand) Handler(u *User, args []string) {
	if c.Access().CheckAndFeedbackAccess(*u) {
		return
	}

	if len(args) != 1 {
		u.Conn().Write([]byte("Invalid arguments.\n"))
		return
	}

	key := args[0]
	err := u.db.Delete(key)
	if err != nil {
		u.Conn().Write([]byte(err.Error() + "\n"))
		return
	}

	u.Conn().Write([]byte("OK\n"))
}

// ------------------------- KeysCommand -------------------------
func (c *keysCommand) Name() string {
	return "keys"
}

func (c *keysCommand) Description() string {
	return "Lists all keys in the database."
}

func (c *keysCommand) Access() AccessPrivilege {
	return ReadAccess
}

func (c *keysCommand) usage() (string, bool) {
	return "", false
}

func (c *keysCommand) Handler(u *User, args []string) {
	if c.Access().CheckAndFeedbackAccess(*u) {
		return
	}

	keys := u.db.Keys()

	u.Conn().Write([]byte(fmt.Sprintf("%v\n", keys)))
}

// ------------------------- Commands -------------------------
var commands = map[string]command{
	CMD_HELP:   &helpCommand{},
	CMD_EXIT:   &exitCommand{},
	CMD_WHOAMI: &whoamiCommand{},
	CMD_GET:    &getCommand{},
	CMD_SET:    &setCommand{},
	CMD_DEL:    &delCommand{},
	CMD_KEYS:   &keysCommand{},
}
