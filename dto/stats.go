package dto

import (
	monitor "github.com/elojah/http-monitor"
)

// Stats represents the stats object to display regularly.
type Stats struct {
	TopHits map[string]int
}

// NewStats create a new display stats object from request hits.
func NewStats(reqs []monitor.RequestHit) Stats {
	stats := Stats{
		TopHits: make(map[string]int, len(reqs)),
	}
	for _, req := range reqs {
		stats.TopHits[req.URL] = req.Hit
	}
	return stats
}
