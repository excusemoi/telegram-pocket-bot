package telegram

import (
	"context"
	"github.com/excusemoi/telegram-pocket-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"net/url"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case b.commands.Start:
		return b.handleStartCommand(message)
	default:
		return errUnknownCommand
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	_, err := b.tokenRepository.Get(message.Chat.ID, repository.AccessToken)
	if err != nil {
		return b.initAuthorizationProcess(message)
	}
	_, err = b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, b.messages.Responses.AlreadyAuthorized))
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	_, err := b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, b.messages.Errors.UnknownCommand))
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.Responses.UrlSavedSuccessfully)

	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		return errInvalidUrl
	}

	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return errUnauthorized
	}

	if err = b.pocketClient.Add(context.Background(), pocket.AddInput{
		URL:         message.Text,
		AccessToken: accessToken,
	}); err != nil {
		return errUnableToSave
	}

	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
