package server

import (
	"fmt"
	"net"

	"github.com/0xRuFFy/mapDB/internal/mapdb"
	"github.com/0xRuFFy/mapDB/internal/utils/globals"
)

// MapDBServer is a server that handles requests from clients
// and manages the mapDB key-value store.
type MapDBServer struct {
	listener net.Listener

	db *mapdb.Database

	users []*User
}

// NewMapDBServer creates a new MapDBServer instance.
// It takes a port, host, and protocol as arguments.
//
// The port is the port number that the server will listen on.
// The host is the host address that the server will listen on.
func NewMapDBServer(port, host string) *MapDBServer {
	logger.Info("Starting server on " + host + ":" + port + "...")
	listener, err := net.Listen(globals.NET_PROTOCOL, host+":"+port)
	if err != nil {
		logger.Fatal(err.Error())
		return nil
	}

	return &MapDBServer{
		listener: listener,
		db:       mapdb.NewDatabase(),
		users:    make([]*User, 0),
	}
}

// Serve starts the server.
// It listens for connections and handles requests.
func (s *MapDBServer) Serve() {

	if s.listener == nil {
		logger.Fatal("Server listener is nil")
		return
	}

	defer s.listener.Close()
	logger.Info("Listening on " + s.listener.Addr().String())

	s.db.Set("test", "10")
	logger.Debug(fmt.Sprintf("Database: %v", s.db.Keys()))
	tmp, _ := s.db.Get("test")
	logger.Debug(fmt.Sprintf("%v", tmp))

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			logger.Error(err.Error())
			continue
		}

		go s.handleConnection(conn)
	}
}

// handleConnection handles a new connection to the server.
// Creates a User that will be used to handle requests.
func (s *MapDBServer) handleConnection(conn net.Conn) {
	logger.Info("New connection from " + conn.RemoteAddr().String())
	defer conn.Close()

	user := NewUser(conn.RemoteAddr().String(), &conn)
	s.users = append(s.users, user)
	user.handle()

	s.removeUser(user)
}

func (s *MapDBServer) removeUser(user *User) {
	for i, u := range s.users {
		if u == user {
			s.users = append(s.users[:i], s.users[i+1:]...)
			break
		}
	}
}
