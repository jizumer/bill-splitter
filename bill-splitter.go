package main

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	log.Println("Bot Token: ", botToken)

	pref := tele.Settings{
		Token:  botToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	startTime := time.Now()

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		log.Println("Hello command received")
		return c.Send("Hello!")
	})

	b.Handle("/info", func(c tele.Context) error {
		log.Println("/info command received")
		return c.Send("App started at: " + startTime.String())
	})

	b.Start()
}
