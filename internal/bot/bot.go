package bot

import (
	g "github.com/bahodurnazarov/CatFacts/internal/getFacts"
	lg "github.com/bahodurnazarov/CatFacts/pkg/utils"
	gt "github.com/bas24/googletranslatefree"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Select Language 🇬🇧"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Выберите Язык 🇷🇺"),
	),
)
var facts = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/facts"),
	),
)
var fact = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/fact"),
	),
)

func Bot(c echo.Context) error {
	token := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		lg.Errl.Fatal(err)
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

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		image := g.GetImage()
		//Extract the command from the Message.
		switch update.Message.Command() {
		case "start":
			msg.ReplyMarkup = numericKeyboard
		case "facts":
			factEN := g.GetFacts()
			factRU, _ := gt.Translate(factEN, "en", "ru")
			InsertToDB(factEN, factRU)
			msg.Text = image + "\n\n\"" + factEN + "\""
		case "fact":
			factEN := g.GetFacts()
			factRU, _ := gt.Translate(factEN, "en", "ru")
			InsertToDB(factEN, factRU)
			msg.Text = image + "\n\n\"" + factRU + "\""
		default:
			msg.Text = "Неверная команда!"
		}

		if update.Message.Text == "Select Language 🇬🇧" {
			msg.ReplyMarkup = facts
			msg.Text = "To get started, select the command \"Facts\""
		} else if update.Message.Text == "Выберите Язык 🇷🇺" {
			msg.ReplyMarkup = fact
			msg.Text = "Чтобы начать, выберете команду \"Факты\""
		}

		if _, err := bot.Send(msg); err != nil {
			lg.Errl.Fatal(err)
		}
	}
	return c.JSON(http.StatusOK, bot)
}
