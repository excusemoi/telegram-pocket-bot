package telegram

import (
	"github.com/excusemoi/telegram-pocket-bot/pkg/config"
	"github.com/excusemoi/telegram-pocket-bot/pkg/repository/boltdb"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

type Bot struct {
	bot             *tgbotapi.BotAPI
	pocketClient    *pocket.Client
	tokenRepository *boltdb.TokenRepository
	redirectUrl     string
	messages        *config.Messages
	commands        *config.Commands
}

func NewBot(bot *tgbotapi.BotAPI,
	client *pocket.Client,
	tr *boltdb.TokenRepository,
	redirectUrl string,
	messages *config.Messages,
	commands *config.Commands) *Bot {
	return &Bot{bot: bot,
		pocketClient:    client,
		tokenRepository: tr,
		redirectUrl:     redirectUrl,
		messages:        messages,
		commands:        commands}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}
	b.handleUpdates(updates)
	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		var handler func(*tgbotapi.Message) error
		if update.Message == nil {
			continue
		} else if update.Message.IsCommand() {
			handler = b.handleCommand
		} else {
			handler = b.handleMessage
		}
		if err := handler(update.Message); err != nil {
			if err = b.handleError(update.Message.Chat.ID, err); err != nil {
				log.Println(err)
			}

		}
	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)

}
