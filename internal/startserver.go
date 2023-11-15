package internal

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func StartServer(port string) {
	li, errStart := net.Listen("tcp", ":"+port)
	if errStart != nil {
		log.Println(errStart)
	}

	fmt.Println(listenMsg + port)
	s := &server{
		listen:   li,
		messages: make(chan message),
		users:    make(map[string]net.Conn),
		mu:       sync.RWMutex{},
	}
	defer s.closeServer()
	go s.write()
	for {
		conn, errConn := s.listen.Accept()
		if errConn != nil {
			conn.Close()
			continue
		}
		if len(s.users) > 10 {
			conn.Write([]byte(fullConn))
			for len(s.users) > 10 {
			}
		}
		go s.handler(conn)
	}
}

func (s *server) closeServer() {
	defer s.listen.Close()
	for _, conn := range s.users {
		conn.Write([]byte(exitServer))
		conn.Close()
	}
}
