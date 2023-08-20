package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/vadimistar/yt-audio-bot-serverless/pkg/agent"
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

	a := agent.New(bot)
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.HandleUpdate)
	return http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}
