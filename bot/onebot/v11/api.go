//
// Bot API
//
// Refer to following link for detailed docs:
//
// https://github.com/botuniverse/onebot-11/blob/master/api/public.md
//

package onebot_v11_impl

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/arttnba3/Shigure-Bot/api/onebot/v11"
)

func (bot *V11Bot) Log(params ...any) {
	if bot.Logger != nil {
		bot.Logger(params...)
	}
}

func (bot *V11Bot) SendPrivateMsg(userId int64, message interface{}, autoEscape bool) (int32, error) {
	params := onebot_v11_api.SendPrivateMsgReq{
		UserID:     userId,
		Message:    message,
		AutoEscape: autoEscape,
	}

	bot.Log(fmt.Sprintf("Send message to private [%v]: %v", userId, message))

	respJson, err := bot.Sender.SendRequestAndGetResult("send_private_msg", params)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while sending private message to user: %v, error: %v", userId, err.Error()))
		return -1, err
	}

	var respData onebot_v11_api.SendPrivateMsgResp
	if err := json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (send_private_msg)", respJson))
		return -1, errors.New("invalid response for (send_private_msg)")
	}

	return respData.MessageID, nil
}

func (bot *V11Bot) SendGroupMsg(groupId int64, message interface{}, autoEscape bool) (int32, error) {
	params := onebot_v11_api.SendGroupMsgReq{
		GroupID:    groupId,
		Message:    message,
		AutoEscape: autoEscape,
	}

	bot.Log(fmt.Sprintf("Send message to group [%v]: %v", groupId, message))

	respJson, err := bot.Sender.SendRequestAndGetResult("send_group_msg", params)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while sending msg to to group [%v], error: %v", groupId, err.Error()))
		return -1, err
	}

	var respData onebot_v11_api.SendGroupMsgResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (send_group_msg)", respJson))
		return -1, errors.New("invalid response for (send_group_msg)")
	}

	return respData.MessageID, nil
}

func (bot *V11Bot) SendMsg(messageType string, userId int64, groupId int64, message interface{}, autoEscape bool) (int32, error) {
	switch messageType {
	/*
		case "private":
			return bot.backend.SendPrivateMsg(userId, message, autoEscape)
		case "group":
			return bot.backend.SendGroupMsg(groupId, message, autoEscape)
	*/
	case "private":
		fallthrough
	case "group":
		reqParam := onebot_v11_api.SendMsgReq{
			MessageType: messageType,
			UserID:      userId,
			GroupID:     groupId,
			Message:     message,
			AutoEscape:  autoEscape,
		}

		bot.Log(fmt.Sprintf("Send message to %v [%v]: %v", messageType, groupId, message))

		respJson, err := bot.Sender.SendRequestAndGetResult("send_msg", reqParam)
		if err != nil {
			bot.Log(fmt.Sprintf("Error occur while sending message, original param: %v, error: %v", reqParam, err.Error()))
			return -1, err
		}

		var respData onebot_v11_api.SendMsgResp
		if err = json.Unmarshal(respJson, &respData); err != nil {
			bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (send_msg)", respJson))
			return -1, errors.New("invalid response for (send_msg)")
		}

		return respData.MessageID, nil

	default:
		bot.Log(
			"Invalid message type: %v, expect private or group",
		)
		return -1, errors.New(fmt.Sprintf("invalid message type: %v", messageType))
	}
}

func (bot *V11Bot) DeleteMsg(messageID int32) error {
	reqParam := onebot_v11_api.DeleteMsgReq{
		MessageID: messageID,
	}

	_, err := bot.Sender.SendRequestAndGetResult("delete_msg", reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while deleting message: %v, error: %v", messageID, err.Error()))
		return err
	}

	return nil
}

func (bot *V11Bot) GetMsg(messageID int32) (interface{}, error) {
	reqParam := onebot_v11_api.GetMsgReq{
		MessageID: messageID,
	}

	respJson, err := bot.Sender.SendRequestAndGetResult("get_msg", reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting message: %v, error: %v", messageID, err.Error()))
		return nil, err
	}

	var respData onebot_v11_api.GetMsgResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (get_msg)", respJson))
		return nil, errors.New("invalid response for (get_msg)")
	}

	return respData.Message, nil
}
