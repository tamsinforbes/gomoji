package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return rand.Intn(max+1-min) + min
}

func randEmoji(block string) int {
	// Grinning Face emoji
	u := 0x1F600

	if block == "chess" {
		// Western chess symbols U+2654 to U+265F
		u = randInt(0x2654, 0x265F)
	} else if block == "symbols" {
		u = randInt(0x1FA70, 0x1FAF8)
		// Symbols and Pictographs Extended A U+1FA70 to U+1FAF8, some gaps
	} else if block == "faces" {
		// Face emojis U+1F600 to U+1F64E
		u = randInt(0x1F600, 0x1F64E)
	} else if block == "alchemical" {
		// Alchemical symbols
		u = randInt(0x1F700, 0x1F774)
	} else {
		u = 0x1F602
	}

	return u
}

func main() {
	name := flag.String("name", "world", "The name to greet.")
	question := flag.String("question", "How's tricks?", "The sentence to emojify.")
	block := flag.String("block", "default", "The Unicode code point block to select an emoji from.")
	flag.Parse()

	emoji := randEmoji(*block)
	fmt.Printf("Hello, %s! %c %s \n", *name, emoji, *question)
}
