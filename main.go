package main

import (
	"blockymon/monster"
	"blockymon/terminal"
	"blockymon/tray"
	"os"
	"sync"
	"time"

	"github.com/getlantern/systray"
	log "github.com/sirupsen/logrus"
)

func main() {
	f, err := os.OpenFile("monsters.log", os.O_WRONLY|os.O_CREATE, 0755) //log file
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)

	var wg sync.WaitGroup

	var monster monster.Monster
	monster.Initialize()
	go systray.Run(tray.OnReady, tray.OnQuit)
	go terminal.Input()
	wg.Wait()
	time.Sleep(time.Hour)
}
