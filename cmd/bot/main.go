package main

import (
	"github.com/excusemoi/telegram-pocket-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {
	b, err := tgbotapi.NewBotAPI("5516775333:AAHEepCjVdoZyLPI56WCZteps_SAYRZS_84") // разумеется в пизду нахуй убрать
	if err != nil {
		log.Panic(err)
	}
	cl, err := pocket.NewClient("102592-ed66400e5014e3e94c36c38") // разумеется в пизду нахуй убрать
	if err != nil {
		log.Panic(err)
	}

	b.Debug = true
	telegramBot := telegram.NewBot(b, cl, "http://localhost/")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
