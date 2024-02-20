package main

import (
	"flag"
	"fmt"
)

func main() {
	numberOfCards := flag.Int("numberOfCards", 32, "Number of Cards to play witth")
	flag.Parse()
	fmt.Printf("Playing with %d Cards\n", *numberOfCards)
}
