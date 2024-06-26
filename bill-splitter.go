package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {

	startTime := time.Now()

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	log.Println("Bot Token: ", botToken)

	pref := tele.Settings{
		Token:  botToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

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

	b.Handle("/balances", func(c tele.Context) error {
		log.Println("/balances command received")
		p := new(Person)
		persons, err := p.FindAll()
		if err != nil {
			log.Println("Error: ", err)
			return c.Send("Error: " + err.Error())
		}
		return c.Send(fmt.Sprintf("Balances: %v", persons))
	})

	b.Start()
}
