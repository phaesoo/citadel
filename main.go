package main

import (
	"log"
)

type App interface {
	Listen() error
	Shutdown() error
}

func main() {
	log.Print("Call main")
	log.Print("Finished")
}
