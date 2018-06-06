package redis

import (
	"strconv"

	monitor "github.com/elojah/http-monitor"
)

const (
	tickKey = "tick:"
)

// CreateTick is the implementation of Tick service by redis.
func (s *Service) CreateTick(tick monitor.Tick) error {

	key := tickKey + strconv.FormatInt(tick.TS.Unix(), 10)

	// Create new tick
	if err := s.Incr(key).Err(); err != nil {
		return err
	}

	// Set expiration time
	return s.Expire(key, tick.TTL).Err()
}

// CountTick is the implementation of Tick service by redis.
func (s *Service) CountTick() (int64, error) {
	keys, err := s.Keys(tickKey + "*").Result()
	if err != nil {
		return 0, err
	}
	var count int64
	for _, key := range keys {
		val, err := s.Get(key).Result()
		if err != nil {
			return 0, err
		}
		n, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, err
		}
		count += n
	}
	return count, nil
}
