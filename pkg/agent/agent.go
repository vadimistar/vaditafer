package agent

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vadimistar/vaditafer/pkg/avianation"
	"github.com/vadimistar/vaditafer/pkg/checkwx"
	"github.com/vadimistar/vaditafer/pkg/opencage"
)

type Agent struct {
	bot            *tgbotapi.BotAPI
	aviaClient     *avianation.Client
	checkwxClient  *checkwx.Client
	opencageClient *opencage.Client
	radialDistance int
}

func New(bot *tgbotapi.BotAPI, aviaClient *avianation.Client, checkwxClient *checkwx.Client, opencageClient *opencage.Client, radialDistance int) *Agent {
	return &Agent{
		bot:            bot,
		aviaClient:     aviaClient,
		checkwxClient:  checkwxClient,
		opencageClient: opencageClient,
		radialDistance: radialDistance,
	}
}
