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
	reqParam := onebot_v11_api.SendPrivateMsgReq{
		UserID:     userId,
		Message:    message,
		AutoEscape: autoEscape,
	}
	reqAction := "send_private_msg"

	bot.Log(fmt.Sprintf("Send message to private [%v]: %v", userId, message))

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while sending private message to user: %v, error: %v", userId, err.Error()))
		return -1, err
	}

	var respData onebot_v11_api.SendPrivateMsgResp
	if err := json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", respJson, reqAction))
		return -1, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData.MessageID, nil
}

func (bot *V11Bot) SendGroupMsg(groupId int64, message interface{}, autoEscape bool) (int32, error) {
	reqParam := onebot_v11_api.SendGroupMsgReq{
		GroupID:    groupId,
		Message:    message,
		AutoEscape: autoEscape,
	}
	reqAction := "send_group_msg"

	bot.Log(fmt.Sprintf("Send message to group [%v]: %v", groupId, message))

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while sending msg to to group [%v], error: %v", groupId, err.Error()))
		return -1, err
	}

	var respData onebot_v11_api.SendGroupMsgResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", respJson, reqAction))
		return -1, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
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
		reqAction := "send_msg"

		bot.Log(fmt.Sprintf("Send message to %v [%v]: %v", messageType, groupId, message))

		respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
		if err != nil {
			bot.Log(fmt.Sprintf("Error occur while sending message, original param: %v, error: %v", reqParam, err.Error()))
			return -1, err
		}

		var respData onebot_v11_api.SendMsgResp
		if err = json.Unmarshal(respJson, &respData); err != nil {
			bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", respJson, reqAction))
			return -1, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
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
	reqAction := "delete_msg"

	_, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
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
	reqAction := "get_msg"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting message: %v, error: %v", messageID, err.Error()))
		return nil, err
	}

	var respData onebot_v11_api.GetMsgResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", respJson, reqAction))
		return nil, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData.Message, nil
}

func (bot *V11Bot) GetForwardMsg(ID string) (interface{}, error) {
	reqParam := onebot_v11_api.GetForwardMsgReq{
		ID: ID,
	}
	reqAction := "get_forward_msg"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting forwaed message: %v, error: %v", ID, err.Error()))
		return nil, err
	}

	var respData onebot_v11_api.GetForwardMsgResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", respJson, reqAction))
		return nil, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData.Message, nil
}

func (bot *V11Bot) SendLike(userID int64, times int64) {
	reqParam := onebot_v11_api.SendLikeReq{
		UserID: userID,
		Times:  times,
	}
	reqAction := "send_like"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while sending like to user: %v, error: %v", userID, err.Error()))
		return
	}

	var respData onebot_v11_api.SendLikeResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", respJson, reqAction))
		return
	}

	return
}

func (bot *V11Bot) SetGroupKick(groupID int64, userID int64, rejectAddRequest bool) {
	reqParam := onebot_v11_api.SetGroupKickReq{
		GroupID:          groupID,
		UserID:           userID,
		RejectAddRequest: rejectAddRequest,
	}
	reqAction := "set_group_kick"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while setting kicking %v in group: %v, error: %v", userID, groupID, err.Error()))
		return
	}

	var respData onebot_v11_api.SetGroupKickResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", respJson, reqAction))
		return
	}

	return
}

func (bot *V11Bot) SetGroupBan(groupID int64, userID int64, duration int64) {
	reqParam := onebot_v11_api.SetGroupBanReq{
		GroupID:  groupID,
		UserID:   userID,
		Duration: duration,
	}
	reqAction := "set_group_ban"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while setting banning %v in group: %v, error: %v", userID, groupID, err.Error()))
		return
	}

	var respData onebot_v11_api.SetGroupBanResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", respJson, reqAction))
		return
	}

	return
}

func (bot *V11Bot) SetGroupAnonymousBan(groupID int64, anonymous interface{}, anonymousFlag string, flag string, duration int64) {
	reqParam := onebot_v11_api.SetGroupAnonymousBanReq{
		GroupID:       groupID,
		Anonymous:     anonymous,
		AnonymousFlag: anonymousFlag,
		Flag:          flag,
		Duration:      duration,
	}
	reqAction := "set_group_anonymous_ban"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while setting banning %v in group: %v, error: %v", anonymous, groupID, err.Error()))
		return
	}

	var respData onebot_v11_api.SetGroupAnonymousBanResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return
	}

	return
}

func (bot *V11Bot) SetGroupWholeBan(groupID int64, enable bool) {
	reqParam := onebot_v11_api.SetGroupWholeBanReq{
		GroupID: groupID,
		Enable:  enable,
	}
	reqAction := "set_group_whole_ban"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while setting whole ban in group: %v, error: %v", groupID, err.Error()))
		return
	}

	var respData onebot_v11_api.SetGroupWholeBanResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return
	}

	return
}

func (bot *V11Bot) SetGroupAdmin(groupID int64, userID int64, enable bool) {
	reqParam := onebot_v11_api.SetGroupAdminReq{
		GroupID: groupID,
		UserID:  userID,
		Enable:  enable,
	}
	reqAction := "set_group_admin"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while setting admin %v in group: %v, error: %v", userID, groupID, err.Error()))
		return
	}

	var respData onebot_v11_api.SetGroupAdminResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return
	}

	return
}

func (bot *V11Bot) SetGroupAnonymous(groupID int64, enable bool) {
	reqParam := onebot_v11_api.SetGroupAnonymousReq{
		GroupID: groupID,
		Enable:  enable,
	}
	reqAction := "set_group_anonymous"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while setting group %v anonymous, error: %v", groupID, err.Error()))
		return
	}

	var respData onebot_v11_api.SetGroupAdminResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return
	}

	return
}

func (bot *V11Bot) SetGroupCard(groupID int64, userID int64, card string) {
	reqParam := onebot_v11_api.SetGroupCardReq{
		GroupID: groupID,
		UserID:  userID,
		Card:    card,
	}
	reqAction := "set_group_card"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while setting card for group %v, error: %v", groupID, err.Error()))
		return
	}

	var respData onebot_v11_api.SetGroupCardResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return
	}

	return
}

func (bot *V11Bot) SetGroupName(groupID int64, groupName string) {
	reqParam := onebot_v11_api.SetGroupNameReq{
		GroupID:   groupID,
		GroupName: groupName,
	}
	reqAction := "set_group_name"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while setting name for group %v, error: %v", groupID, err.Error()))
		return
	}

	var respData onebot_v11_api.SetGroupNameResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return
	}

	return
}

func (bot *V11Bot) SetGroupLeave(groupID int64, isDismiss bool) {
	reqParam := onebot_v11_api.SetGroupLeaveReq{
		GroupID:   groupID,
		IsDismiss: isDismiss,
	}
	reqAction := "set_group_leave"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while setting leave for group %v, error: %v", groupID, err.Error()))
		return
	}

	var respData onebot_v11_api.SetGroupNameResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return
	}

	return
}

func (bot *V11Bot) SetGroupSpecialTitle(groupID int64, userID int64, specialTitle string, duration int64) {
	reqParam := onebot_v11_api.SetGroupSpecialTitleReq{
		GroupID:      groupID,
		UserID:       userID,
		SpecialTitle: specialTitle,
		Duration:     duration,
	}
	reqAction := "set_group_special_title"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while setting special title for %v in group %v, error: %v", userID, groupID, err.Error()))
		return
	}

	var respData onebot_v11_api.SetGroupSpecialTitleResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return
	}

	return
}

func (bot *V11Bot) SetFriendAddRequest(flag string, approve bool, remark string) {
	reqParam := onebot_v11_api.SetFriendAddRequestReq{
		Flag:    flag,
		Approve: approve,
		Remark:  remark,
	}
	reqAction := "set_friend_add_request"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while setting friend add request, error: %v", err.Error()))
		return
	}

	var respData onebot_v11_api.SetFriendAddRequestResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return
	}

	return
}

func (bot *V11Bot) SetGroupAddRequest(flag string, Type string, subType string, approve bool, reason string) {
	reqParam := onebot_v11_api.SetGroupAddRequestReq{
		Flag:    flag,
		Type:    Type,
		SubType: subType,
		Approve: approve,
		Reason:  reason,
	}
	reqAction := "set_group_add_request"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while setting group add request, error: %v", err.Error()))
		return
	}

	var respData onebot_v11_api.SetGroupAddRequestResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return
	}

	return
}

func (bot *V11Bot) GetLoginInfo() (onebot_v11_api.GetLoginInfoResp, error) {
	reqParam := onebot_v11_api.GetLoginInfoReq{}
	reqAction := "get_login_info"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting login info, error: %v", err.Error()))
		return onebot_v11_api.GetLoginInfoResp{}, err
	}

	var respData onebot_v11_api.GetLoginInfoResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.GetLoginInfoResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) GetStrangerInfo(userID int64, noCache bool) (onebot_v11_api.GetStrangerInfoResp, error) {
	reqParam := onebot_v11_api.GetStrangerInfoReq{
		UserID:  userID,
		NoCache: noCache,
	}
	reqAction := "get_stranger_info"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting stranger info, error: %v", err.Error()))
		return onebot_v11_api.GetStrangerInfoResp{}, err
	}

	var respData onebot_v11_api.GetStrangerInfoResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.GetStrangerInfoResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) GetFriendList() (onebot_v11_api.GetFriendListResp, error) {
	reqParam := onebot_v11_api.GetFriendListReq{}
	reqAction := "get_friend_list"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting friend list, error: %v", err.Error()))
		return onebot_v11_api.GetFriendListResp{}, err
	}

	var respData onebot_v11_api.GetFriendListResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.GetFriendListResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) GetGroupInfo(groupID int64, noCache bool) (onebot_v11_api.GetGroupInfoResp, error) {
	reqParam := onebot_v11_api.GetGroupInfoReq{
		GroupID: groupID,
		NoCache: noCache,
	}
	reqAction := "get_group_info"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting info of group %v, error: %v", groupID, err.Error()))
		return onebot_v11_api.GetGroupInfoResp{}, err
	}

	var respData onebot_v11_api.GetGroupInfoResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.GetGroupInfoResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) GetGroupList() (onebot_v11_api.GetGroupListResp, error) {
	reqParam := onebot_v11_api.GetGroupListReq{}
	reqAction := "get_group_list"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting group list, error: %v", err.Error()))
		return onebot_v11_api.GetGroupListResp{}, err
	}

	var respData onebot_v11_api.GetGroupListResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.GetGroupListResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) GetGroupMemberInfo(groupID int64, userID int64, noCache bool) (onebot_v11_api.GetGroupMemberInfoResp, error) {
	reqParam := onebot_v11_api.GetGroupMemberInfoReq{
		GroupID: groupID,
		UserID:  userID,
		NoCache: noCache,
	}
	reqAction := "get_group_member_info"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting member %v info in group %v, error: %v", userID, groupID, err.Error()))
		return onebot_v11_api.GetGroupMemberInfoResp{}, err
	}

	var respData onebot_v11_api.GetGroupMemberInfoResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.GetGroupMemberInfoResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) GetGroupMemberList(groupID int64) (onebot_v11_api.GetGroupMemberListResp, error) {
	reqParam := onebot_v11_api.GetGroupMemberListReq{
		GroupID: groupID,
	}
	reqAction := "get_group_member_list"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting member list of group %v, error: %v", groupID, err.Error()))
		return onebot_v11_api.GetGroupMemberListResp{}, err
	}

	var respData onebot_v11_api.GetGroupMemberListResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.GetGroupMemberListResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) GetGroupHonourInfo(groupID int64, Type string) (onebot_v11_api.GetGroupHonourInfoResp, error) {
	reqParam := onebot_v11_api.GetGroupHonourInfoReq{
		GroupID: groupID,
		Type:    Type,
	}
	reqAction := "get_group_honour_info"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting honour info of group %v, error: %v", groupID, err.Error()))
		return onebot_v11_api.GetGroupHonourInfoResp{}, err
	}

	var respData onebot_v11_api.GetGroupHonourInfoResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.GetGroupHonourInfoResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) GetCookies(domain string) (onebot_v11_api.GetCookiesResp, error) {
	reqParam := onebot_v11_api.GetCookiesReq{
		Domain: domain,
	}
	reqAction := "get_cookies"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting cookies of domain %v, error: %v", domain, err.Error()))
		return onebot_v11_api.GetCookiesResp{}, err
	}

	var respData onebot_v11_api.GetCookiesResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.GetCookiesResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) GetCSRFToken() (onebot_v11_api.GetCSRFTokenResp, error) {
	reqParam := onebot_v11_api.GetCSRFTokenReq{}
	reqAction := "get_csrf_token"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting CSRF token, error: %v", err.Error()))
		return onebot_v11_api.GetCSRFTokenResp{}, err
	}

	var respData onebot_v11_api.GetCSRFTokenResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.GetCSRFTokenResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) GetCredentials(domain string) (onebot_v11_api.GetCredentialsResp, error) {
	reqParam := onebot_v11_api.GetCredentialsReq{
		Domain: domain,
	}
	reqAction := "get_credentials"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting credentials of domain %v, error: %v", domain, err.Error()))
		return onebot_v11_api.GetCredentialsResp{}, err
	}

	var respData onebot_v11_api.GetCredentialsResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.GetCredentialsResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) GetRecord(file string, outFormat string) (onebot_v11_api.GetRecordResp, error) {
	reqParam := onebot_v11_api.GetRecordReq{
		File:      file,
		OutFormat: outFormat,
	}
	reqAction := "get_record"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while transforming format of file %v, error: %v", file, err.Error()))
		return onebot_v11_api.GetRecordResp{}, err
	}

	var respData onebot_v11_api.GetRecordResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log("Invalid response data: %v, expect resp format for (%v)", err, reqAction)
		return onebot_v11_api.GetRecordResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) GetImage(file string) (onebot_v11_api.GetImageResp, error) {
	reqParam := onebot_v11_api.GetImageReq{
		File: file,
	}
	reqAction := "get_image"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting image file %v, error: %v", file, err.Error()))
		return onebot_v11_api.GetImageResp{}, err
	}

	var respData onebot_v11_api.GetImageResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.GetImageResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) CanSendImage() (onebot_v11_api.CanSendImageResp, error) {
	reqParam := onebot_v11_api.CanSendImageReq{}
	reqAction := "can_send_image"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while checking %v, error: %v", reqAction, err.Error()))
		return onebot_v11_api.CanSendImageResp{}, err
	}

	var respData onebot_v11_api.CanSendImageResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.CanSendImageResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) CanSendRecord() (onebot_v11_api.CanSendRecordResp, error) {
	reqParam := onebot_v11_api.CanSendRecordReq{}
	reqAction := "can_send_record"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while checking %v, error: %v", reqAction, err.Error()))
		return onebot_v11_api.CanSendRecordResp{}, err
	}

	var respData onebot_v11_api.CanSendRecordResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.CanSendRecordResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) GetStatus() (onebot_v11_api.GetStatusResp, error) {
	reqParam := onebot_v11_api.GetStatusReq{}
	reqAction := "get_status"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting status, error: %v", err.Error()))
		return onebot_v11_api.GetStatusResp{}, err
	}

	var respData onebot_v11_api.GetStatusResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.GetStatusResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) GetVersionInfo() (onebot_v11_api.GetVersionInfoResp, error) {
	reqParam := onebot_v11_api.GetVersionInfoReq{}
	reqAction := "get_version_info"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while getting version info, error: %v", err.Error()))
		return onebot_v11_api.GetVersionInfoResp{}, err
	}

	var respData onebot_v11_api.GetVersionInfoResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return onebot_v11_api.GetVersionInfoResp{}, errors.New(fmt.Sprintf("invalid response for (%v)", reqAction))
	}

	return respData, nil
}

func (bot *V11Bot) SetRestart() {
	reqParam := onebot_v11_api.SetRestartReq{}
	reqAction := "set_restart"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while setting restart, error: %v", err.Error()))
		return
	}

	var respData onebot_v11_api.SetRestartResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return
	}

	return
}

func (bot *V11Bot) CleanCache() {
	reqParam := onebot_v11_api.CleanCacheReq{}
	reqAction := "clean_cache"

	respJson, err := bot.Sender.SendRequestAndGetResult(reqAction, reqParam)
	if err != nil {
		bot.Log(fmt.Sprintf("Error occur while cleaning cache, error: %v", err.Error()))
		return
	}

	var respData onebot_v11_api.CleanCacheResp
	if err = json.Unmarshal(respJson, &respData); err != nil {
		bot.Log(fmt.Sprintf("Invalid response data: %v, expect resp format for (%v)", err, reqAction))
		return
	}

	return
}
