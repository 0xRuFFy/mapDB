package main

import (
	"flag"

	"github.com/0xRuFFy/mapDB/internal/server"
	"github.com/0xRuFFy/mapDB/internal/utils/globals"
)

func main() {

	var host string
	flag.StringVar(&host, "host", globals.HOST, "host on which the server listents")

	flag.Parse()

	server := server.NewMapDBServer("8080", host)
	server.Serve()

}
