package main

import (
	"encoding/json"
	"fmt"
	"github.com/arttnba3/Shigure-Bot/api/onebot/v11/event"
	onebot_v11_api_message "github.com/arttnba3/Shigure-Bot/api/onebot/v11/message"
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
	event, ok2 := params[1].(onebot_v11_api_event.PrivateMessage)
	if !ok1 || !ok2 {
		Logger("Error: parameter type mismatch in PrivateMessageHandler")
		return
	}

	Logger("Got private message:", event.Message)

	_, err := bot.SendPrivateMsg(event.UserId, event.Message, false)
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
	event, ok2 := params[1].(onebot_v11_api_event.GroupMessage)
	if !ok1 || !ok2 {
		Logger("Error: parameter type mismatch in GroupMessageHandler")
		return
	}

	Logger("Got group message:", event.Message)
	_, err := bot.SendGroupMsg(event.GroupId, event.Message, false)
	if err != nil {
		Logger(fmt.Sprintf("Error: error sending group message: %v", err.Error()))
	}
}

func PrivateRecallHandler(params ...any) {
	if len(params) < 2 {
		Logger("Error: insufficient parameters for PrivateRecallHandler")
		return
	}

	bot, ok1 := params[0].(*onebot_v11_impl.V11Bot)
	event, ok2 := params[1].(onebot_v11_api_event.FriendMessageRecalled)
	if !ok1 || !ok2 {
		Logger("Error: parameter type mismatch in PrivateRecallHandler")
		return
	}

	Logger("Got private recall event, recalled message id: ", event.MessageID)

	recalledMsgData, err := bot.GetMsg(int32(event.MessageID)) // that's how the doc define...
	if err != nil {
		Logger(fmt.Sprintf("Error: error getting recall message: %v", err.Error()))
		return
	}

	recalledMsgJson, err := json.Marshal(recalledMsgData)
	if err != nil {
		Logger(fmt.Sprintf("Error: error marshalling recall message data: %v", err.Error()))
	}

	var recalledMsg onebot_v11_api_message.MessageArray
	err = json.Unmarshal(recalledMsgJson, &recalledMsg)
	if err != nil {
		Logger("Unable to parse recalled message as MessageArray")
		return
	}

	replyMsg := onebot_v11_api_message.MessageArray{
		onebot_v11_api_message.MessageSegment{
			Type: "text",
			Data: onebot_v11_api_message.MessageSegmentDataText{
				Text: "Detected recalled message: \n",
			},
		},
	}
	replyMsg = append(replyMsg, recalledMsg...)

	_, err = bot.SendPrivateMsg(event.UserID, replyMsg, false)
	if err != nil {
		Logger(fmt.Sprintf("Error: error sending private message: %v", err.Error()))
	}
}

func GroupRecallHandler(params ...any) {
	if len(params) < 2 {
		Logger("Error: insufficient parameters for GroupRecallHandler")
		return
	}

	bot, ok1 := params[0].(*onebot_v11_impl.V11Bot)
	event, ok2 := params[1].(onebot_v11_api_event.GroupMessageRecalled)
	if !ok1 || !ok2 {
		Logger("Error: parameter type mismatch in GroupRecallHandler")
		return
	}

	Logger("Got group recall event, recalled message id: ", event.MessageID)

	recalledMsgData, err := bot.GetMsg(int32(event.MessageID)) // that's how the doc define...
	if err != nil {
		Logger(fmt.Sprintf("Error: error getting recall message: %v", err.Error()))
		return
	}

	recalledMsgJson, err := json.Marshal(recalledMsgData)
	if err != nil {
		Logger(fmt.Sprintf("Error: error marshalling recall message data: %v", err.Error()))
	}

	var recalledMsg onebot_v11_api_message.MessageArray
	err = json.Unmarshal(recalledMsgJson, &recalledMsg)
	if err != nil {
		Logger("Unable to parse recalled message as MessageArray")
		return
	}

	replyMsg := onebot_v11_api_message.MessageArray{
		onebot_v11_api_message.MessageSegment{
			Type: "text",
			Data: onebot_v11_api_message.MessageSegmentDataText{
				Text: "Detected recalled message: \n",
			},
		},
	}

	// Note that if there's picture in the msg, it'll fail to send, maybe it's the problem of v11 impl backend :(
	replyMsg = append(replyMsg, recalledMsg...)

	_, err = bot.SendGroupMsg(event.GroupID, replyMsg, false)
	if err != nil {
		Logger(fmt.Sprintf("Error: error sending group message: %v", err.Error()))
	}
}

var handlers map[string]func(params ...any) = map[string]func(params ...any){
	"message_private":      PrivateMessageHandler,
	"message_group":        GroupMessageHandler,
	"notice_friend_recall": PrivateRecallHandler,
	"notice_group_recall":  GroupRecallHandler,
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
