package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"fmt"
	"io/ioutil"
    "strings"
    "errors"
    "strconv"
    _ "github.com/lib/pq"
    "database/sql"    
    "bufio"
    "os"    
)

type ListElement struct {
    Idx float32
    Element string
}

var (
	BotToken []byte
	err error
    IsDevelop = true
    DevData = []ListElement{
        {1,"Аптека"},
        {1.001,"Канефрон"},
        {1.002,"Йод"},
        {2,"Зоо магазин"},
        {2.001,"Феликс 10 пакетиков"},
        {3,"Овощи, фрукты"},
        {3.001,"огурцы"},
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
    RelShort = []string{"","/ADD","/LIST","/DONE","/DEL"}
    db_connect_string = ""
    db *sql.DB    
)

func init () {
    var tt []string

	BotToken, err = ioutil.ReadFile("./listmaster.key")
	if err != nil {
		log.Panic(err)
	}

    var file *os.File
    file, err = os.Open("./listmaster.db")
    defer file.Close()    
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
            t := scanner.Text()
            t = strings.TrimSpace(t)
            if len(t) > 0 {
                tt = append(tt,t)
            }
    }    
    db_connect_string = strings.Join(tt," ")
}

func main() {
    db, err = sql.Open("postgres", db_connect_string)
    if err != nil {
        panic(err)
    }
    defer db.Close()    

	bot, err := tgbotapi.NewBotAPI(string(BotToken))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60

	updates, err := bot.GetUpdatesChan(ucfg)

    // Lists := []ListElement{}
    // if IsDevelop {
    //     Lists = DevData
    // }

    for update := range updates {

        CID := update.Message.Chat.ID;

        rows, err_db := db.Query(`SELECT idx,item FROM tt where chat_id = $1 and show='t'`,CID)
        if err_db != nil {
            fmt.Printf("%#v",err_db)
            panic("Error on select");
        }

        Lists := []ListElement{}
        for rows.Next() {
            var idx float32
            var text string
            err = rows.Scan(&idx, &text)
            new_el := ListElement{idx,text}
            Lists = append(Lists,new_el)

        }        

        code, idx, element,err := ParseCommand(update.Message.Text,Lists)

        if code == 1 {
            Lists = AddElement(Lists,idx,element)
        }

        msg_text := ""
        if err != nil {
            msg_text = err.Error()
        } else {
            msg_text = fmt.Sprintf("ParseCommand returned: \n    code=%d,\n    idx =%v,\n    el  =%s\n",code,idx,element)
            msg_text = ShowList(Lists)
        }

        // msg := tgbotapi.NewMessage(update.Message.Chat.	ID, update.Message.Text)
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, msg_text)
        if err != nil { 
            msg.ReplyToMessageID = update.Message.MessageID
        }
        log.Printf("Chat ID=%#v\n\tcommand=%s\n\tsourse=\"%s\"\n\titem=\"%s\"\n",CID,RelShort[code],update.Message.Text,element)
        bot.Send(msg)
    }}
//////////////////////
//
//   ParseCommand
//
//////////////////////
func ParseCommand(command string, lists []ListElement) (code int, idx float32, element string,err error) {
    if strings.HasPrefix(command,"/") != true {
        return 0,0.0,"",errors.New("it's not command")
    }

    words := CheckWords(strings.Fields(command))
    code,err = GetCommandCode(words[0])
    if err != nil {
        return 0,0.0,"",err
    }
    if len(words) >= 2 {
        idx,_ = GetListIdx(code,words[1],lists)
        element = words[len(words)-1]
    }
    return code,idx,element,nil
}

func AddElement(lists []ListElement, idx float32, element string) []ListElement {
    ret := []ListElement{}
    if idx == 0 {
        ret = lists
        ret = append(ret,ListElement{GetMaxIdx(lists),element})
    } else {
        lastId := float32(0.0)
        for _,e := range lists {
            if element != "" && e.Idx == (idx+1) {
                new_el := ListElement{lastId+0.001,element}
                ret = append(ret,new_el)
                element = ""
            } else {
                lastId = e.Idx
            }
            ret = append(ret,e)
        }

        if element != "" {
            new_el := ListElement{lastId+0.001,element}
            ret = append(ret,new_el)
        }
    }
    return ret
}

func GetMaxIdx(lists []ListElement) float32 {
    if  len(lists) > 0 {
        return float32(int(lists[len(lists)-1].Idx+1))    
    } else {
        return 1;
    }
    
}

func ShowList(lists []ListElement) string {
    out := ""
    for _,e := range lists {
        if (e.Idx - float32(int(e.Idx))) == 0 {
            out = out + fmt.Sprintf("%.f. %s\n",e.Idx,e.Element)
        } else {
            out = out + fmt.Sprintf("    %.3f. %s\n",e.Idx,e.Element)
        }
    }
    return out
}

func GetListIdx(code int, word string,lists []ListElement) (idx float32,err error){
    idx64,err := strconv.ParseFloat(word,32)
    if err != nil {
        for _,e := range lists {
            if e.Element == word {
                idx64 = float64(e.Idx)
            }
        }
    }
    return float32(idx64),nil
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
        w := words[i]

        if strings.HasPrefix(w,"\"") {
            w = w[1:]
        }
        if strings.HasSuffix(w,"\"") {
            w = w[:len(w)-1]
        }

        if isWord == false {
            out = append(out,w)
        } else {
            out[len(out)-1] = out[len(out)-1]+" "+w
        }

        if isWord == false && strings.HasPrefix(words[i],"\"") {
            isWord = true
        } else if isWord == true && strings.HasSuffix(words[i],"\"") {
            isWord = false
        }
    }
    return out
}
