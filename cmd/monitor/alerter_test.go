package main

import (
	"testing"

	monitor "github.com/elojah/http-monitor"
	"github.com/elojah/http-monitor/mocks"
)

func TestStart(t *testing.T) {
	t.Run("0_alert_raised", func(t *testing.T) {
		var mappers monitor.Mappers
		mappers.TickMapper = &mocks.TickMapper{
			CountTickFunc: func(subset monitor.TickSubset) (int, error) {
				return 9, nil
			},
		}
		config := AlerterConfig{
			Treshold:       10,
			TriggerRange:   "2s",
			ReccurGap:      "2s",
			TriggerRecover: "10s",
			ReboundGap:     "10s",
		}
		alerter := NewAlerter(mappers)
		if err := alerter.Dial(config); err != nil {
			t.Error(err)
			return
		}
	})
}
