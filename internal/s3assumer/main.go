package main

import (
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }
