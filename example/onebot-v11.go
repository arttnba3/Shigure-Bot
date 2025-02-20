package main

import (
	"fmt"
	"github.com/arttnba3/Shigure-Bot/bot"
	"github.com/arttnba3/Shigure-Bot/bot/onebot/v11"
	"os"
	"time"
)

func MessageHandler(rawData []byte) {
	fmt.Println(fmt.Sprintf("[%v]", time.Now().Format("2006-01-02 15:04:05")), "Got message:", string(rawData))
}

func Logger(params ...interface{}) {
	fmt.Println(fmt.Sprintf("[%v] ", time.Now().Format("2006-01-02 15:04:05")) + fmt.Sprint(params...))
}

func main() {
	var err error
	Logger("Starting...")

	configJson, err := os.ReadFile("example/config.json")
	if err != nil {
		Logger("Unable to read config file, error: ", err.Error())
		return
	}

	bot, err := shigure.NewShigureBot("OneBot-V11", configJson, Logger, MessageHandler)
	if err != nil {
		Logger("Unable to create bot, error: ", err.Error())
		return
	}

	for i := 0; i < 2; i++ {
		var err error

		// TODO: Change this to a better way
		if i%2 == 0 {
			_, err = bot.Bot.(*onebot_v11_impl.V11Bot).SendGroupMsg(1145141919, fmt.Sprintf("Test from Shigure-%v", i), false)
		} else {
			_, err = bot.Bot.(*onebot_v11_impl.V11Bot).SendPrivateMsg(1145141919, fmt.Sprintf("Test from Shigure-%v", i), false)
		}

		if err != nil {
			fmt.Println("Got error: ", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
