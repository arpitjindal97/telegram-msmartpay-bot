package main

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
	"github.com/tebeka/selenium"
	"os"
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

var telegram_bot_api string
func main() {
	bot, err := tgbotapi.NewBotAPI(telegram_bot_api)
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	const (
		// These paths will be different on your system.
		seleniumPath    = "selenium-server-standalone.jar"
		geckoDriverPath = "geckodriver"
		port            = 8081
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		messageProcessor(bot,update)

	}
	fmt.Println("will not Reach here")

}
func messageProcessor(bot *tgbotapi.BotAPI,update tgbotapi.Update) {

	//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Processing ...")
	bot.Send(msg)
	var wd selenium.WebDriver
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
			caps := selenium.Capabilities{"browserName": "firefox"}
			wd,_ = selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 8081))
			
			vals[0] = main1(vals,wd)

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
			wd.Close()
			wd.Quit()
		},
	}.Do()
}