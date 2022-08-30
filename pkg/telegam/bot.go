package telegam

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	o   *owm.CurrentWeatherData
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start(weather func(message *tgbotapi.Message) (answer string, err error)) error {

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		log.Fatal(err)
	}
	b.handleUpdates(updates, weather)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel, weather func(message *tgbotapi.Message) (answer string, err error)) {
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			weather, err := weather(update.Message)
			if err != nil {
				weather = "Отправьте мне локацию"
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, weather)
			b.bot.Send(msg)

			fmt.Println(weather)
		}

	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u), nil
}
