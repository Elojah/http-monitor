package redis

import (
	monitor "github.com/elojah/http-monitor"

	"github.com/go-redis/redis"
)

const (
	sectionHitKey = "section_hit:"
)

// AddSectionHit is the implementation of Request service by redis.
func (s *Service) AddSectionHit(section string) error {
	return s.Client.ZIncr(sectionHitKey, redis.Z{Score: 1, Member: section}).Err()
}

// ListSectionHit is the implementation of Request service by redis.
func (s *Service) ListSectionHit(subset monitor.SectionHitSubset) ([]monitor.SectionHit, error) {
	if subset.TopHits != nil {
		cmd := s.ZRevRangeByScoreWithScores(
			sectionHitKey,
			redis.ZRangeBy{
				Count: int64(*subset.TopHits),
				Min:   "-inf",
				Max:   "+inf",
			})
		vals, err := cmd.Result()
		if err != nil {
			return nil, err
		}
		reqs := make([]monitor.SectionHit, len(vals))
		for i, val := range vals {
			reqs[i] = monitor.SectionHit{
				Section: val.Member.(string),
				Hit:     int(val.Score),
			}
		}
		return reqs, nil
	}
	return nil, nil
}

// ResetSectionHit reset all registered request hits.
func (s *Service) ResetSectionHit() error {
	return s.Del(sectionHitKey).Err()
}
