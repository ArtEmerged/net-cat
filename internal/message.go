package internal

import (
	"fmt"
	"time"
)

func (s *server) GetMessage(msg message) {
	fmt.Println("GetMessage")
	s.messages <- msg
}

func (s *server) switchMsg() {
	for {
		select {
		case msg := <-s.messages:
			fmt.Println("switchMsg msg")
			s.write(msg)
		case <-s.live:
		}
	}
}

func (s *server) toString(msg message) string {
	fmt.Println("toString")
	if msg.time == "" {
		return fmt.Sprintf("%s%s", msg.user, msg.text)
	}
	text := fmt.Sprintf("\n[%s][%s]:%s", msg.time, msg.user, msg.text)
	// s.saveMessage(text)
	return text
}

func (s *server) write(msg message) {
	s.mu.Lock()
	fmt.Println("write:start")
	defer s.mu.Unlock()
	text := s.toString(msg)
	time := time.Now().Format(dateFormat)
	var datename string
	for name, userAddr := range s.users {
		datename = fmt.Sprintf("[%s][%s]:", time, name)
		userAddr.Write([]byte(datename))
		fmt.Println("write:for")
		if msg.user == name {
			continue
		}
		userAddr.Write([]byte(text))
		// userAddr.Write([]byte(datename))
	}
	fmt.Println("write:end")
}

func (s *server) saveMessage(msg string) {
	fmt.Println("write:saveMessage")
	s.allmessages += msg
}
