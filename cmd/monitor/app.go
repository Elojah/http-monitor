package main

import (
	"github.com/elojah/http-monitor"
)

// App is the main monitor app responsible fo reading logs and displaying stats.
type App struct {
	monitor.RequestMapper

	logFile       string
	statsInterval uint
}

// Dial configure app with right settings.
func (a *App) Dial(c Config) error {
	a.logFile = c.LogFile
	a.statsInterval = c.StatsInterval
	return nil
}

// NewApp returns a new app.
func NewApp() *App {
	return nil
}
