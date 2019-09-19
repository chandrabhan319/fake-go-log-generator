package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/pflag"
)

const usage = `
Usage: log [options]

Options:
  -l, --level string       log level. ("debug"|"info"|"warn"|"error")

`

func init() {
	pflag.Usage = printUsage
}

func main() {
	help := pflag.BoolP("help", "h", false, "Show usage")
	level := pflag.StringP("level", "l", "", "Log level")
	interval := pflag.IntP("interval", "i", 1, "Interval between random log generation")
	duration := pflag.IntP("duration", "d", 1, "Duration for which random log will be generated")

	pflag.Parse()

	if *help {
		printUsage()
		os.Exit(0)
	}

	logLevel, err := log.ParseLevel(*level)
	if err != nil {
		log.Errorf("invalid log level: %s", *level)
		os.Exit(1)
	}
	log.SetLevel(logLevel)

	randomLog(*interval, *duration, logLevel)
}

func printUsage() {
	fmt.Printf(usage)
}

func randomLog(i int, d int, lvl log.Level) {
	//create the log type slice
	var l []string
	switch lvl.String() {
	case "info":
		l = []string{"info", "warn", "error"}
	case "debug":
		l = []string{"debug", "info", "warn", "error"}
	case "warn", "warning":
		l = []string{"warn", "error"}
	case "error":
		l = []string{"error"}
	default:
		l = []string{"debug", "info", "warn", "error"}
	}
	rand.Seed(time.Now().Unix()) // initialize pseudo random generator
	for start := time.Now(); time.Since(start) < time.Duration(d)*time.Minute; {
		message := l[rand.Intn(len(l))]
		switch message {
		case "info":
			log.Infof("Generated dummy application %s log", message)
		case "debug":
			log.Debugf("Generated dummy application %s log", message)
		case "warn":
			log.Warnf("Generated dummy application %s log", message)
		case "error":
			log.Errorf("Generated dummy application %s log", message)
		}
		time.Sleep(time.Duration(i) * time.Second)
	}
}
