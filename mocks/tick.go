package mocks

import (
	"sync/atomic"

	monitor "github.com/elojah/http-monitor"
)

// TickMapper mocks Tick mapper.
type TickMapper struct {
	CreateTickFunc  func(monitor.Tick) error
	CreateTickCount int32
	CountTickFunc   func(monitor.TickSubset) (int, error)
	CountTickCount  int32
}

// CreateTick mocks tick mapper.
func (m *TickMapper) CreateTick(tick monitor.Tick) error {
	atomic.AddInt32(&m.CreateTickCount, 1)
	if m.CreateTickFunc == nil {
		return nil
	}
	return m.CreateTickFunc(tick)
}

// CountTick mocks tick mapper.
func (m *TickMapper) CountTick(subset monitor.TickSubset) (int, error) {
	atomic.AddInt32(&m.CountTickCount, 1)
	if m.CountTickFunc == nil {
		return 0, nil
	}
	return m.CountTickFunc(subset)
}
