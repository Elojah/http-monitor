package main

import (
	"time"

	"github.com/hpcloud/tail"

	"github.com/elojah/http-monitor"
)

// App is the main monitor app responsible fo reading logs and displaying stats.
type App struct {
	monitor.RequestMapper

	ticker *time.Ticker
	done   chan struct{}

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
	a.ticker = time.NewTicker(time.Second * a.statsInterval)
	return nil
}

func (a *App) Close() {
	a.ticker.Stop()
	done <- struct{}{}
}

// Start starts the reading log process + regular display of stats.
func (a *App) Start() error {
	t, err := tail.TailFile(a.logFile, tail.Config{Follow: true})
	if err != nil {
		return err
	}
	for {
		select {
		case <-done:
			return
		case t := <-a.ticker.C:
			a.LogStats(t)
		case t.Lines:
		}
	}
}

// LogStats log stats at time t.
func (a *App) LogStats(t time.Time) {
}
