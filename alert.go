package monitor

import (
	"sync/atomic"
	"time"
)

// AlertStatus represents the status of an alert.
type AlertStatus int32

const (
	// Down means the alert is not raised or was raised and is not anymore.
	Down AlertStatus = 0
	// Up means the alert treshold triggered.
	Up AlertStatus = 1
)

// Load load an alert status atomically.
func (a AlertStatus) Load() AlertStatus {
	return AlertStatus(atomic.LoadInt32((*int32)(&a)))
}

// Store stores an alert status atomically.
func (a *AlertStatus) Store(status AlertStatus) {
	atomic.StoreInt32((*int32)(a), int32(status))
}

// Alert represents a raised or recovered alert.
type Alert struct {
	Ticks  int
	TS     time.Time
	Status AlertStatus
}

// AlertMapper maps alert operations.
type AlertMapper interface {
	LogAlert(Alert)
}
