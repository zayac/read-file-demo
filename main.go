package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Event struct {
	Date        string `json:"date"`
	Description string `json:"description"`
}

func ReadEvents(fname string) []Event {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var events []Event
	if json.NewDecoder(bufio.NewReader(f)).Decode(&events); err != nil {
		log.Fatal(err)
	}
	return events
}

func main() {
	fmt.Println(ReadEvents("input.json"))
}
