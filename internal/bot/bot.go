package bot

import (
	"fmt"
	g "github.com/bahodurnazarov/CatFacts/internal/getFacts"
	gt "github.com/bas24/googletranslatefree"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Bot(c echo.Context) error {
	bot, err := tgbotapi.NewBotAPI("6263473565:AAEopw_EaoLRP83Io-aniGU2w7m6T1nfcDk")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "start":
			msg.Text = "Чтобы начать, выберете команду \"Факты\""
		case "facts":
			factEN := g.GetFacts()
			factRU, _ := gt.Translate(factEN, "en", "ru")
			msg.Text = "ENGLISH: 🇬🇧 " + "\"" + factEN + "\"" + "\n" + "\nRUSSIAN: 🇷🇺 " + "\"" + factRU + "\""
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
	return c.JSON(http.StatusOK, bot)
}

func Route() {
	var timeDuration = 10
	timer := time.NewTicker(time.Second * time.Duration(timeDuration))
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			resp, err := http.Get("http://localhost:1323/bot")

			if err != nil {
				log.Fatal(err)
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(body))
		}
	}
}
