package redis

import (
	monitor "github.com/elojah/http-monitor"

	"github.com/go-redis/redis"
)

// CreateRequest is the implementation of Request service by mem.
func (s *Service) CreateRequest(req monitor.Request) error {
	return s.ZIncrXX(req.URL, redis.Z{Score: 1, Member: nil}).Err()
}

// ListRequest is the implementation of Request service by mem.
func (s *Service) ListRequest(monitor.RequestSubset) ([]monitor.Request, error) {
	return nil, nil
}

// CountRequest is the implementation of Request service by mem.
func (s *Service) CountRequest(monitor.RequestSubset) (int, error) {
	return 0, nil
}
