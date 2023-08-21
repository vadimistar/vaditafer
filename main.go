package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/vadimistar/vaditafer/pkg/agent"
	"github.com/vadimistar/vaditafer/pkg/avianation"
	"github.com/vadimistar/vaditafer/pkg/checkwx"
	"github.com/vadimistar/vaditafer/pkg/opencage"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	godotenv.Load()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		return errors.Wrap(err, "create telegram API")
	}

	webhook, err := bot.GetWebhookInfo()
	if err != nil {
		return errors.Wrap(err, "get webhook from telegram bot")
	}
	log.Printf("Webhook Info: %#v\n", webhook)

	var aviaClient avianation.Client
	checkwxClient := checkwx.New(os.Getenv("CHECKWX_API_KEY"))
	opencageClient := opencage.New(os.Getenv("OPENCAGE_API_KEY"))

	radialDistance, err := strconv.Atoi(os.Getenv("RADIAL_DISTANCE"))
	if err != nil {
		return errors.Wrap(err, "parse radial distance")
	}

	a := agent.New(bot, &aviaClient, checkwxClient, opencageClient, radialDistance)
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.HandleUpdate)
	return http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}
