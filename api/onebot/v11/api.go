//
// OneBot V11 API
//
// Refer to following link for detailed docs:
//
// 		https://github.com/botuniverse/onebot-11/blob/master/api/public.md
//		https://github.com/botuniverse/onebot-11/blob/master/api/hidden.md
//

package onebot_v11_api

type V11BotAPI interface {
	SendPrivateMsg(userId int64, message interface{}, autoEscape bool) (int32, error)
	SendGroupMsg(groupId int64, message interface{}, autoEscape bool) (int32, error)
	SendMsg(messageType string, userId int64, groupId int64, message interface{}, autoEscape bool) (int32, error)
	DeleteMsg(messageID int32) error
	GetMsg(messageID int32) (interface{}, error)
	// TODO: implement more API
}

type BotAction struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	UUID   string      `json:"uuid"`
}

// send_private_msg

type SendPrivateMsgReq struct {
	UserID     int64       `json:"user_id"`
	Message    interface{} `json:"message"`
	AutoEscape bool        `json:"auto_escape"`
}

type SendPrivateMsgResp struct {
	MessageID int32 `json:"message_id"`
}

// send_group_msg

type SendGroupMsgReq struct {
	GroupID    int64       `json:"group_id"`
	Message    interface{} `json:"message"`
	AutoEscape bool        `json:"auto_escape"`
}

type SendGroupMsgResp struct {
	MessageID int32 `json:"message_id"`
}

type SendMsgReq struct {
	MessageType string      `json:"message_type"`
	UserID      int64       `json:"user_id"`
	GroupID     int64       `json:"group_id"`
	Message     interface{} `json:"message"`
	AutoEscape  bool        `json:"auto_escape"`
}

type SendMsgResp struct {
	MessageID int32 `json:"message_id"`
}

type DeleteMsgReq struct {
	MessageID int32 `json:"message_id"`
}

// DeleteMsgResp : Only a placeholder here
type DeleteMsgResp struct {
}

type GetMsgReq struct {
	MessageID int32 `json:"message_id"`
}
type GetMsgResp struct {
	Time        int32       `json:"time"`
	MessageType string      `json:"message_type"`
	MessageID   int32       `json:"message_id"`
	RealID      int32       `json:"real_id"`
	Sender      interface{} `json:"sender"`
	Message     interface{} `json:"message"`
}

// TODO: implement more APIs
