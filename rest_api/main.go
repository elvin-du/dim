package main

import (
	"log"
	"os"
)

func main() {
	defer func() {
		log.Println("hi")
	}()
	os.Exit(0)
}
