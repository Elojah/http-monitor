package dto

import (
	"fmt"

	monitor "github.com/elojah/http-monitor"
)

// Stats represents the stats object to display regularly.
type Stats struct {
	TopHits map[string]int
}

// NewStats create a new display stats object from section hits.
func NewStats(shs []monitor.SectionHit) Stats {
	stats := Stats{
		TopHits: make(map[string]int, len(shs)),
	}
	for _, sh := range shs {
		stats.TopHits[sh.Section] = sh.Hit
	}
	return stats
}

// String returns the string representation sent for logs.
func (s Stats) String() string {
	var str string
	for key, value := range s.TopHits {
		str += fmt.Sprintf("%s: %d\n", key, value)
	}
	return str
}
