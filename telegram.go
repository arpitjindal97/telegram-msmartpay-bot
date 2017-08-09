package main

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
)

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}
type Exception interface{}

func Throw(e Exception) {
	panic(e)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

func main() {
	bot, err := tgbotapi.NewBotAPI("392780450:AAEdek0ZeKYbVUBP5HpvoPVnAF8uumzHbQw")
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		Block{

			Try: func() {

				str := update.Message.Text
				count := 0
				var vals [3]string

				for _, r := range str {
					if r == ',' {
						count++
						continue
					}
					vals[count] = vals[count] + string(r)
				}
				if count == 0{
					panic( "invalid syntax")
				}

				vals[0] = main1(vals)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, vals[0])
				//msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

			},
			Catch: func(e Exception) {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error Occured !!!")
				fmt.Println(e)
				bot.Send(msg)

			},
			Finally: func() {
				//fmt.Println("Finally called")
			},
		}.Do()
	}
	fmt.Println("will not Reach here")

}
