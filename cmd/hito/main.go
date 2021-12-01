package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ad9311/hito/internal/app"
	"github.com/ad9311/hito/internal/console"
	"github.com/ad9311/hito/internal/handler"
)

var config *app.InitConfig
var data *app.InitData

func main() {
	setupCloseHandler()
	console.InitMessage()

	config, data = app.New()

	handler.New(config, data)

	console.ServerInfo(config.PortNumber)
	console.AssertPanic(serve().ListenAndServe())
}

func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		mainClose()
		os.Exit(0)
	}()
}

func mainClose() {
	console.ClosingMessage("CTRL+C signal detected\nClosing application...")
	err := config.ConnDB.SQL.Close()
	console.AssertPanic(err)
	err = serve().Close()
	console.AssertPanic(err)
	console.Message("\nGooddbye. =D")
}
