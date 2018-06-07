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

	mappers := monitor.Mappers{}
	mappers.SectionMapper = redisx
	mappers.TickMapper = redisx

	logReader := NewLogReader(mappers)
	alerter := NewAlerter(mappers)
	if err := logReader.Dial(cfg.LogReader); err != nil {
		log.WithField("dial", "log_reader").Error(err)
		return
	}
	if err := alerter.Dial(cfg.Alerter); err != nil {
		log.WithField("dial", "alerter").Error(err)
		return
	}

	go func() {
		defer logReader.Close()
		if err := logReader.Start(); err != nil {
			log.WithField("routine", "log_reader").Error(err)
			return
		}
	}()
	go func() {
		defer alerter.Close()
		if err := alerter.Start(); err != nil {
			log.WithField("routine", "log_reader").Error(err)
			return
		}
	}()
	select {}
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	var filepath string
	flag.StringVar(&filepath, "c", "", "configuration file in JSON")

	flag.Parse()

	run(filepath)
}
