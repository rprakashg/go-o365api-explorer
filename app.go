package main

import (
	"github.com/rprakashg/go-o365api-explorer/server"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

// main entry point.
func main() {
	server.Start()
}
