package server

import (
	"fmt"
	"net"
	"strings"

	"github.com/0xRuFFy/mapDB/internal/mapdb"
	"github.com/0xRuFFy/mapDB/internal/utils/globals"
)

type User struct {
	addr string
	conn *net.Conn
	db   *mapdb.Database
	exit bool
}

func NewUser(addr string, conn *net.Conn) *User {
	return &User{
		addr: addr,
		conn: conn,
		exit: false,
	}
}

func (u *User) Addr() string {
	return u.addr
}

func (u *User) Conn() net.Conn {
	return *u.conn
}

func (u *User) Disconnect() {
	u.exit = true
}

func (u *User) handle() {
	conn := *u.conn
	defer conn.Close()

	conn.Write([]byte("Welcome to mapDB!\n"))

	buf := make([]byte, 0, globals.BUFFER_SIZE) // this is a buffer to hold the data that is read from the connection
	tmp := make([]byte, globals.READ_BUFFER)    // this is a temporary buffer to read data from the connection

	for conn != nil {
		n, err := conn.Read(tmp)
		if err != nil {
			if err.Error() != "EOF" {
				logger.Debug("error was here")
				logger.Error(err.Error())
			}
			break
		}

		buf = append(buf, tmp[:n]...)
		if n < globals.READ_BUFFER {
			logger.Info(fmt.Sprintf("[%s] Received: %s", conn.RemoteAddr().String(), string(buf)))
			conn.Write([]byte("Message received.\n"))

			u.msgHandler(string(buf))
			if u.exit {
				break
			}

			buf = make([]byte, 0, globals.BUFFER_SIZE)
			tmp = make([]byte, globals.READ_BUFFER)
		}
	}

	logger.Info("Connection closed from " + conn.RemoteAddr().String())
}

func (u *User) msgHandler(msg string) {
	split := strings.Split(strings.Trim(msg, "\n"), " ")
	cmd := split[0]
	if handel, ok := commands[cmd]; ok {
		handel.Handler(u, split[1:])
	} else {
		u.Conn().Write([]byte("Invalid command.\n"))
	}

}
