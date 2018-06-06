package main

import (
	"flag"

	log "github.com/sirupsen/logrus"

	monitor "github.com/elojah/http-monitor"
	"github.com/elojah/http-monitor/storage/redis"
)

func run(filepath string) {

	cfg, err := NewConfig(filepath)
	if err != nil {
		log.WithField("read", "config").Error(err)
		return
	}
	if err := cfg.Check(); err != nil {
		log.WithField("check", "config").Error(err)
		return
	}
	redisx := redis.NewService()
	if err := redisx.Dial(cfg.Redis); err != nil {
		log.WithField("dial", "redis").Error(err)
		return
	}

	services := monitor.Services{}
	services.SectionMapper = redisx
	services.TickMapper = redisx
	app := NewApp(services)
	if err := app.Dial(cfg); err != nil {
		log.WithField("dial", "app").Error(err)
		return
	}
	if err := app.Start(); err != nil {
		log.WithField("routine", "app").Error(err)
		return
	}
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	filepath := flag.String("c", "bin/config.json", "configuration file in JSON")

	run(*filepath)
}
