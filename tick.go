package monitor

import "time"

// Tick represents an alert for high traffic during a period of time.
type Tick struct {
	TS  time.Time
	TTL time.Duration
}

// TickMapper is a data interface for Tick object.
type TickMapper interface {
	IncrTick(Tick) error
	CountTick() (int64, error)
}
