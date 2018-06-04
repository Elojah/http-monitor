package redis

import "errors"

// Config is the config required for redis service.
type Config struct {
	Addr     string
	Password string
	DB       string
}

func (c Config) Check() error {
	if c.Addr == "" {
		return errors.New("redis address url cannot be empty")
	}
	if c.DB == "" {
		return errors.New("redis DB cannot be empty")
	}
	return nil
}
