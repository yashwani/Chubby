package client

const rpcServerLeader = "Server.Leader"
const rpcServerKeepAlive = "Server.KeepAlive"
const rpcServerCreateSession = "Server.CreateSession"

// LeaderRequest is the RPC request for if a server is a leader.
type LeaderRequest struct {
}

// LeaderResponse is the RPC response that returns the leader's address.
type LeaderResponse struct {
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
