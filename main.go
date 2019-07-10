package main

import (
	"fmt"
	"time"
	"math/rand"
)

var messages = []string {"Not a chance", "Sounds good", "Yes", "No", "Think again"}

func main() {
	for {
		fmt.Printf("%s\n", messages[rand.Intn(len(messages))])
		time.Sleep(800 * time.Millisecond)
	}
}
