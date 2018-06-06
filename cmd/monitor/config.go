package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/elojah/http-monitor/storage/redis"
)

const (
	defaultLogPath = "/var/log/access.log"
)

// Config is the configuration structure for monitor.
type Config struct {
	LogFile          string       `json:"log_file"`
	StatsInterval    uint         `json:"stats_interval"`
	TopDisplay       uint         `json:"top_display"`
	AlertReqPerSec   uint         `json:"alert_req_per_sec"`
	AlertTriggerTime uint         `json:"alert_trigger_time"`
	AlertReboundTime uint         `json:"alert_rebound_time"`
	AlertReccurTime  uint         `json:"alert_reccur_time"`
	Redis            redis.Config `json:"redis"`
}

// NewConfig creates a new config initialized from filepath in JSON format.
func NewConfig(filepath string) (Config, error) {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}

	var c Config
	err = json.Unmarshal(raw, &c)
	return c, err
}

// Check check if config fields are valid.
func (c Config) Check() error {
	if c.LogFile == "" {
		return errors.New("log filepath cannot be empty")
	}
	if c.StatsInterval == 0 {
		return errors.New("interval between each stats display cannot be 0")
	}
	if c.TopDisplay == 0 {
		return errors.New("number of top hits to display cannot be 0")
	}
	if c.AlertReqPerSec == 0 {
		return errors.New("number of requests required to trigger an alert cannot be 0")
	}
	if c.AlertTriggerTime == 0 {
		return errors.New("number of seconds required to trigger an alert cannot be 0")
	}
	if c.AlertReccurTime == 0 {
		return errors.New("number of seconds required between two alerts check cannot be 0")
	}
	return nil
}
