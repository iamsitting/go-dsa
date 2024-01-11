package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select { // there is no value in messages channel, so default is selected
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"

	select { // messages channel has no buffer or receiver, so default is selected
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Print("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
