package main

import (
	"Chubby/client"
	"log"
)

func main() {

	chubby, err := client.Open()
	if err != nil {
		log.Fatal(err)
	}

	chubby.Close()
}
