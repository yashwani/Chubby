package main

import (
	"Chubby/store"
	"log"
	"time"
)

func main() {

	store := store.New(false, ".cell/server1", "localhost:5379")

	if err := store.Open(true, "1"); err != nil {
		log.Fatalf("Unable to open store: %s", err)
	}

	// time for leader election
	time.Sleep(3 * time.Second)

	if err := store.Set("testKey", "testValue"); err != nil {
		log.Fatalf("Unable to set key: %s", err)
	}

	value, err := store.Get("testKey")
	if err != nil {
		log.Fatalf("Unable to get key: %s", err)
	}

	log.Printf("Value is %s", value)

	if err := store.Delete("testKey"); err != nil {
		log.Fatalf("Unable to delete key: %s", err)
	}
}
