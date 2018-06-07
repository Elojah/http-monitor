package main

import (
	"time"

	"github.com/elojah/http-monitor"
	"github.com/elojah/http-monitor/dto"
)

// Alerter is the main monitor app responsible fo reading logs and displaying stats.
type Alerter struct {
	monitor.TickMapper

	ticker    *time.Ticker
	lastAlert time.Time

	treshold       uint
	triggerRange   time.Duration
	triggerRecover time.Duration
	reboundGap     time.Duration

	status monitor.AlertStatus
}

// NewAlerter returns a new alerter.
func NewAlerter(mappers monitor.Mappers) *Alerter {
	return &Alerter{
		TickMapper: mappers,
		status:     monitor.Down,
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
	a.triggerRecover, err = time.ParseDuration(c.TriggerRecover)
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
		status := a.status.Load()
		switch status {
		case monitor.Down:
			if ticks >= int(a.treshold) && ts.Sub(a.lastAlert) > a.reboundGap {
				a.status.Store(monitor.Up)
				a.lastAlert = ts
				alert := monitor.Alert{Ticks: ticks, TS: ts, Status: monitor.Up}
				dto.NewAlert(alert).Log()
			}
		case monitor.Up:
			if ticks < int(a.treshold) && ts.Sub(a.lastAlert) > a.triggerRecover {
				a.status.Store(monitor.Down)
				alert := monitor.Alert{Ticks: ticks, TS: ts, Status: monitor.Down}
				dto.NewAlert(alert).Log()
			}
		}
	}
	return nil
}
