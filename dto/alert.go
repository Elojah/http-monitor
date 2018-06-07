package dto

import (
	"fmt"

	monitor "github.com/elojah/http-monitor"
)

// AlertService implements domain alert mapper.
type AlertService struct{}

// LogAlert logs the alert on stdout.
func (AlertService) LogAlert(a monitor.Alert) {
	switch a.Status {
	case monitor.Down:
		fmt.Printf("Alert recovered - hits = %d, triggered at %s", a.Ticks, a.TS.String())
	case monitor.Up:
		fmt.Printf("High traffic generated an alert - hits = %d, triggered at %s", a.Ticks, a.TS.String())
	}
}
