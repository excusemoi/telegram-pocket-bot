package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	errInvalidUrl     = errors.New("Invalid url")
	errUnauthorized   = errors.New("Unauthorized user")
	errUnableToSave   = errors.New("Unable to save url")
	errUnknownCommand = errors.New("Unknown command")
)

func (b *Bot) handleError(chatID int64, err error) error {
	msg := tgbotapi.NewMessage(chatID, b.messages.Errors.UnknownError)
	switch err {
	case errInvalidUrl:
		msg.Text = b.messages.Errors.InvalidUrl
	case errUnauthorized:
		msg.Text = b.messages.Errors.Unauthorized
	case errUnableToSave:
		msg.Text = b.messages.Errors.UnableToSave
	case errUnknownCommand:
		msg.Text = b.messages.Errors.UnknownCommand
	}
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
