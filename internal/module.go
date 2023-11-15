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
	nameUsed      = "There's already a user with that name\n"
	nameIncorr    = "Use only Latin letters\n"
	nameMsg       = "[ENTER YOUR NAME]:"
	listenMsg     = "Listening on the port :"
	welcomeMsg    = "static/welcome.txt"
	fullConn      = "The server's full. Do you want to wait for someone to come out?\n"
)

type message struct {
	time string
	user string
	text string
}

// type user struct {
// 	name string
// 	msg  message
// }

type server struct {
	listen      net.Listener
	messages    chan message
	live        chan message
	users       map[string]net.Conn
	allmessages string
	mu          sync.RWMutex
}
