package tray

import (
	"os"
	"os/signal"
	"syscall"

	"blockymon/icon"

	"github.com/getlantern/systray"
)

func OnReady() {
	systray.SetIcon(icon.Data)

	mQuit := systray.AddMenuItem("Quit", "Quit example tray application")

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case <-mQuit.ClickedCh:
			systray.Quit()
		case <-sigc:
			systray.Quit()
		}
	}
}

func OnQuit() {
}
