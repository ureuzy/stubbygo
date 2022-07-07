package main

import (
	"log"

	"github.com/masanetes/stubbygo/pkg/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
