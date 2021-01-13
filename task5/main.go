package main

import (
	"context"
	"log"

	"github.com/bend-is/task5/pkg/server"
)

func main() {
	pprofCtx, pprofCancel := context.WithCancel(context.Background())

	go func() {
		defer pprofCancel()
		if err := server.StartServer(context.Background()); err != nil {
			log.Println(err)
		}
	}()

	if err := server.StartPprofServer(pprofCtx); err != nil {
		log.Println(err)
	}
}
