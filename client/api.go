package client

// Handle is a Chubby client handler.
type Handle struct {

	// Lock is the node or file name
	Lock string
}

// Open creates Chubby client handler and starts a session with the Chubby cell.
func Open(lock string) (*Handle, error) {
	// send RPC call to create session with the server
	return nil, nil
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
