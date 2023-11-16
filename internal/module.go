package internal

import (
	"net"
	"sync"
)

const (
	DefPort       = "8989"
	portUse       = "This port is already in use"
	IncorrectPort = "[USAGE]: ./TCPChat $port"
	dateFormat    = "2006-01-02 15:04:05"
	joinMsg       = " has joined our chat..."
	leftMsg       = " has left our chat..."
	nameUsed      = "There's already a user with that name.\n"
	nameIncorr    = "Use only Latin letters.\n"
	nameVeryLong  = "Your name must not exceed 10 Latin characters.\n"
	nameMsg       = "[ENTER YOUR NAME]:"
	listenMsg     = "Listening on the port :"
	welcomeMsg    = "static/welcome.txt"
	fullConn      = "The server's full. Do you want to wait for someone to come out?\n"
	exitServer    = "Sorry. The server has been suspended"
)

type message struct {
	time string
	user string
	text string
}

type server struct {
	listen       net.Listener
	messages     chan message
	users        map[string]net.Conn
	allmessages  string
	mu           sync.RWMutex
}
