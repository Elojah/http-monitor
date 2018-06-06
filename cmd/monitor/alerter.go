package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/elojah/http-monitor"
)

// Alerter is the main monitor app responsible fo reading logs and displaying stats.
type Alerter struct {
	monitor.TickMapper

	ticker    *time.Ticker
	lastAlert time.Time

	reqPerSec   uint
	triggerTime time.Duration
	reboundTime time.Duration
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
	a.triggerTime = time.Duration(c.AlertTriggerTime) * time.Second
	a.reboundTime = time.Duration(c.AlertReboundTime) * time.Second
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
		min := ts.Add(-a.triggerTime)
		max := ts
		ticks, err := a.CountTick(monitor.TickSubset{Min: &min, Max: &max})
		if err != nil {
			return err
		}
		if ticks > int(a.reqPerSec) && ts.Sub(a.lastAlert) > a.reboundTime {
			a.lastAlert = ts
			a.LogAlert(ticks, ts)
		}
	}
	return nil
}

// LogAlert log an alert of ticks at time ts.
func (a *Alerter) LogAlert(ticks int, ts time.Time) {
	log.WithFields(log.Fields{
		"type":  "alert",
		"ticks": ticks,
		"ts":    ts.Unix(),
	}).Info("alert triggered")
}
