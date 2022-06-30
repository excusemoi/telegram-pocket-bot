package main

import (
	"github.com/boltdb/bolt"
	"github.com/excusemoi/telegram-pocket-bot/pkg/config"
	"github.com/excusemoi/telegram-pocket-bot/pkg/repository"
	"github.com/excusemoi/telegram-pocket-bot/pkg/repository/boltdb"
	"github.com/excusemoi/telegram-pocket-bot/pkg/server"
	"github.com/excusemoi/telegram-pocket-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {

	cfg, err := config.Init()
	if err != nil {
		log.Panic(err)

	}

	b, err := tgbotapi.NewBotAPI(cfg.TelegramToken) // разумеется в пизду нахуй убрать
	if err != nil {
		log.Panic(err)
	}
	b.Debug = true

	cl, err := pocket.NewClient(cfg.PocketConsumerKey) // разумеется в пизду нахуй убрать
	if err != nil {
		log.Panic(err)
	}

	db, err := initDb()
	if err != nil {
		log.Panic(err)
	}

	tr := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(b, cl, tr, cfg.AuthServerUrl, cfg.Messages, cfg.Commands)

	authorizationServer := server.NewAuthorizationServer(cl, tr, cfg.TelegramBotUrl)

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err = authorizationServer.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDb() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessToken))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestToken))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return db, nil
}
