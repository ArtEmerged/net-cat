package internal

import (
	"net"
	"sync"
)

const (
	DefPort       = "8989"
	IncorrectPort = "[USAGE]: ./TCPChat $port"
	dateFormat    = "2006-01-02 15:04:05"
	joinMsg       = " has joined our chat...\n"
	liveMsg       = " has left our chat...\n"
	nameMsg       = "[ENTER YOUR NAME]:"
	listenMsg     = "Listening on the port :"
	welcomeMsg    = "static/welcome.txt"
	fullConn      = "The server's full. Do you want to wait for someone to come out?\n"
)

type message struct {
	text string
	time string
}
// type user struct {
// 	name string
// 	msg  message
// }

type server struct {
	listen   net.Listener
	messages chan message
	users    map[string]net.Conn
	mu       sync.RWMutex
}
