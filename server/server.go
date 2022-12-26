package server

import (
	"Chubby/store"
	"log"
	"net"
	"net/rpc"
	"time"
)

type Config struct {
	RaftBind string
	RaftDir  string
	NodeID   string
	Inmem    bool
	Listen   string
	Join     string
}

type Server struct {
	store *store.Store
}

func Run(config Config) {

	server := &Server{
		store: store.New(config.Inmem, config.RaftDir, config.RaftBind),
	}

	if err := server.store.Open(config.Join == "", config.NodeID); err != nil {
		log.Fatalf("Failed to open store: %s", err)
	}

	if config.Join != "" {

		client, err := rpc.Dial("tcp", config.Join)
		if err != nil {
			log.Fatal(err)
		}

		request := JoinRequest{
			NodeID: config.NodeID,
			Addr:   config.RaftBind,
		}

		response := &JoinResponse{}

		if err := client.Call("Server.Join", request, response); err != nil {
			log.Fatalf("Failed to make RPC call to join Chubby cell: %s", err)
		}

		if response.Error != nil {
			log.Fatalf("Failed to join existing Chubby cell: %s", err)
		}
	}

	rpc.Register(server)

	listener, err := net.Listen("tcp", config.Listen)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	println(server.store.Raft.LeaderWithID())

	rpc.Accept(listener)

}
