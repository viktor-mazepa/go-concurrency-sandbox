package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)
	go shout(ping, pong)

	fmt.Println("Type something and press ENETR (enter Q to quit)")

	for {
		fmt.Print("-->")

		var userInput string

		_, _ = fmt.Scanln(&userInput)

		if "q" == strings.ToLower(userInput) {
			break
		}

		ping <- userInput

		response := <-pong

		fmt.Println("Response:", response)
	}

	fmt.Println("All done. CLosing channels")
	close(ping)
	close(pong)
}
