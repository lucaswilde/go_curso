package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// cards := newDeck()
	// cards.shuffe()
	// cards.print()
}

func getEvents(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}
