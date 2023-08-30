package main

import (
	"github.com/iarsham/sender/app"
	"gopkg.in/telebot.v3"
	"os"
	"strconv"
	"time"
)

func init() {
	app.GetEnvs()
	app.GetLogger()
	app.FileManager()
}

func main() {
	Log := app.Log
	ttlTime, _ := strconv.Atoi(os.Getenv("TTL"))
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &telebot.LongPoller{Timeout: time.Duration(ttlTime) * time.Second},
	})
	if err != nil {
		Log.Fatal(err)
	}
	Log.Info("Sender Started Successfully")
	bot.Handle(telebot.OnText, app.HelloHandler)
	go app.SendNewLine(bot, Log)
	bot.Start()
}
