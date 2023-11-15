package internal

import (
	"log"
)

func VaidPort(port string) {
	for _, digit := range port {
		if digit < 48 || digit > 57 {
			log.Fatalln(IncorrectPort)
		}
	}
}
