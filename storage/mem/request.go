package mem

import monitor "github.com/elojah/http-monitor"

// CreateRequest is the implementation of Request service by mem.
func (s *Service) CreateRequest(monitor.Request) error {
	return nil
}

// ListRequest is the implementation of Request service by mem.
func (s *Service) ListRequest(monitor.RequestSubset) ([]monitor.Request, error) {
	return nil, nil
}

// CountRequest is the implementation of Request service by mem.
func (s *Service) CountRequest(monitor.RequestSubset) (int, error) {
	return 0, nil
}
