package onebot_v12_api_interface

import "github.com/arttnba3/Shigure-Bot/api/onebot/v12/connect"

type MetaConnect struct {
	onebot_v12_api_connect.EventBase
	Version onebot_v12_api_connect.ActionResponse `json:"version"`
}

type MetaHeartbeat struct {
	onebot_v12_api_connect.EventBase
	Interval int64 `json:"interval"`
}

type MetaStatusUpdate struct {
	onebot_v12_api_connect.EventBase
	Status onebot_v12_api_connect.ActionResponse `json:"status"`
}
