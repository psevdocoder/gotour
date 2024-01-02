package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			<-time.After(200 * time.Millisecond)
			//time.Sleep(200 * time.Millisecond)
			c1 <- "one"
		}
	}()
	go func() {
		for {
			<-time.After(1 * time.Second)
			//time.Sleep(1 * time.Second)
			c2 <- "two"
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
