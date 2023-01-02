package main

import (
	"Chubby/client"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Kill, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	id := flag.String("id", "", "ID of the client")

	flag.Parse()

	chubby, err := client.Open(*id)
	if err != nil {
		log.Fatal(err)
	}

	println("finished open")

	chubby.Close()

	println("finished close")

	<-quitCh
}
