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
	monitor.LogSectionMapper
	monitor.SectionMapper
	monitor.TickMapper

	ticker *time.Ticker

	logFile    string
	topDisplay uint
	tickTTL    time.Duration
}

// NewLogReader returns a new log reader.
func NewLogReader(mappers monitor.Mappers) *LogReader {
	return &LogReader{
		LogSectionMapper: mappers,
		SectionMapper:    mappers,
		TickMapper:       mappers,
	}
}

// Dial configure app with right settings.
func (lr *LogReader) Dial(c LogReaderConfig) error {
	var err error
	lr.logFile = c.LogFile
	lr.topDisplay = c.TopDisplay
	statsGap, err := time.ParseDuration(c.StatsGap)
	if err != nil {
		return err
	}
	lr.ticker = time.NewTicker(statsGap)
	return nil
}

// Close interrupts the ticker and log reading,
func (lr *LogReader) Close() {
	lr.ticker.Stop()
}

// Start starts the reading log process + regular display of stats.
func (lr *LogReader) Start() error {
	t, err := tail.TailFile(lr.logFile, tail.Config{Follow: true, Logger: tail.DiscardingLogger})
	if err != nil {
		return err
	}
	for {
		select {
		case _, ok := <-lr.ticker.C:
			if !ok {
				return nil
			}
			if err := lr.logSection(); err != nil {
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
			if err := lr.CreateTick(monitor.Tick{RequestID: req.ID, TS: time.Now()}); err != nil {
				return err
			}
		}
	}
}

// logSection log stats at time t.
func (lr *LogReader) logSection() error {
	reqs, err := lr.ListSection(monitor.SectionSubset{})
	if err != nil {
		return err
	}
	lr.LogSection(reqs, lr.topDisplay)
	return nil
}
