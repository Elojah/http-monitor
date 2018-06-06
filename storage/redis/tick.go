package redis

import (
	"strconv"

	"github.com/go-redis/redis"

	monitor "github.com/elojah/http-monitor"
)

const (
	tickKey = "tick:"
)

// CreateTick is the implementation of Tick service by redis.
func (s *Service) CreateTick(tick monitor.Tick) error {
	return s.ZAdd(tickKey, redis.Z{Score: float64(tick.TS.Unix()), Member: tick.RequestID.String()}).Err()
}

// CountTick is the implementation of Tick service by redis.
func (s *Service) CountTick(subset monitor.TickSubset) (int, error) {
	min := "-inf"
	max := "+inf"
	if subset.Min != nil {
		min = strconv.FormatInt(subset.Min.Unix(), 10)
	}
	if subset.Max != nil {
		max = strconv.FormatInt(subset.Max.Unix(), 10)
	}
	cmd := s.ZRangeByScore(tickKey, redis.ZRangeBy{
		Min: min,
		Max: max,
	})
	vals, err := cmd.Result()
	if err != nil {
		return 0, err
	}
	return len(vals), nil
}
