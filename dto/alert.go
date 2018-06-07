package dto

import (
	"fmt"

	monitor "github.com/elojah/http-monitor"
)

// LogAlert logs the alert on stdout.
func (Service) LogAlert(a monitor.Alert) {
	switch a.Status {
	case monitor.Down:
		fmt.Printf("Alert recovered - hits = %d, triggered at %s", a.Ticks, a.TS.String())
	case monitor.Up:
		fmt.Printf("High traffic generated an alert - hits = %d, triggered at %s", a.Ticks, a.TS.String())
	}
}
