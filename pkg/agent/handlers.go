package agent

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"github.com/vadimistar/vaditafer/pkg/taf"
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
	lat, lng, timezone, err := a.opencageClient.ForwardGeocode(update.Message.Text)
	if err != nil {
		a.sendMessage(update.Message.From.ID, "Неверный запрос")
		return errors.Wrap(err, "forward geocode")
	}

	ids, err := a.aviaClient.ClosestAirports(lat, lng, a.radialDistance)
	if err != nil || len(ids) <= 0 {
		a.sendMessage(update.Message.From.ID, fmt.Sprintf("Не найдено аэропортов в радиусе %d км по координатам %f, %f", int(float64(a.radialDistance)*1.6), lat, lng))
		return errors.Wrap(err, "closest airports")
	}

	for _, id := range ids {
		taf, err := a.checkwxClient.Taf(id)
		if err != nil {
			log.Printf("cannot get taf for %s: %s", id, err.Error())
			continue
		}

		if err := localizeTime(taf, timezone); err != nil {
			return err
		}

		msg := tgbotapi.NewMessage(update.Message.From.ID, createMessage(taf))
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "Markdown"

		_, err = a.bot.Send(msg)
		if err != nil {
			return err
		}

		break
	}

	return nil
}

func localizeTime(taf *taf.Taf, location string) (err error) {
	defer func() {
		err = errors.Wrap(err, "localize time")
	}()

	loc, err := time.LoadLocation(location)
	if err != nil {
		return err
	}

	taf.CreatedAt = taf.CreatedAt.In(loc)
	taf.From = taf.From.In(loc)
	taf.To = taf.To.In(loc)
	for i := range taf.Forecasts {
		taf.Forecasts[i].Header.Start = taf.Forecasts[i].Header.Start.In(loc)
		taf.Forecasts[i].Header.End = taf.Forecasts[i].Header.End.In(loc)
	}

	return nil
}

func createMessage(taf *taf.Taf) string {
	var s strings.Builder
	fmt.Fprintf(&s, "**Прогноз создан __%s__\nДействует с __%s__ до __%s__**:\n\n", taf.CreatedAt.Format(timeLayout), taf.From.Format(timeLayout), taf.To.Format(timeLayout))

	for i, forecast := range taf.Forecasts {
		if i != 0 {
			s.WriteString("\n\n")
		}

		s.WriteString("**")
		if forecast.Header.Kind != "" {
			s.WriteString(forecast.Header.Kind + " ")
		}
		fmt.Fprintf(&s, "с %s до %s", forecast.Header.Start.Format(timeLayout), forecast.Header.End.Format(timeLayout))
		s.WriteString("**:\n")

		if forecast.Wind != nil && forecast.Wind.Speed != 0 {
			s.WriteString("ветер ")
			if forecast.Wind.Direction != "" {
				s.WriteString(forecast.Wind.Direction + " ")
			}
			fmt.Fprintf(&s, "%d м/c", forecast.Wind.Speed)
			if forecast.Wind.Gusts != 0 {
				fmt.Fprintf(&s, " (порывы %d м/c)", forecast.Wind.Gusts)
			}
			s.WriteString(", ")
		}

		if forecast.Visibility != 0 {
			if forecast.Visibility >= 9999 {
				s.WriteString("видимость 10 км и более, ")
			} else {
				fmt.Fprintf(&s, "видимость %d км, ", forecast.Visibility)
			}
		}

		for _, w := range forecast.Weather {
			fmt.Fprintf(&s, "%s, ", w)
		}

		for _, cloud := range forecast.CloudLayers {
			fmt.Fprintf(&s, "%s облачность на высоте %d м, ", cloud.Quantity, cloud.Height)
		}
	}

	result := s.String()
	return result[:len(result)-2]
}

const timeLayout = "02/01 15:04"

func (a *Agent) sendMessage(chatID int64, text string) error {
	_, err := a.bot.Send(tgbotapi.NewMessage(chatID, text))
	return err
}
