package client

import (
	"log"
	"net/rpc"
)

// cell is the network of servers in the chubby cell.
var cell = map[string]bool{
	"localhost:7134": true,
	"localhost:7234": true,
	"localhost:7334": true,
}

// Handle is a Chubby client handler.
type Handle struct {

	// Lock is the node or file name
	Lock string

	// Leader is the address of the cell's leader
	Leader string

	// client is the RPC client
	client *rpc.Client

	// quitSession is added to when the client wants to terminate its session with the chubby cell
	quitSession chan bool
}

// Open creates Chubby client handler and starts a session with the Chubby cell.
func Open() (*Handle, error) {

	handle := &Handle{}

	for server := range cell {

		log.Print(server)

		client, err := rpc.Dial("tcp", server)
		if err != nil {
			log.Printf("Unable to dial server at %s", server)
		}

		request := LeaderRequest{}
		response := &LeaderResponse{}

		if err := client.Call(rpcServerLeader, request, response); err != nil {

			log.Printf("Failed to make %s RPC call to server %s: %s", rpcServerLeader, server, err)

			continue
		}

		if response.Address != "" {

			handle.Leader = response.Address

			break
		}
	}

	// ping leader

	var err error

	handle.client, err = rpc.Dial("tcp", "localhost:7134")
	if err != nil {
		log.Printf("Unable to dial server at %s", handle.Leader)
	}

	request := CreateSessionRequest{
		ID: "123",
	}
	response := &CreateSessionResponse{}

	if err = handle.client.Call(rpcServerCreateSession, request, response); err != nil {
		log.Fatalf("error herer1: %s", err)
	}

	log.Print(6)

	done := make(chan bool)

	go handle.KeepAlive(done)

	//       if this succeeds, then spawn keepAlive goroutine

	//       return successfully

	return handle, nil
}

func (h *Handle) KeepAlive(done chan bool) {

	for {
		select {

		case <-h.quitSession:

			log.Printf("terminating session")

			done <- true

			return

		default:
			log.Printf("sending keep alive")
			request := KeepAliveRequest{
				ID: "123",
			}

			response := &KeepAliveResponse{}

			if err := h.client.Call(rpcServerKeepAlive, request, response); err != nil {
				log.Fatalf(err.Error())
			}
		}
	}
}

// Close tears down the session with the Chubby cell thereby destroying the handle.
func (h *Handle) Close() error {
	//
	return nil
}

// GetContentsAndStat reads the contents of the lock.
func (h *Handle) GetContentsAndStat() error {
	return nil
}

// SetContents writes contents to the lock.
func (h *Handle) SetContents() error {
	return nil
}

// Delete deletes the lock.
func (h *Handle) Delete() error {
	return nil
}

// Acquire tries to acquire the lock.
func (h *Handle) Acquire() error {
	return nil
}

// Release releases the lock.
func (h *Handle) Release() error {
	return nil
}
