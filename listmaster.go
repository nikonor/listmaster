package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"fmt"
	"io/ioutil"
)

var (
	BotToken []byte
	err error
)

func init () {
	BotToken, err = ioutil.ReadFile("./listmaster.key")
	if err != nil {
		log.Panic(err)
		
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(string(BotToken))
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%#+v\n",bot);

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	
	updates, err := bot.GetUpdatesChan(ucfg)

    for update := range updates {
        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
        msg.ReplyToMessageID = update.Message.MessageID

        bot.Send(msg)
    }}
