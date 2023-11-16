package internal

import (
	"bufio"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
)

func (s *server) handler(conn net.Conn) {
	s.welcome(conn)
	name := s.getUserName(conn)
	s.usersNotification(conn, name)
	s.mu.Lock()
	s.users[name] = conn
	s.mu.Unlock()
	conn.Write([]byte(s.allmessages))
	s.client(conn, name)
}

func (s *server) client(conn net.Conn, name string) {
	defer s.closeConn(conn, name)
	var text string
	buf := bufio.NewScanner(conn)
	for buf.Scan() {
		text = buf.Text()
		msg := message{
			time: time.Now().Format(dateFormat),
			user: name,
			text: text,
		}
		s.messages <- msg
	}
}

func (s *server) welcome(conn net.Conn) {
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
		if len(name) > 8 {
			conn.Write([]byte(nameIncorr + nameMsg))
		} else if !s.checkName(name) {
			conn.Write([]byte(nameIncorr + nameMsg))
		} else if _, ok := s.users[name]; ok {
			conn.Write([]byte(nameUsed + nameMsg))
		} else {
			break
		}
	}
	return name
}

func (s *server) usersNotification(conn net.Conn, name string) {
	msg := message{
		text: joinMsg,
		user: name,
		time: "",
	}
	s.messages <- msg
}

func (s *server) closeConn(conn net.Conn, name string) {
	s.mu.Lock()
	defer conn.Close()
	defer s.mu.Unlock()
	msg := message{
		text: leftMsg,
		user: name,
		time: "",
	}
	delete(s.users, name)
	s.messages <- msg
}

func (s *server) checkName(name string) bool {
	pattern := regexp.MustCompile(`^[[:alpha:]]+$`)
	return pattern.MatchString(name)
}

func (s *server) checkEmpty(text string) bool {
	trimmedText := strings.TrimSpace(text)
	return len(trimmedText) == 0
}

func (s *server) checkRune(text string) bool {
	for _, letter := range text {
		if letter < 32 {
			return true
		}
	}
	return false
}
