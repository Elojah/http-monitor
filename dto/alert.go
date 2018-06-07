package dto

import (
	"fmt"

	monitor "github.com/elojah/http-monitor"
)

// Alert aliases a domain alert
type Alert monitor.Alert

// NewAlert returns a dto alert from a domain alert.
func NewAlert(alert monitor.Alert) Alert {
	return Alert(alert)
}

// Log logs the alert on stdout.
func (a Alert) Log() {
	switch a.Status {
	case monitor.Down:
		fmt.Printf("Alert recovered - hits = %d, triggered at %s", a.Ticks, a.TS.String())
	case monitor.Up:
		fmt.Printf("High traffic generated an alert - hits = %d, triggered at %s", a.Ticks, a.TS.String())
	}
}
