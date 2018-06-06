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

// AlerterConfig is the configuration object for alerter.
type AlerterConfig struct {
	Treshold     uint   `json:"treshold"`
	TriggerRange string `json:"trigger_range"`
	ReboundGap   string `json:"rebound_gap"`
	ReccurGap    string `json:"reccur_gap"`
}

// Check check if config fields are valid.
func (c AlerterConfig) Check() error {
	if c.Treshold == 0 {
		return errors.New("missing treshold field")
	}
	if c.TriggerRange == "" {
		return errors.New("missing trigger_range field")
	}
	if c.ReccurGap == "" {
		return errors.New("missing reccur_gap field")
	}
	return nil
}

// LogReaderConfig is the configuration object for log reader.
type LogReaderConfig struct {
	LogFile    string `json:"log_file"`
	StatsGap   string `json:"stats_gap"`
	TopDisplay uint   `json:"top_display"`
}

// Check check if config fields are valid.
func (c LogReaderConfig) Check() error {
	if c.LogFile == "" {
		return errors.New("missing log_file field")
	}
	if c.StatsGap == "" {
		return errors.New("missing stats_gap field")
	}
	if c.TopDisplay == 0 {
		return errors.New("missing top_display field")
	}
	return nil
}

// Config is the configuration structure for monitor.
type Config struct {
	Alerter   AlerterConfig   `json:"alerter"`
	LogReader LogReaderConfig `json:"log_reader"`
	Redis     redis.Config    `json:"redis"`
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
	if err := c.Alerter.Check(); err != nil {
		return err
	}
	if err := c.LogReader.Check(); err != nil {
		return err
	}
	return c.Redis.Check()
}
