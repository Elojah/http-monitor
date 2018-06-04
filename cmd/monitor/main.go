package main

import (
	"flag"

	log "github.com/sirupsen/logrus"

	"github.com/elojah/http-monitor/storage/mem"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	filepath := flag.String("c", "bin/config.json", "configuration file in JSON")

	cfg, err := NewConfig(*filepath)
	if err != nil {
		log.Error(err)
		return
	}
	memx := mem.NewService()
	app := NewApp(memx)
	if err := app.Dial(cfg); err != nil {
		log.Error(err)
		return
	}
	_ = app
}
