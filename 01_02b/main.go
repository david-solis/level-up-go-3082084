package main

import (
	"log"
	"time"
	"strings"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	msgs := strings.Split(msg, " ")
	for _, e := range msgs {
			var word []string
			for i, c := range e {
				word = append(word, strings.Repeat(string(c), i + 1))
			}
			print(strings.Join(word, ""))
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}