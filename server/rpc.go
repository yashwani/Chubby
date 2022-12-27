package server

// LeaderRequest is the RPC request for if a server is a leader.
type LeaderRequest struct {
}

// LeaderResponse is the RPC response that returns the leader's address.
type LeaderResponse struct {
	Address string
}

// Leader returns the address of the leader of the cell.
func (s *Server) Leader(req LeaderRequest, resp *LeaderResponse) error {

	address, _ := s.store.Raft.LeaderWithID()

	resp.Address = string(address)

	return nil
}
