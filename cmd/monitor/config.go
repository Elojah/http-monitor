package main

import (
	"encoding/json"
	"io/ioutil"
)

const (
	defaultLogPath = "/var/log/access.log"
)

// Config is the configuration structure for monitor.
type Config struct {
	LogFile          string
	StatsInterval    uint
	AlertReqPerSec   uint
	AlertTriggerTime uint
	AlertReportTime  uint
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
