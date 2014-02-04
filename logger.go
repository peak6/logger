package logger

import (
	"flag"
	"log"
	"os"
	"time"
)

var Lerr log.Logger
var Linfo log.Logger
var flags = log.Ltime | log.Lmicroseconds
var withLineNumbers = flag.Bool("ll", false, "Enable line numbers in log statements")

func InitLogger() {
	if *withLineNumbers {
		flags |= log.Lshortfile
	}
	Lerr = *log.New(os.Stdout, "E", flags)
	Linfo = *log.New(os.Stdout, "I", flags)
	setupLoggers()
	go resetLoggers()
}

func setupLoggers() {
	t := time.Now()
	dstr := t.Format("0102")
	Lerr.SetPrefix("E" + dstr + " ")
	Linfo.SetPrefix("I" + dstr + " ")
}

func resetLoggers() {
	for {
		t := time.Now()
		tonightMidnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1)
		tick := tonightMidnight.Sub(t)
		if tick > 1 {
			ticker := time.NewTicker(tick)
			Linfo.Println("Will reset in", tick, "at", tonightMidnight)
			<-ticker.C
			setupLoggers()
		} else {
			Lerr.Println("Failed to setup new date reset timer, please bounce this app")
			return
		}
	}

}
