package redis

import (
	monitor "github.com/elojah/http-monitor"

	"github.com/go-redis/redis"
)

const (
	sectionKey = "section:"
)

// IncrSection is the implementation of Section service by redis.
func (s *Service) IncrSection(name string) error {
	return s.Client.ZIncr(sectionKey, redis.Z{Score: 1, Member: name}).Err()
}

// ListSection is the implementation of Section service by redis.
func (s *Service) ListSection(subset monitor.SectionSubset) ([]monitor.Section, error) {
	var count int64
	if subset.TopHits != nil {
		count = int64(*subset.TopHits)
	}
	cmd := s.ZRevRangeByScoreWithScores(
		sectionKey,
		redis.ZRangeBy{
			Count: count,
			Min:   "-inf",
			Max:   "+inf",
		})
	vals, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	reqs := make([]monitor.Section, len(vals))
	for i, val := range vals {
		reqs[i] = monitor.Section{
			Name: val.Member.(string),
			Hit:  int(val.Score),
		}
	}
	return reqs, nil
}

// ResetSection reset all registered request hits.
func (s *Service) ResetSection() error {
	return s.Del(sectionKey).Err()
}
