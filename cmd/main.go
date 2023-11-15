package main

import (
	"log"
	"net-cat/internal"
	"os"
)

func main() {
	port := internal.DefPort
	args := os.Args[1:]
	if len(args) > 1 {
		log.Fatalln(internal.IncorrectPort)
	}
	if len(args) == 1 {
		port = args[0]
		internal.VaidPort(port)
	}
	internal.StartServer(port)
}
