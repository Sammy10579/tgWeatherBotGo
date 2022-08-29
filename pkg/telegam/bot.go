package telegam

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start(
	weatherByCity func(string) (string, error),
	weatherByLocation func(float64, float64) (string, error),
) error {

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		log.Fatal(err)
	}
	b.handleUpdates(updates, weatherByCity, weatherByLocation)

	return nil
}

func (b *Bot) handleUpdates(
	updates tgbotapi.UpdatesChannel,
	weatherByCity func(string) (string, error),
	weatherByLocation func(float64, float64) (string, error),
) {
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Location != nil {
				weatherInLocation, err := weatherByLocation(update.Message.Location.Longitude, update.Message.Location.Latitude)
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отправьте мне локацию")
					b.bot.Send(msg)
					continue
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, weatherInLocation)
				b.bot.Send(msg)
			} else {
				weatherInCity, err := weatherByCity(update.Message.Text)
				if err != nil {
					fmt.Println(err)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите пожалуйста существующий город")
					b.bot.Send(msg)
					continue
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, weatherInCity)
				b.bot.Send(msg)

				fmt.Println(weatherInCity)
			}

		}
	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u), nil
}
