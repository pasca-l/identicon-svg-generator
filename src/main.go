package main

import (
	"log"

	"github.com/pasca-l/identicon-generator/identicon"
)

func main() {
	err := identicon.GenerateIdenticon("pasca-l")
	if err != nil {
		log.Fatal(err)
	}
}
