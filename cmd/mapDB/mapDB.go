package main

import "github.com/0xRuFFy/mapDB/internal/server"

func main() {

	server := server.NewMapDBServer("8080", "localhost")
	server.Serve()

}
