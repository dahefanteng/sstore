package main

import (
	"github.com/dahefanteng/sstore/mon"
	"log"
)

func main() {
	log.Println("start the mon server..")
	mon.StartServer()
}
