package server

// JoinRequest is the RPC request to join a Chubby cell.
type JoinRequest struct {

	// NodeID is the ID of the server wishing to join the Chubby cell
	NodeID string

	// addr is the address of the server wishing to join the Chubby cell
	Addr string
}

// JoinResponse is the RPC response to joining a Chubby cell.
type JoinResponse struct {
	Error error
}

// LeaderRequest is the RPC request for if a server is a leader.
type LeaderRequest struct {
}

// LeaderResponse is the RPC response that returns the leader's address.
type LeaderResponse struct {

	// Address of the leader
	Address string
}

// KeepAliveRequest is the RPC request to keep the session alive.
type KeepAliveRequest struct {

	// ID of the client sending the request
	ID string
}

// KeepAliveResponse is the RPC response for the client's keepAlive request.
type KeepAliveResponse struct {
	Something bool
}

type CreateSessionRequest struct {
	ID string
}

type CreateSessionResponse struct {
}
