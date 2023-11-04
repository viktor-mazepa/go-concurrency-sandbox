package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {

	for {
		time.Sleep(6 * time.Second)
		ch <- "This is from server_1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server_2"
	}
}
func main() {
	fmt.Println("Select with channels")
	fmt.Println("--------------------")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		select {
		case s1 := <-channel1:
			fmt.Println("Case ONE:", s1)
		case s2 := <-channel1:
			fmt.Println("Case TWO:", s2)
		case s3 := <-channel2:
			fmt.Println("Case THREE:", s3)
		case s4 := <-channel2:
			fmt.Println("Case FOUR:", s4)
		}
	}
}
