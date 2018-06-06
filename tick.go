package monitor

import "time"

// Tick represents an alert for high traffic during a period of time.
type Tick struct {
	TS time.Time
}

// TickMapper is a data interface for Tick object.
type TickMapper interface {
	CreateTick(Tick) error
	CountTick(TickSubset) (int, error)
}

// TickSubset retreives a set of ticks indexed on their ts.
type TickSubset struct {
	Min *time.Time
	Max *time.Time
}
