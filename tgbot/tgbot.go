package main

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  "6656832916:AAG3f2U6RAizDuPsDu6snP3W6qqR_k4JpGU",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// echo
	b.Handle(tele.OnText, func(ctx tele.Context) error {
		user := ctx.Sender()
		text := ctx.Text()
		msg, err := b.Send(user, text)
		if err != nil {
			return err
		}
		return ctx.Send(msg)
	})

	// buttons
	var (
		// Universal markup builders.
		menu = &tele.ReplyMarkup{ResizeKeyboard: true}

		// Reply buttons.
		btnHelp     = menu.Text("Help")
		btnSettings = menu.Text("Settings")
	)

	menu.Reply(
		menu.Row(btnHelp),
		menu.Row(btnSettings),
	)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello!", menu)
	})

	// On reply button pressed (message)
	b.Handle(&btnHelp, func(c tele.Context) error {
		user := c.Sender()
		msg, err := b.Send(user, "Some help")
		if err != nil {
			return err
		}
		return c.Send(msg)
	})
	b.Start()
}
