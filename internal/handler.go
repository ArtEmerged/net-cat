package internal

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
)

func (s *server) client(conn net.Conn, name string) {
	fmt.Println("client->start")
	var text string
	buf := bufio.NewScanner(conn)
	// datename := fmt.Sprintf("[%s][%s]:", time.Now().Format(dateFormat), name)
	// conn.Write([]byte(datename))
	for buf.Scan() {
		fmt.Println("client->scan")
		text = buf.Text()
		if s.checkEmpty(text) {
			fmt.Println("client->scan->if == true")
			continue
		}
		msg := message{
			time: time.Now().Format(dateFormat),
			user: name,
			text: text,
		}
		fmt.Println("handler:client->scan->GetMessage-> start")
		s.GetMessage(msg)
		fmt.Println("handler:client->scan->GetMessage-> end")
		// datename = fmt.Sprintf("[%s][%s]:", time.Now().Format(dateFormat), name)
		// conn.Write([]byte(datename))
	}
	fmt.Println("client->end")
}

func (s *server) handler(conn net.Conn) {
	fmt.Println("handler:welcome-> start")
	s.welcome(conn)
	fmt.Println("handler:getNAme-> start")
	name := s.getUserName(conn)
	fmt.Println("handler:usersNotification-> start")
	s.usersNotification(conn, name)
	s.mu.Lock()
	s.users[name] = conn
	s.mu.Unlock()
	fmt.Println("handler:Write->allmessages-> start")
	conn.Write([]byte(s.allmessages))
	s.client(conn, name)
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
		if !s.checkName(name) {
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
	fmt.Println("usersNotification")
	msg := message{
		text: joinMsg,
		user: name,
		time: "",
	}
	s.GetMessage(msg)
}

func (s *server) checkName(name string) bool {
	pattern := regexp.MustCompile(`^[[:alpha:]]+$`)
	return pattern.MatchString(name)
}

func (s *server) checkEmpty(text string) bool {
	trimmedText := strings.TrimSpace(text)
	return len(trimmedText) == 0
}
