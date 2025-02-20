//
// Notice Event
//
// Refer to following link for detailed docs:
//
// 		https://github.com/botuniverse/onebot-11/blob/master/event/notice.md
//

package onebot_v11_api_event

// internal but not defined in specification

type NoticeEventBase struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
}

// Group file upload

type GroupFileInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Busid int64  `json:"busid"` // OneBot spec also doesn't know what it is
}

type GroupFileUpload struct {
	Time       int64         `json:"time"`
	SelfID     int64         `json:"self_id"`
	PostType   string        `json:"post_type"`
	NoticeType string        `json:"notice_type"`
	GroupID    int64         `json:"group_id"`
	UserID     int64         `json:"user_id"`
	File       GroupFileInfo `json:"file"`
}

// Group administrator change

type GroupAdminChange struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
}

// Group member removed

type GroupMemberRemoved struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	OperatorID int64  `json:"operator_id"`
	UserID     int64  `json:"user_id"`
}

// Group member added

type GroupMemberAdded struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	OperatorID int64  `json:"operator_id"`
	UserID     int64  `json:"user_id"`
}

// Group speak banned

type GroupSpeakBanned struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	OperatorID int64  `json:"operator_id"`
	UserID     int64  `json:"user_id"`
	Duration   int64  `json:"duration"`
}

// Friend Added

type FriendAdded struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	UserID     int64  `json:"user_id"`
}

// Group Message Recalled

type GroupMessageRecalled struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	OperatorID int64  `json:"operator_id"`
	MessageID  int64  `json:"message_id"`
}

// Friend Message Recalled

type FriendMessageRecalled struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	UserID     int64  `json:"user_id"`
	MessageID  int64  `json:"message_id"`
}

// Group chuoyichuo (how to translate this perfectly?)

type GroupChuoyichuo struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	TargetID   int64  `json:"target_id"`
}

// Group Red Packet King of Luck

type GroupRedPacketKingOfLuck struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	TargetID   int64  `json:"target_id"`
}

// Group member honours changed

type GroupMemberHonoursChanged struct {
	Time       int64  `json:"time"`
	SelfID     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	HonorType  string `json:"honor_type"`
	UserID     int64  `json:"user_id"`
}
