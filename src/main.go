package main

import (
	"log"

	"github.com/pasca-l/identicon-svg-generator/server"
)

func main() {
	err := server.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
