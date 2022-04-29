package main

import (
	"blockymon/monster"
	"blockymon/tray"
	"container/list"
	"fmt"
	"time"

	"github.com/getlantern/systray"
)

func main() {
	/*f, err := os.OpenFile("monsters.log", os.O_WRONLY|os.O_CREATE, 0755) //log file
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)*/

	var myMonster monster.Monster

	var petName string
	fmt.Print("Enter pet name: ")
	//	fmt.Scanf("%s", &petName)
	petName = "lulu"
	myMonster.Initialize("001", petName, time.Now())
	myAlbum := list.New()
	fmt.Println(myMonster.AlbumShot(time.Now(), myAlbum))

	go func() {
		var action string
		for {
			fmt.Print("Enter your action: ")
			fmt.Scanf("%s", &action)
			myMonster.TakeAction(action)
			fmt.Println(myMonster.AlbumShot(time.Now(), myAlbum))
		}
	}()
	systray.Run(tray.OnReady, tray.OnQuit)

}
