package main

import (
	"flag"
	"log"

	"github.com/ureuzy/stubbygo/pkg/server"
)

var (
	c = flag.String("c", "", "config file")
)

func main() {
	flag.Parse()

	if *c == "" {
		log.Fatal("must be specified config.yaml")
	}

	if err := server.Run(*c); err != nil {
		log.Fatal(err)
	}
}
