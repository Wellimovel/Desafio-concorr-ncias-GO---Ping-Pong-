package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}

}
func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func pingar2(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}

}

func main() {

	var c chan string = make(chan string)

	go pingar(c)
	go pingar2(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}
