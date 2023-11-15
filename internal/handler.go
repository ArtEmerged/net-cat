package internal

import (
	"bufio"
	"log"
	"net"
	"os"
)

func (s *server) handler(conn net.Conn) {
	s.welcom(conn)
	name := s.getUserName(conn)
}

func (s *server) welcom(conn net.Conn) {
	welcom, err := os.ReadFile(welcomeMsg)
	if err != nil {
		log.Println(err)
	}
	conn.Write(welcom)
	conn.Write([]byte(nameMsg))

}
func (s *server) getUserName(conn net.Conn) string {
	var name string
	buf := bufio.NewScanner(conn)
	for buf.Scan() {
		name = buf.Text()
	}
	return name
}

func (s *server) usersNotification(conn net.Conn, name string) {
	s.users[name] = conn
	
}
