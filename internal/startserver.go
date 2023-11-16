package internal

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

func StartServer(port string) {
	li, errStart := net.Listen("tcp", ":"+port)
	if errStart != nil {
		log.Println(portUse)
		os.Exit(0)
	}
	defer li.Close()
	fmt.Println(listenMsg + port)
	s := &server{
		listen:   li,
		messages: make(chan message),
		users:    make(map[string]net.Conn),
		mu:       sync.RWMutex{},
	}
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
