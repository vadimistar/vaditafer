package agent

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Agent struct {
	bot *tgbotapi.BotAPI
}

func New(bot *tgbotapi.BotAPI) *Agent {
	return &Agent{
		bot: bot,
	}
}
