package main

import (
	"time"

	"github.com/elojah/http-monitor"
	"github.com/hpcloud/tail"
)

// App is the main monitor app responsible fo reading logs and displaying stats.
type App struct {
	monitor.RequestMapper

	ticker *time.Ticker

	logFile string
}

// NewApp returns a new app.
func NewApp(rm monitor.RequestMapper) *App {
	return &App{
		RequestMapper: rm,
	}
}

// Dial configure app with right settings.
func (a *App) Dial(c Config) error {
	a.logFile = c.LogFile
	a.ticker = time.NewTicker(time.Second * time.Duration(c.StatsInterval))
	return nil
}

// Close interrupts the ticker and log reading,
func (a *App) Close() {
	a.ticker.Stop()
}

// Start starts the reading log process + regular display of stats.
func (a *App) Start() error {
	t, err := tail.TailFile(a.logFile, tail.Config{Follow: true})
	if err != nil {
		return err
	}
	for {
		select {
		case t, ok := <-a.ticker.C:
			if !ok {
				return nil
			}
			a.LogStats(t)
		case line := <-t.Lines:
			_ = line
		}
	}
}

// LogStats log stats at time t.
func (a *App) LogStats(t time.Time) {
}
