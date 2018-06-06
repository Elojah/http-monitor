package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/elojah/http-monitor"
	"github.com/elojah/http-monitor/dto"
	"github.com/hpcloud/tail"
)

// LogReader is the main monitor app responsible fo reading logs and displaying stats.
type LogReader struct {
	monitor.SectionMapper
	monitor.TickMapper

	ticker *time.Ticker

	logFile    string
	topDisplay uint
	tickTTL    time.Duration
}

// NewLogReader returns a new log reader.
func NewLogReader(services monitor.Services) *LogReader {
	return &LogReader{
		SectionMapper: services,
		TickMapper:    services,
	}
}

// Dial configure app with right settings.
func (lr *LogReader) Dial(c Config) error {
	lr.logFile = c.LogFile
	lr.topDisplay = c.TopDisplay
	lr.ticker = time.NewTicker(time.Second * time.Duration(c.StatsInterval))
	return nil
}

// Close interrupts the ticker and log reading,
func (lr *LogReader) Close() {
	lr.ticker.Stop()
}

// Start starts the reading log process + regular display of stats.
func (lr *LogReader) Start() error {
	t, err := tail.TailFile(lr.logFile, tail.Config{Follow: true})
	if err != nil {
		return err
	}
	for {
		select {
		case _, ok := <-lr.ticker.C:
			if !ok {
				return nil
			}
			if err := lr.LogStats(); err != nil {
				return err
			}
			if err := lr.ResetSection(); err != nil {
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
			if err := lr.IncrSection(req.Section()); err != nil {
				return err
			}
			if err := lr.CreateTick(monitor.Tick{TS: req.TS}); err != nil {
				return err
			}
		}
	}
}

// LogStats log stats at time t.
func (lr *LogReader) LogStats() error {
	reqs, err := lr.ListSection(monitor.SectionSubset{TopHits: &lr.topDisplay})
	if err != nil {
		return err
	}
	log.Info(dto.NewStats(reqs).String())
	return nil
}
