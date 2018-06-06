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

	treshold     uint
	triggerRange time.Duration
	reboundGap   time.Duration
}

// NewAlerter returns a new alerter.
func NewAlerter(services monitor.Services) *Alerter {
	return &Alerter{
		TickMapper: services,
	}
}

// Dial configure app with right settings.
func (a *Alerter) Dial(c AlerterConfig) error {
	var err error
	a.treshold = c.Treshold
	a.triggerRange, err = time.ParseDuration(c.TriggerRange)
	if err != nil {
		return err
	}
	a.reboundGap, err = time.ParseDuration(c.ReboundGap)
	if err != nil {
		return err
	}
	reccurGap, err := time.ParseDuration(c.ReccurGap)
	if err != nil {
		return err
	}
	a.ticker = time.NewTicker(reccurGap)
	return nil
}

// Close interrupts the ticker and log reading,
func (a *Alerter) Close() {
	a.ticker.Stop()
}

// Start starts the alert service.
func (a *Alerter) Start() error {
	for ts := range a.ticker.C {
		min := ts.Add(-a.triggerRange)
		max := ts
		ticks, err := a.CountTick(monitor.TickSubset{Min: &min, Max: &max})
		if err != nil {
			return err
		}
		if ticks > int(a.treshold) && ts.Sub(a.lastAlert) > a.reboundGap {
			a.lastAlert = ts
			a.LogAlert(ticks, ts)
		}
	}
	return nil
}

// LogAlert log an alert of ticks at time ts.
func (a *Alerter) LogAlert(ticks int, ts time.Time) {
	log.Infof("High traffic generated an alert - hits = %d, triggered at %s", ticks, ts.String())
}
