package server

import (
	"net"
	// "github.com/0xRuFFy/mapDB/internal/utils/errors"
)

const (
	NET_PROTOCOL = "tcp"
)

// MapDBServer is a server that handles requests from clients
// and manages the mapDB key-value store.
type MapDBServer struct {
	listener net.Listener
}

// NewMapDBServer creates a new MapDBServer instance.
// It takes a port, host, and protocol as arguments.
//
// The port is the port number that the server will listen on.
// The host is the host address that the server will listen on.
func NewMapDBServer(port, host string) *MapDBServer {
	logger.Info("Starting server on " + host + ":" + port + "...")
	listener, err := net.Listen(NET_PROTOCOL, host+":"+port)
	if err != nil {
		logger.Fatal(err.Error())
		return nil
	}

	return &MapDBServer{
		listener: listener,
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

	// TODO: implement me...

	logger.Info("Connection closed from " + conn.RemoteAddr().String())

	// panic(errors.NotYetImplemented())
}
