package telegram

import (
	"context"
	"fmt"
	"github.com/excusemoi/telegram-pocket-bot/pkg/repository"
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
