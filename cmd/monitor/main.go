package main

import (
	"flag"

	log "github.com/sirupsen/logrus"

	"github.com/elojah/http-monitor/storage/redis"
)

func run(filepath string) {

	cfg, err := NewConfig(filepath)
	if err != nil {
		log.Error(err)
		return
	}
	redisx := redis.NewService()
	if err := redisx.Dial(cfg.Redis); err != nil {
		log.Error(err)
		return
	}

	app := NewApp(redisx)
	if err := app.Dial(cfg); err != nil {
		log.Error(err)
		return
	}
	if err := app.Start(); err != nil {
		log.Error(err)
		return
	}
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	filepath := flag.String("c", "bin/config.json", "configuration file in JSON")

	run(*filepath)
}
