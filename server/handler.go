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

func (s *Server) Join(req JoinRequest, resp *JoinResponse) error {

	resp.Error = s.store.Join(req.NodeID, req.Addr)

	return resp.Error
}

type OpenRequest struct {
	ClientID string
}

type OpenResponse struct {
}

func (s *Server) Open(req OpenRequest, resp *OpenResponse) error {

	return nil
}
