package shigure

import (
	"errors"
	"github.com/arttnba3/Shigure-Bot/bot/onebot/v11"
)

type ShigureBot struct {
	Bot      interface{}
	handlers map[string]func(...any)
}

func NewShigureBot(botType string, configJson []byte, logger func(...any), handler func(rawData []byte)) (*ShigureBot, error) {
	switch botType {
	case "OneBot-V11":
		bot, err := onebot_v11_impl.NewV11Bot(configJson, logger, handler) // TODO: switch from handler to handlers table
		if err != nil {
			return nil, err
		}

		return &ShigureBot{
			Bot:      bot,
			handlers: nil,
		}, nil
	default:
		return nil, errors.New("unknown bot type [" + botType + "]")
	}
}
