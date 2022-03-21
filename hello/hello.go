package main

import (
	"example/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Amp")
	fmt.Println(message)
}
