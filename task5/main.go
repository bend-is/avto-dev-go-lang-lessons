package main

import (
	"context"
	"log"

	"github.com/bend-is/task5/pkg/server"
)

func main() {
	ctx := context.Background()

	go func() {
		if err := server.StartPprofServer(ctx); err != nil {
			log.Println(err)
		}
	}()

	if err := server.StartServer(ctx); err != nil {
		log.Fatal(err)
	}
}
