package onebot_v11_impl

import (
	"encoding/json"
	"fmt"
	"github.com/arttnba3/Shigure-Bot/api/onebot/v11/event"
)

func RevokeOperators(cmd string, operators map[string]func(...any), params ...any) {
	operator := operators[cmd]
	if operator != nil {
		operator(params...)
	}
}

func RevokeHandlers(cmd string, handlers map[string]func(...any), params ...any) {
	handler := handlers[cmd]
	if handler != nil {
		handler(params...)
	}
}

func (bot *V11Bot) V11MessageEventHandler(eventData []byte, logger func(params ...any), handlers map[string]func(...any)) {
	var base onebot_v11_api_event.InternalMessageBase
	var privateMessageEvent onebot_v11_api_event.PrivateMessage
	var groupMessageEvent onebot_v11_api_event.GroupMessage

	err := json.Unmarshal(eventData, &base)
	if err != nil {
		logger(fmt.Sprintf("Unable to parse event data: %v, error: %v", eventData, err))
		return
	}

	logger("Handling message event")

	switch base.MessageType {
	case "private":
		err = json.Unmarshal(eventData, &privateMessageEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse as privateMessageEvent, data: %v, error: %v", eventData, err))
			return
		}
		RevokeHandlers("receive_private_message", handlers, bot, privateMessageEvent)
		break
	case "group":
		err = json.Unmarshal(eventData, &groupMessageEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse as groupMessageEvent, data: %v, error: %v", eventData, err))
			return
		}
		RevokeHandlers("receive_group_message", handlers, bot, groupMessageEvent)
		break
	default:
		logger(fmt.Sprintf("Unknown message type: %v", base.MessageType))
		return
	}
}

func (bot *V11Bot) V11NoticeEventHandler(eventData []byte, logger func(params ...any), handlers map[string]func(...any)) {

}

func (bot *V11Bot) V11RequestEventHandler(eventData []byte, logger func(params ...any), handlers map[string]func(...any)) {

}

func (bot *V11Bot) V11MetaEventHandler(eventData []byte, logger func(params ...any), handlers map[string]func(...any)) {

}

func (bot *V11Bot) ParseV11Event(eventData []byte, logger func(params ...any), handlers map[string]func(...any)) {
	var base onebot_v11_api_event.EventBase

	err := json.Unmarshal(eventData, &base)
	if err != nil {
		logger(fmt.Sprintf("Unable to parse event data: %v, error: %v", eventData, err))
		return
	}

	switch base.PostType {
	case "message":
		bot.V11MessageEventHandler(eventData, logger, handlers)
		break
	case "notice":
		bot.V11NoticeEventHandler(eventData, logger, handlers)
		break
	case "request":
		bot.V11RequestEventHandler(eventData, logger, handlers)
		break
	case "meta_event":
		bot.V11MetaEventHandler(eventData, logger, handlers)
		break
	default:
		logger(fmt.Sprintf("Unknown post_type [%v] from event backend", base.PostType))
		return
	}
}
