package telegram

import (
	"context"
	"fmt"
	"github.com/excusemoi/telegram-pocket-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) generateAuthorizationLink(chatId int64) (string, error) {
	redirectUrlWithChatId := b.generateRedirectUrl(chatId)
	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectUrlWithChatId)
	if err != nil {
		return "", err
	}
	if err = b.tokenRepository.Save(chatId, requestToken, repository.RequestToken); err != nil {
		return "", err
	}
	return b.pocketClient.GetAuthorizationURL(requestToken, redirectUrlWithChatId)
}

func (b *Bot) generateRedirectUrl(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", b.redirectUrl, chatID)
}

func (b *Bot) getAccessToken(chatID int64) (string, error) {
	return b.tokenRepository.Get(chatID, repository.AccessToken)
}

func (b *Bot) initAuthorizationProcess(message *tgbotapi.Message) error {
	authLink, err := b.generateAuthorizationLink(message.Chat.ID)
	_, err = b.bot.Send(tgbotapi.NewMessage(message.Chat.ID,
		fmt.Sprintf(b.messages.Responses.Start, authLink)))
	return err
}
