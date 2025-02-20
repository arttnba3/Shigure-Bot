package main

import (
	"fmt"
	"github.com/arttnba3/Shigure-Bot/api/onebot/v11/event"
	"github.com/arttnba3/Shigure-Bot/bot"
	"github.com/arttnba3/Shigure-Bot/bot/onebot/v11"
	"os"
	"time"
)

func Logger(params ...any) {
	fmt.Println(fmt.Sprintf("[%v] ", time.Now().Format("2006-01-02 15:04:05")) + fmt.Sprint(params...))
}

func PrivateMessageHandler(params ...any) {
	if len(params) < 2 {
		Logger("Error: insufficient parameters for PrivateMessageHandler")
		return
	}

	bot, ok1 := params[0].(*onebot_v11_impl.V11Bot)
	message, ok2 := params[1].(onebot_v11_api_event.PrivateMessage)
	if !ok1 || !ok2 {
		Logger("Error: parameter type mismatch in PrivateMessageHandler")
		return
	}

	Logger(fmt.Sprintf("[%v]", time.Now().Format("2006-01-02 15:04:05")), "Got private message:", message.Message)

	_, err := bot.SendPrivateMsg(message.UserId, message.Message, false)
	if err != nil {
		Logger(fmt.Sprintf("Error: error sending private message: %v", err.Error()))
	}
}

func GroupMessageHandler(params ...any) {
	if len(params) < 2 {
		Logger("Error: insufficient parameters for GroupMessageHandler")
		return
	}

	bot, ok1 := params[0].(*onebot_v11_impl.V11Bot)
	message, ok2 := params[1].(onebot_v11_api_event.GroupMessage)
	if !ok1 || !ok2 {
		Logger("Error: parameter type mismatch in GroupMessageHandler")
		return
	}

	Logger(fmt.Sprintf("[%v]", time.Now().Format("2006-01-02 15:04:05")), "Got group message:", message.Message)
	_, err := bot.SendGroupMsg(message.GroupId, message.Message, false)
	if err != nil {
		Logger(fmt.Sprintf("Error: error sending group message: %v", err.Error()))
	}
}

var handlers map[string]func(params ...any) = map[string]func(params ...any){
	"receive_private_message": PrivateMessageHandler,
	"receive_group_message":   GroupMessageHandler,
}

func main() {
	var err error
	Logger("Starting...")

	configJson, err := os.ReadFile("example/config.json")
	if err != nil {
		Logger("Unable to read config file, error: ", err.Error())
		return
	}

	_, err = shigure.NewShigureBot("OneBot-V11", configJson, Logger, handlers)
	if err != nil {
		Logger("Unable to create bot, error: ", err.Error())
		return
	}

	Logger("Sleeping...")

	for {
		time.Sleep(1 * time.Second)
	}
}
