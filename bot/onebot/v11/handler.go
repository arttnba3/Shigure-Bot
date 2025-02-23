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

	switch base.MessageType {
	case "private":
		err = json.Unmarshal(eventData, &privateMessageEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse as privateMessageEvent, data: %v, error: %v", eventData, err))
			return
		}
		RevokeHandlers("message_private", handlers, bot, privateMessageEvent)
		break
	case "group":
		err = json.Unmarshal(eventData, &groupMessageEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse as groupMessageEvent, data: %v, error: %v", eventData, err))
			return
		}
		RevokeHandlers("message_group", handlers, bot, groupMessageEvent)
		break
	default:
		logger(fmt.Sprintf("Unknown message type: %v", base.MessageType))
		return
	}
}

func (bot *V11Bot) V11NoticeEventHandler(eventData []byte, logger func(params ...any), handlers map[string]func(...any)) {
	var noticeBase onebot_v11_api_event.NoticeEventBase

	err := json.Unmarshal(eventData, &noticeBase)
	if err != nil {
		logger(fmt.Sprintf("Unable to parse event data: %v, error: %v", eventData, err))
		return
	}

	switch noticeBase.NoticeType {
	case "group_upload":
		var groupFileUploadEvent onebot_v11_api_event.GroupFileUpload
		err = json.Unmarshal(eventData, &groupFileUploadEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse as groupFileUploadEvent, data: %v, error: %v", eventData, err))
			return
		}
		RevokeHandlers("notice_group_upload", handlers, bot, groupFileUploadEvent)
		break
	case "group_admin":
		var groupAdminChangedEvent onebot_v11_api_event.GroupAdminChange
		err = json.Unmarshal(eventData, &groupAdminChangedEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse as groupAdminChangedEvent, data: %v, error: %v", eventData, err))
			return
		}
		RevokeHandlers("notice_group_admin", handlers, bot, groupAdminChangedEvent)
		break
	case "group_decrease":
		var groupMemberRemovedEvent onebot_v11_api_event.GroupMemberRemoved
		err = json.Unmarshal(eventData, &groupMemberRemovedEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse as groupMemberRemovedEvent, data: %v, error: %v", eventData, err))
			return
		}
		RevokeHandlers("notice_group_decrease", handlers, bot, groupMemberRemovedEvent)
		break
	case "group_increase":
		var groupMemberAddedEvent onebot_v11_api_event.GroupMemberAdded
		err = json.Unmarshal(eventData, &groupMemberAddedEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse as groupMemberAddedEvent, data: %v, error: %v", eventData, err))
			return
		}
		RevokeHandlers("notice_group_increase", handlers, bot, groupMemberAddedEvent)
		break
	case "group_ban":
		var groupSpeakBannedEvent onebot_v11_api_event.GroupSpeakBanned
		err = json.Unmarshal(eventData, &groupSpeakBannedEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse as groupSpeakBannedEvent, data: %v, error: %v", eventData, err))
			return
		}
		RevokeHandlers("notice_group_ban", handlers, bot, groupSpeakBannedEvent)
		break
	case "friend_add":
		var friendAddedEvent onebot_v11_api_event.FriendAdded
		err = json.Unmarshal(eventData, &friendAddedEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse as friendAddedEvent, data: %v, error: %v", eventData, err))
			return
		}
		RevokeHandlers("notice_friend_add", handlers, bot, friendAddedEvent)
		break
	case "group_recall":
		var groupMessageRecalledEvent onebot_v11_api_event.GroupMessageRecalled
		err = json.Unmarshal(eventData, &groupMessageRecalledEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse as groupMessageRecalledEvent, data: %v, error: %v", eventData, err))
			return
		}
		RevokeHandlers("notice_group_recall", handlers, bot, groupMessageRecalledEvent)
		break
	case "friend_recall":
		var friendMessageRecalledEvent onebot_v11_api_event.FriendMessageRecalled
		err = json.Unmarshal(eventData, &friendMessageRecalledEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse as friendMessageRecalledEvent, data: %v, error: %v", eventData, err))
			return
		}
		RevokeHandlers("notice_friend_recall", handlers, bot, friendMessageRecalledEvent)
		break
	case "notify":
		switch noticeBase.SubType {
		case "poke":
			var groupChuoyichuoEvent onebot_v11_api_event.GroupChuoyichuo
			err = json.Unmarshal(eventData, &groupChuoyichuoEvent)
			if err != nil {
				logger(fmt.Sprintf("Unable to parse as groupChuoyichuoEvent, data: %v, error: %v", eventData, err))
				return
			}
			RevokeHandlers("notice_notify_poke", handlers, bot, groupChuoyichuoEvent)
			break
		case "lucky_king":
			var groupRedPacketKingOfLuckEvent onebot_v11_api_event.GroupRedPacketKingOfLuck
			err = json.Unmarshal(eventData, &groupRedPacketKingOfLuckEvent)
			if err != nil {
				logger(fmt.Sprintf("Unable to parse as groupRedPacketKingOfLuckEvent, data: %v, error: %v", eventData, err))
				return
			}
			RevokeHandlers("notice_notify_lucky_king", handlers, bot, groupRedPacketKingOfLuckEvent)
			break
		case "honor":
			var groupMemberHonoursChangedEvent onebot_v11_api_event.GroupMemberHonoursChanged
			err = json.Unmarshal(eventData, &groupMemberHonoursChangedEvent)
			if err != nil {
				logger(fmt.Sprintf("Unable to parse as groupMemberHonoursChangedEvent, data: %v, error: %v", eventData, err))
				return
			}
			RevokeHandlers("notice_notify_honor", handlers, bot, groupMemberHonoursChangedEvent)
			break
		default:
			logger(fmt.Sprintf("Unknown notice sub_typee: %v", noticeBase.SubType))
			return
		}
		break
	default:
		logger(fmt.Sprintf("Unknown notice_type: %v", noticeBase.NoticeType))
		return
	}
}

func (bot *V11Bot) V11RequestEventHandler(eventData []byte, logger func(params ...any), handlers map[string]func(...any)) {
	var requestEventBase onebot_v11_api_event.RequestEventBase

	err := json.Unmarshal(eventData, &requestEventBase)
	if err != nil {
		logger(fmt.Sprintf("Unable to parse event data: %v, error: %v", eventData, err))
		return
	}

	switch requestEventBase.RequestType {
	case "friend":
		var friendAddedRequestEvent onebot_v11_api_event.FriendAddRequest
		err = json.Unmarshal(eventData, &friendAddedRequestEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse as friendAddedRequestEvent, data: %v, error: %v", eventData, err))
			return
		}
		RevokeHandlers("request_friend", handlers, bot, friendAddedRequestEvent)
		break
	case "group":
		var groupAddedRequestEvent onebot_v11_api_event.GroupAddRequest
		err = json.Unmarshal(eventData, &groupAddedRequestEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse as groupAddedRequestEvent, data: %v, error: %v", eventData, err))
			return
		}
		RevokeHandlers("request_group", handlers, bot, groupAddedRequestEvent)
		break
	}
}

func (bot *V11Bot) V11MetaEventHandler(eventData []byte, logger func(params ...any), handlers map[string]func(...any)) {
	var metaEventBase onebot_v11_api_event.MetaEventBase

	err := json.Unmarshal(eventData, &metaEventBase)
	if err != nil {
		logger(fmt.Sprintf("Unable to parse event data: %v, error: %v", eventData, err))
		return
	}

	switch metaEventBase.MetaEventType {
	case "lifecycle":
		var lifeCycleEvent onebot_v11_api_event.LifeCycle

		err = json.Unmarshal(eventData, &lifeCycleEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse event data: %v, error: %v", eventData, err))
			return
		}

		RevokeHandlers("request_lifecycle", handlers, bot, lifeCycleEvent)
		break
	case "heartbeat":
		var heartBeatEvent onebot_v11_api_event.HeartBeat

		err = json.Unmarshal(eventData, &heartBeatEvent)
		if err != nil {
			logger(fmt.Sprintf("Unable to parse event data: %v, error: %v", eventData, err))
			return
		}

		RevokeHandlers("request_heartbeat", handlers, bot, heartBeatEvent)
		break
	default:
		logger(fmt.Sprintf("Unknown meta_event_type: %v", metaEventBase.MetaEventType))
		return
	}
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
