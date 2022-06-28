package boltdb

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/excusemoi/telegram-pocket-bot/pkg/repository"
	"strconv"
)

type TokenRepository struct {
	db *bolt.DB
}

func NewTokenRepository(db *bolt.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (r *TokenRepository) Save(chatID int64, token string, bucket repository.Bucket) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Put(int64ToBytes(chatID), []byte(token))
	})
}
func (r *TokenRepository) Get(chatID int64, bucket repository.Bucket) (string, error) {
	var (
		err   error
		token string
	)
	err = r.db.View(func(tx *bolt.Tx) error {
		token = string(tx.Bucket([]byte(bucket)).Get(int64ToBytes(chatID)))
		return nil
	})
	if token == "" && err == nil {
		err = errors.New(fmt.Sprintf("В базе данных нет токена по chat_id = %d", chatID))
	}
	return token, err
}

func int64ToBytes(n int64) []byte {
	return []byte(strconv.FormatInt(n, 10))
}
