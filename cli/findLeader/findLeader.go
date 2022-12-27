package main

import (
	"Chubby/client"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Kill, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	chubby, err := client.Open()
	if err != nil {
		log.Fatal(err)
	}

	chubby.Close()

	for {
		// Exit on signal.
		select {
		case <-quitCh:
			break
		default:
			continue
		}
	}
}
