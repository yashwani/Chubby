package main

import (
	"Chubby/server"
	"flag"
	"os"
	"os/signal"
)

func main() {

	raftBind := flag.String("raftBind", "", "port bound to")
	raftDir := flag.String("raftDir", "./", "store directory")
	nodeID := flag.String("nodeID", "1", "node id")
	inmem := flag.Bool("inmem", false, "raft db stored in memory")
	listen := flag.String("listen", "", "where server listens for RPC")
	join := flag.String("join", "", "address of node in network")

	flag.Parse()

	config := server.Config{
		RaftBind: *raftBind,
		RaftDir:  *raftDir,
		NodeID:   *nodeID,
		Inmem:    *inmem,
		Listen:   *listen,
		Join:     *join,
	}

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	go server.Run(config)

	<-c
}
