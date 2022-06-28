package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart              = "start"
	commandUnknown            = "Я не знаю такой команды"
	commandReplyStartTemplate = "Привет! Чтобы сохранять ссылки в своем Pocket аккаунте, для начала тебе необходимо дать мне на это " +
		"доступ. Для этого переходи по ссылке:\n%s"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	authLink, err := b.generateAuthorizationLink(message.Chat.ID)
	_, err = b.bot.Send(tgbotapi.NewMessage(message.Chat.ID,
		fmt.Sprintf(commandReplyStartTemplate, authLink)))
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	_, err := b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, commandUnknown))
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return err
}
