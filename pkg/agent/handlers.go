package agent

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *Agent) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	log.Printf("request url: %s method: %s\n", r.URL, r.Method)

	var update tgbotapi.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		w.WriteHeader(http.StatusNoContent)
		log.Printf("cannot decode update: %s\n", err)
		return
	}

	if update.Message == nil || update.Message.Text == "" {
		return
	}

	if err := a.SendResponse(update); err != nil {
		w.WriteHeader(http.StatusNoContent)
		log.Printf("cannot send response: %s\n", err)
		return
	}
}

func (a *Agent) SendResponse(update tgbotapi.Update) error {
	return errors.New("todo")
}
