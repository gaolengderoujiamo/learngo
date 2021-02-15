package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type DefaultFieldHook struct {
}

func (hook *DefaultFieldHook) Fire(entry *log.Entry) error {
	entry.Data["appName"] = "MyAppName"
	return nil
}

func (hook *DefaultFieldHook) Levels() []log.Level {
	return log.AllLevels
}

func main() {
	log.AddHook(&DefaultFieldHook{})
	log.WithFields(log.Fields{
		"a":    "AAA",
		"date": time.Now(),
	}).Info("hook test")
}
