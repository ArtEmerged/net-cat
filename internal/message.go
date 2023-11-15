package internal

import (
	"fmt"
	"time"
)

func (s *server) getMessage(msg message) {
	if s.checkEmpty(msg.text) {
		return
	}
	s.messages <- msg
}

// func (s *server) switchMsg() {
// 	for {
// 		select {
// 		case msg := <-s.messages:
// 			s.write(msg)
// 		case <-s.live:
// 		}
// 	}
// }

func (s *server) write() {
	for {
		msg := <-s.messages
		s.mu.Lock()
		text := s.toString(msg)
		time := time.Now().Format(dateFormat)
		var datename string
		for name, userAddr := range s.users {
			datename = fmt.Sprintf("\n[%s][%s]:", time, name)
			if msg.user == name {
				userAddr.Write([]byte(datename[1:]))
				continue
			}
			userAddr.Write([]byte(text))
			userAddr.Write([]byte(datename))
		}
		s.mu.Unlock()
	}
}

func (s *server) toString(msg message) string {
	if msg.time == "" {
		return fmt.Sprintf("\n%s%s", msg.user, msg.text)
	}
	text := fmt.Sprintf("\n[%s][%s]:%s", msg.time, msg.user, msg.text)
	s.saveMessage(text[1:] + "\n")
	return text
}

func (s *server) saveMessage(msg string) {
	s.allmessages += msg
}
