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

	//       if this succeeds, then spawn keepAlive goroutine

	//       return successfully

	return handle, nil
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
