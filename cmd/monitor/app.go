package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/elojah/http-monitor"
	"github.com/elojah/http-monitor/dto"
	"github.com/hpcloud/tail"
)

// App is the main monitor app responsible fo reading logs and displaying stats.
type App struct {
	monitor.RequestHitMapper

	ticker *time.Ticker

	logFile    string
	topDisplay uint
}

// NewApp returns a new app.
func NewApp(rm monitor.RequestHitMapper) *App {
	return &App{
		RequestHitMapper: rm,
	}
}

// Dial configure app with right settings.
func (a *App) Dial(c Config) error {
	a.logFile = c.LogFile
	a.topDisplay = c.TopDisplay
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
		case _, ok := <-a.ticker.C:
			if !ok {
				return nil
			}
			if err := a.LogStats(); err != nil {
				return err
			}
			if err := a.ResetRequestHit(); err != nil {
				return err
			}
		case line := <-t.Lines:
			if line.Err != nil {
				return err
			}
			clf, err := dto.NewCLF(line.Text)
			if err != nil {
				// Don't stop if a line has wrong format
				log.Error(err)
				continue
			}
			req, err := clf.NewRequest()
			if err != nil {
				// Don't stop if a line has wrong format
				log.Error(err)
				continue
			}
			if err := a.AddRequestHit(req); err != nil {
				return err
			}
		}
	}
}

// LogStats log stats at time t.
func (a *App) LogStats() error {
	reqs, err := a.ListRequestHit(monitor.RequestSubset{TopHits: &a.topDisplay})
	if err != nil {
		return err
	}
	log.Info(dto.NewStats(reqs).String())
	return nil
}
