package telegram

import (
	"fmt"

	tgbotapi "gopkg.in/telebot.v3"
)

func (b *Bot) handlerEcho(c tgbotapi.Context) error {
	user, text := c.Sender(), c.Text()
	msg, err := b.bot.Send(user, text)
	if err != nil {
		return err
	}
	return c.Send(msg)
}

func (b *Bot) handlerStart(c tgbotapi.Context) error {
	fmt.Println(c.Message().Payload)
	return nil
}
