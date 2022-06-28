package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

type Bot struct {
	bot          *tgbotapi.BotAPI
	pocketClient *pocket.Client
	redirectUrl  string
}

func NewBot(bot *tgbotapi.BotAPI, client *pocket.Client, redirectUrl string) *Bot {
	return &Bot{bot: bot, pocketClient: client, redirectUrl: redirectUrl}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}
	err = b.handleUpdates(updates)
	if err != nil {
		return err
	}
	return err
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	var err error
	for update := range updates {
		if update.Message == nil {
			continue
		} else if update.Message.IsCommand() {
			err = b.handleCommand(update.Message)
		} else {
			err = b.handleMessage(update.Message)
		}
	}
	return err
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)

}
