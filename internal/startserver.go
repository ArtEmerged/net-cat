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
		log.Fatal(errStart)
	}
	defer li.Close()
	fmt.Println(listenMsg + port)
	s := &server{
		listen:   li,
		messages: make(chan message),
		users:    make(map[string]net.Conn),
		mu:       sync.RWMutex{},
	}
	go s.switchMsg()
	for {
		conn, errConn := s.listen.Accept()
		if errConn != nil {
			log.Fatal(errConn)
		}
		if len(s.users) > 10 {
			conn.Write([]byte(fullConn))
			for len(s.users) > 10 {
			}
		}
		go s.handler(conn)
	}
}
