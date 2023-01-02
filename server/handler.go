package server

import (
	"log"
	"time"
)

// Join adds a requesting server to the existing chubby cell.
func (s *Server) Join(req JoinRequest, resp *JoinResponse) error {

	resp.Error = s.store.Join(req.NodeID, req.Addr)

	return resp.Error
}

// Leader returns the address of the leader of the cell.
func (s *Server) Leader(req LeaderRequest, resp *LeaderResponse) error {

	address, _ := s.store.Raft.LeaderWithID()

	resp.Address = string(address)

	return nil
}

func (s *Server) CreateSession(req CreateSessionRequest, resp *CreateSessionResponse) error {

	s.sessions[req.ID] = true

	s.timeouts[req.ID] = time.Now().Add(req.Timeout)

	log.Printf("Timeout after creating session: %s", s.timeouts[req.ID])

	return nil
}

// KeepAlive extends the client's session by a pre-defined interval.
func (s *Server) KeepAlive(req KeepAliveRequest, resp *KeepAliveResponse) error {

	log.Printf("Received message from %v", req.ID)

	if !s.sessions[req.ID] {

		s.sessions[req.ID] = true

		s.timeouts[req.ID] = time.Now()
	}

	s.timeouts[req.ID] = s.timeouts[req.ID].Add(req.Extension)

	log.Printf("Update %s timeout to %s", req.ID, s.timeouts[req.ID])

	time.Sleep(req.Extension)

	log.Printf("Return keep alive to client %s", req.ID)

	return nil
}
