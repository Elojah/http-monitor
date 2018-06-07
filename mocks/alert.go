package mocks

import (
	"sync/atomic"

	monitor "github.com/elojah/http-monitor"
)

// AlertMapper mocks Alert mapper.
type AlertMapper struct {
	LogAlertFunc      func(monitor.Alert)
	LogAlertUpCount   int32
	LogAlertDownCount int32
}

// LogAlert mocks alert mapper.
func (m *AlertMapper) LogAlert(alert monitor.Alert) {
	switch alert.Status {
	case monitor.Down:
		atomic.AddInt32(&m.LogAlertDownCount, 1)
	case monitor.Up:
		atomic.AddInt32(&m.LogAlertUpCount, 1)
	}
	if m.LogAlertFunc == nil {
		return
	}
	m.LogAlertFunc(alert)
}
