package main

import (
	"fmt"
	"time"
	"math/rand"
)

var messages = []string {"Not a chance", "Sounds good", "Yes", "No", "Think again", "Give me the antibiotics", "Things are looking up"}

func main() {
	for {
		fmt.Printf("%s\n", messages[rand.Intn(len(messages))])
		time.Sleep(800 * time.Millisecond)
	}
}
