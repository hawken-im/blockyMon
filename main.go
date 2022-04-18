package main

import (
	"blockymon/actions"
	"blockymon/monster"
	"blockymon/tray"
	"fmt"
	"os"

	"github.com/getlantern/systray"
	log "github.com/sirupsen/logrus"
)

func main() {
	f, err := os.OpenFile("YP.log", os.O_WRONLY|os.O_CREATE, 0755) //log file
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)

	var monster monster.Monster
	monster.Initialize()
	systray.Run(tray.OnReady, tray.OnQuit)
	for {
		var action string
		fmt.Print("Enter your action: ")
		fmt.Scanf("%s", &action)

		switch action {
		case "pet":
			fmt.Println(action)
			monster.Moody(actions.Pet())
			log.Info("Ped monster, mood change:", actions.Pet(), ". mood became:", monster.Mood)
		case "clean":
			fmt.Println(action)
		case "feed":
			fmt.Println(action)
		case "heal":
			fmt.Println(action)
		default:
			fmt.Println("unknown action")
		}
	}

}
