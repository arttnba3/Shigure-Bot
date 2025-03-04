package onebot_v11_impl

import (
	"encoding/json"
	"github.com/arttnba3/Shigure-Bot/api/onebot/v11"
	"github.com/arttnba3/Shigure-Bot/bot/onebot/v11/backend/http"
)

type V11Bot struct {
	Receiver onebot_v11_api.V11ReceiverAPI
	Sender   onebot_v11_api.V11SenderAPI
	Logger   func(params ...interface{})
	Handlers map[string]func(...any)
}

type V11HTTPServerConfig struct {
	Port int `json:"port"`
}

type V11HTTPPostConfig struct {
	Host string `json:host`
	Port int    `json:port`
}

type V11BotInfo struct {
	HTTPServer *V11HTTPServerConfig `json:"http_server,omitempty"`
	HTTPPost   *V11HTTPPostConfig   `json:"http_post,omitempty"`
}

func NewV11Bot(configJson []byte, logger func(params ...any), handlers map[string]func(...any)) (*V11Bot, error) {
	var receiver onebot_v11_api.V11ReceiverAPI = nil
	var sender onebot_v11_api.V11SenderAPI = nil
	var err error
	var botInfo V11BotInfo
	var bot V11Bot

	err = json.Unmarshal(configJson, &botInfo)
	if err != nil {
		return nil, err
	}

	if botInfo.HTTPServer != nil {
		receiver, err = onebot_v11_impl.NewV11HTTPReceiver(
			botInfo.HTTPServer.Port,
			logger,
			func(data []byte) {
				bot.ParseV11Event(data, logger, handlers)
			},
		)
		if err != nil {
			return nil, err
		}
	}

	if botInfo.HTTPPost != nil {
		sender, err = onebot_v11_impl.NewV11HTTPSender(botInfo.HTTPPost.Host, botInfo.HTTPPost.Port, logger)
		if err != nil {
			return nil, err
		}
	}

	bot.Receiver = receiver
	bot.Sender = sender
	bot.Logger = logger
	bot.Handlers = handlers

	return &bot, nil
}
