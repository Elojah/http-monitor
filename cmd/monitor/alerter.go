package main

import (
	"time"

	"github.com/elojah/http-monitor"
)

// Alerter is the main monitor app responsible fo reading logs and displaying stats.
type Alerter struct {
	monitor.TickMapper

	ticker *time.Ticker

	reqPerSec   uint
	triggerTime uint
}

// NewAlerter returns a new alerter.
func NewAlerter(services monitor.Services) *Alerter {
	return &Alerter{
		TickMapper: services,
	}
}

// Dial configure app with right settings.
func (a *Alerter) Dial(c Config) error {
	a.reqPerSec = c.AlertReqPerSec
	a.triggerTime = c.AlertTriggerTime
	a.ticker = time.NewTicker(time.Second * time.Duration(c.AlertReccurTime))
	return nil
}

// Close interrupts the ticker and log reading,
func (a *Alerter) Close() {
	a.ticker.Stop()
}

// Start starts the alert service.
func (a *Alerter) Start() error {
	for ts := range a.ticker.C {
		ticks, err := a.CountTick()
		if err != nil {
			return err
		}
		if ticks > int64(a.reqPerSec) {
			a.LogAlert(ticks, ts)
		}
	}
	return nil
}

// LogAlert log an alert of ticks at time ts.
func (a *Alerter) LogAlert(ticks int64, ts time.Time) {

}
