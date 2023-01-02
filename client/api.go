package client

import (
	"log"
	"net/rpc"
	"time"
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

	// timeout is the client's approximation of the server's lease timeout
	timeout time.Time

	// jeapardy
	jeopardy bool

	// client ID
	id string
}

// Open creates Chubby client handler and starts a session with the Chubby cell.
func Open(id string) (*Handle, error) {

	handle := &Handle{
		quitSession: make(chan bool),
		id:          id,
	}

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

		client.Close()
	}

	var err error

	println(handle.Leader)

	handle.client, err = rpc.Dial("tcp", "127.0.0.1:7134") // TODO convert from raftport to normal port later
	if err != nil {
		log.Printf("Unable to dial server at %s", handle.Leader)
	}

	request := CreateSessionRequest{
		ID:      handle.id,
		Timeout: 1 * time.Second,
	}

	handle.timeout = time.Now().Add(request.Timeout)

	log.Printf("Timeout after creating session: %s", handle.timeout)

	handle.monitorSession()

	response := &CreateSessionResponse{}

	if err = handle.client.Call(rpcServerCreateSession, request, response); err != nil {
		log.Fatalf("error herer1: %s", err)
	}

	go handle.KeepAlive()

	println("here")

	return handle, nil
}

func (h *Handle) monitorSession() {

	ticker := time.NewTicker(time.Second)

	go func() {

		for {

			now := <-ticker.C

			if now.After(h.timeout) {

				if h.jeopardy {

					h.jeopardy = false

					log.Print("Jeopard timeout, quitting session.")

					h.quitSession <- true

					return
				}

				log.Print("Entering jeopardy.")

				h.jeopardy = true

				h.timeout = h.timeout.Add(45 * time.Second)

			}
		}
	}()
}

func (h *Handle) KeepAlive() {

	lastSuccess := true

	var err error

	for {

		select {

		case <-h.quitSession:

			log.Printf("terminating session")

			return

		default:

			if h.jeopardy {

				h.client, err = rpc.Dial("tcp", "127.0.0.1:7134") // TODO convert from raftport to normal port later

				if err != nil {

					log.Printf("Unable to dial server at %s", h.Leader)

					time.Sleep(time.Second)

					continue
				}

				h.jeopardy = false

				h.timeout = time.Now().Add(time.Second)

				lastSuccess = true
			}

			log.Printf("sending keep alive")
			request := KeepAliveRequest{
				ID:        h.id,
				Extension: 12 * time.Second,
				Buffer:    2 * time.Second,
			}

			if lastSuccess {
				h.timeout = h.timeout.Add(request.Extension)
			}

			log.Printf("Update timeout to %s", h.timeout)

			response := &KeepAliveResponse{}

			err := h.client.Call(rpcServerKeepAlive, request, response)
			if err != nil {

				log.Printf("keep alive failed: %s", err.Error())

				lastSuccess = false

				time.Sleep(time.Second)
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
