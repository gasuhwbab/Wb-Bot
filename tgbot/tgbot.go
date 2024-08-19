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

	// buttons
	var (
		// Universal markup builders.
		menu = &tele.ReplyMarkup{ResizeKeyboard: true}
		// Reply buttons.
		btnSlotsInforming                = menu.Text("Начать информирование о слотах")
		btnAddWarehouseTracking          = menu.Text("Добавить склад на отслеживание")
		btnAddSortingCentreTracking      = menu.Text("Добавить СЦ на отслеживание")
		btnStopInformingSlots            = menu.Text("Приостановить информирование о слотах")
		btnDeleteInformingWarehouseSlots = menu.Text("Удалить информирование о слотах на складах")
	)

	menu.Reply(
		menu.Row(btnSlotsInforming),
		menu.Row(btnAddWarehouseTracking),
		menu.Row(btnAddSortingCentreTracking),
		menu.Row(btnStopInformingSlots),
		menu.Row(btnDeleteInformingWarehouseSlots),
	)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello!", menu)
	})

	b.Handle(&btnSlotsInforming, func(ctx tele.Context) error {
		user, text := ctx.Sender(), ctx.Text()
		msg, err := b.Send(user, text)
		if err != nil {
			return err
		}
		return ctx.Send(msg)
	})

	b.Handle(&btnAddWarehouseTracking, func(ctx tele.Context) error {
		user, text := ctx.Sender(), ctx.Text()
		msg, err := b.Send(user, text)
		if err != nil {
			return err
		}
		return ctx.Send(msg)
	})

	b.Handle(&btnAddSortingCentreTracking, func(ctx tele.Context) error {
		user, text := ctx.Sender(), ctx.Text()
		msg, err := b.Send(user, text)
		if err != nil {
			return err
		}
		return ctx.Send(msg)
	})

	b.Handle(&btnStopInformingSlots, func(ctx tele.Context) error {
		user, text := ctx.Sender(), ctx.Text()
		msg, err := b.Send(user, text)
		if err != nil {
			return err
		}
		return ctx.Send(msg)
	})

	b.Handle(&btnDeleteInformingWarehouseSlots, func(ctx tele.Context) error {
		user, text := ctx.Sender(), ctx.Text()
		msg, err := b.Send(user, text)
		if err != nil {
			return err
		}
		return ctx.Send(msg)
	})

	b.Start()
}
