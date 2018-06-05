package redis

import (
	monitor "github.com/elojah/http-monitor"

	"github.com/go-redis/redis"
)

const (
	requestHitKey = "request_hit:"
)

// AddRequestHit is the implementation of Request service by redis.
func (s *Service) AddRequestHit(req monitor.Request) error {
	return s.Client.ZIncr(requestHitKey, redis.Z{Score: 1, Member: req.URL}).Err()
}

// ListRequestHit is the implementation of Request service by redis.
func (s *Service) ListRequestHit(subset monitor.RequestSubset) ([]monitor.RequestHit, error) {
	if subset.TopHits != nil {
		cmd := s.ZRevRangeByScoreWithScores(
			requestHitKey,
			redis.ZRangeBy{
				Count: int64(*subset.TopHits),
				Min:   "-inf",
				Max:   "+inf",
			})
		vals, err := cmd.Result()
		if err != nil {
			return nil, err
		}
		reqs := make([]monitor.RequestHit, len(vals))
		for i, val := range vals {
			reqs[i] = monitor.RequestHit{
				Request: monitor.Request{URL: val.Member.(string)},
				Hit:     int(val.Score),
			}
		}
		return reqs, nil
	}
	return nil, nil
}

// ResetRequestHit reset all registered request hits.
func (s *Service) ResetRequestHit() error {
	return s.Del(requestHitKey).Err()
}
