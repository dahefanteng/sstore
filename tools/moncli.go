package main

import (
	"context"
	"github.com/dahefanteng/sstore/mon"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	conn, err := grpc.Dial("localhost:2379", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	c := mon.NewMonClient(conn)
	req := &mon.SetMapRequest{"aa", "cc"}
	r, err := c.SetMap(context.Background(), req)
	log.Println(r, err)
}
