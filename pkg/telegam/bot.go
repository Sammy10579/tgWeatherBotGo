package telegam

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot            *tgbotapi.BotAPI
	massageHandler func(message *tgbotapi.Message) (answer string, err error)
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) MassageHandler(fn func(message *tgbotapi.Message) (answer string, err error)) {
	b.massageHandler = fn
	return
}

func (b *Bot) Start() error {

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		log.Fatal(err)
	}
	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {

			weather, err := b.massageHandler(update.Message)
			if err != nil {
				fmt.Println(err)
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
