package main

import (
	"sync/atomic"
	"testing"
	"time"

	monitor "github.com/elojah/http-monitor"
	"github.com/elojah/http-monitor/mocks"
)

func TestStart(t *testing.T) {
	t.Run("0_alert_up", func(t *testing.T) {
		var mappers monitor.Mappers
		tm := &mocks.TickMapper{}
		tm.CountTickFunc = func(subset monitor.TickSubset) (int, error) {
			return 10, nil
		}
		am := &mocks.LogAlertMapper{}
		mappers.TickMapper = tm
		mappers.LogAlertMapper = am

		config := AlerterConfig{
			Treshold:       11,
			TriggerRange:   "0s",
			ReccurGap:      "100ms",
			TriggerRecover: "10s",
			ReboundGap:     "10s",
		}
		alerter := NewAlerter(mappers)
		if err := alerter.Dial(config); err != nil {
			t.Error(err)
			return
		}
		go func() { alerter.Start() }()
		time.Sleep(300 * time.Millisecond)
		alerter.Close()
		if am.LogAlertUpCount != 0 {
			t.Errorf(`expected=0, actual=%d`, am.LogAlertUpCount)
		}
		if am.LogAlertDownCount != 0 {
			t.Errorf(`expected=0, actual=%d`, am.LogAlertDownCount)
		}
	})

	t.Run("1_alert_up", func(t *testing.T) {
		var mappers monitor.Mappers
		tm := &mocks.TickMapper{}
		tm.CountTickFunc = func(subset monitor.TickSubset) (int, error) {
			return 10, nil
		}
		am := &mocks.LogAlertMapper{}
		mappers.TickMapper = tm
		mappers.LogAlertMapper = am

		config := AlerterConfig{
			Treshold:       10,
			TriggerRange:   "0s",
			ReccurGap:      "100ms",
			TriggerRecover: "10s",
			ReboundGap:     "10s",
		}
		alerter := NewAlerter(mappers)
		if err := alerter.Dial(config); err != nil {
			t.Error(err)
			return
		}
		go func() { alerter.Start() }()
		time.Sleep(300 * time.Millisecond)
		alerter.Close()
		logAlertUpCount := atomic.LoadInt32(&am.LogAlertUpCount)
		logAlertDownCount := atomic.LoadInt32(&am.LogAlertDownCount)
		if logAlertUpCount != 1 {
			t.Errorf(`expected=1, actual=%d`, logAlertUpCount)
		}
		if logAlertDownCount != 0 {
			t.Errorf(`expected=0, actual=%d`, logAlertDownCount)
		}
	})

	t.Run("1_alert_up_down", func(t *testing.T) {
		var mappers monitor.Mappers
		tm := &mocks.TickMapper{}
		tm.CountTickFunc = func(subset monitor.TickSubset) (int, error) {
			switch atomic.LoadInt32(&tm.CountTickCount) {
			case 1:
				return 11, nil
			default:
				return 0, nil
			}
		}
		am := &mocks.LogAlertMapper{}
		mappers.TickMapper = tm
		mappers.LogAlertMapper = am

		config := AlerterConfig{
			Treshold:       10,
			TriggerRange:   "0s",
			ReccurGap:      "100ms",
			TriggerRecover: "200ms",
			ReboundGap:     "10s",
		}
		alerter := NewAlerter(mappers)
		if err := alerter.Dial(config); err != nil {
			t.Error(err)
			return
		}
		go func() { alerter.Start() }()
		time.Sleep(500 * time.Millisecond)
		alerter.Close()
		logAlertUpCount := atomic.LoadInt32(&am.LogAlertUpCount)
		logAlertDownCount := atomic.LoadInt32(&am.LogAlertDownCount)
		if logAlertUpCount != 1 {
			t.Errorf(`expected=1, actual=%d`, logAlertUpCount)
		}
		if logAlertDownCount != 1 {
			t.Errorf(`expected=1, actual=%d`, logAlertDownCount)
		}
	})

	t.Run("2_alert_up_bound", func(t *testing.T) {
		var mappers monitor.Mappers
		tm := &mocks.TickMapper{}
		tm.CountTickFunc = func(subset monitor.TickSubset) (int, error) {
			return 11, nil
		}
		am := &mocks.LogAlertMapper{}
		mappers.TickMapper = tm
		mappers.LogAlertMapper = am

		config := AlerterConfig{
			Treshold:       10,
			TriggerRange:   "0ms",
			ReccurGap:      "100ms",
			TriggerRecover: "200ms",
			ReboundGap:     "300ms",
		}
		alerter := NewAlerter(mappers)
		if err := alerter.Dial(config); err != nil {
			t.Error(err)
			return
		}
		go func() { alerter.Start() }()
		time.Sleep(1 * time.Second)
		alerter.Close()
		logAlertUpCount := atomic.LoadInt32(&am.LogAlertUpCount)
		logAlertDownCount := atomic.LoadInt32(&am.LogAlertDownCount)
		if logAlertUpCount != 3 {
			t.Errorf(`expected=3, actual=%d`, logAlertUpCount)
		}
		if logAlertDownCount != 0 {
			t.Errorf(`expected=0, actual=%d`, logAlertDownCount)
		}
	})
}
