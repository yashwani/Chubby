package server

import (
	"errors"
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

	// s.extend[req.ID] = make(chan time.Duration)

	s.timeouts[req.ID] = time.Now().Add(req.Timeout)

	// go func() {
	// 	s.extend[req.ID] <- req.Timeout
	// }()

	// time.Sleep(1 * time.Second)

	// go s.MaintainSession(req.ID)

	return nil
}

// func (s *Server) MaintainSession(ID string) {

// 	for {

// 		select {
// 		case interval := <-s.extend[ID]:
// 			log.Printf("Sleeping for %v seconds", int(interval))
// 			time.Sleep(interval)
// 		default:
// 			log.Printf("Terminating session for client %v", ID)
// 			s.sessions[ID] = false
// 			return
// 		}
// 	}
// }

// KeepAlive extends the client's session by a pre-defined interval.
func (s *Server) KeepAlive(req KeepAliveRequest, resp *KeepAliveResponse) error {

	if !s.sessions[req.ID] {
		return errors.New("keep alive failed because current session not created or terminated")
	}

	s.timeouts[req.ID] = s.timeouts[req.ID].Add(req.Extension)

	log.Printf("Update %s timeout to %s", req.ID, s.timeouts[req.ID])

	time.Sleep(req.Extension - req.Buffer)

	log.Printf("Return keep alive to client %s", req.ID)

	return nil
}
