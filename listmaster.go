package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"fmt"
	"io/ioutil"
    "strings"
    "errors"
)

type ListElement struct {
    Idx float32
    Element string
}

var (
	BotToken []byte
	err error
    IsDevelop = true
    Data = []ListElement{
        {1,"Аптека"},
        {1.1,"Канефрон"},
        {1.2,"Йод"},
        {2,"Зоо магазин"},
        {2.1,"Феликс 10 пакетиков"},
    }
    Rel = map[string]int{
        "/ADD":1,
        "/ДОБ":1,
        "/ФВВ":1,
        "/LIST":2,
        "/ПОКАЗ":2,
        "/ДШЫЕ":2,
        "/DONE":3,
        "/ГОТ":3,
        "/ВЩТУ":3,
        "/DEL":4,
        "/УДАЛ":4,
        "/ВУД":4,
    }
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

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	
	updates, err := bot.GetUpdatesChan(ucfg)

    for update := range updates {
        // log.Printf("[%#+v] %#+v", update.Message.From.UserName, update.Message.Text)

        ParseCommand(update.Message.Text)

        msg := tgbotapi.NewMessage(update.Message.Chat.	ID, update.Message.Text)
        msg.ReplyToMessageID = update.Message.MessageID

        bot.Send(msg)
    }}
//////////////////////
//
//   ParseCommand
//    
//////////////////////    
func ParseCommand(command string) (code int, idx float32, element string,err error) {
    fmt.Printf("Получили строку: %s\n",command)
    if strings.HasPrefix(command,"/") != true {
        return 0,0.0,"",errors.New("it's not command")
    }
    words := CheckWords(strings.Fields(command))
    code,err = GetCommandCode(words[0])
    if err != nil {
        return 0,0.0,"",err
    }
    return code,1.0,"test",nil
}

func GetCommandCode(in string) (code int,err error) {
    code = 0
    code = Rel[strings.ToUpper(in)]
    if code == 0 {
        return 0,errors.New("Unknown command")
    } else {
        return code,nil
    }
}

func CheckWords (words []string) []string {
    var (
        out []string
        isWord = false
    )
    
    for i := range words {
        if isWord == false {
            out = append(out,words[i])
        } else {
            out[len(out)-1] = out[len(out)-1]+" "+words[i]
        }

        if isWord == false && strings.HasPrefix(words[i],"\"") {
            isWord = true
        } else if isWord == true && strings.HasSuffix(words[i],"\"") {
            isWord = false
        }
    }
    return out
}