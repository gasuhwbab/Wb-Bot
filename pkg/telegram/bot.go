package telegram

import tgbotapi "gopkg.in/telebot.v3"

type Bot struct {
	bot *tgbotapi.Bot
}

func NewBot(bot *tgbotapi.Bot) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() error {
	b.bot.Handle("/start", b.handlerStart)
	b.bot.Handle(tgbotapi.OnText, b.handlerEcho)

	b.bot.Start()
	return nil
}
