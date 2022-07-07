package main

import (
	"log"

	"github.com/masanetes/gostub/pkg/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
