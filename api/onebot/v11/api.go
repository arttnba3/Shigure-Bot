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
	GetForwardMsg(ID string) (interface{}, error)
	SendLike(userID int64, times int64)
	SetGroupKick(groupID int64, userID int64, rejectAddRequest bool)
	SetGroupBan(groupID int64, userID int64, duration int64)
	SetGroupAnonymousBan(groupID int64, anonymous interface{}, anonymousFlag string, flag string, duration int64)
	SetGroupWholeBan(groupID int64, enable bool)
	SetGroupAdmin(groupID int64, userID int64, enable bool)
	SetGroupAnonymous(groupID int64, enable bool)
	SetGroupCard(groupID int64, userID int64, card string)
	SetGroupName(groupID int64, groupName string)
	SetGroupLeave(groupID int64, isDismiss bool)
	SetGroupSpecialTitle(groupID int64, userID int64, specialTitle string, duration int64)
	SetFriendAddRequest(flag string, approve bool, remark string)
	SetGroupAddRequest(flag string, Type string, subType string, approve bool, reason string)
	GetLoginInfo() (GetLoginInfoResp, error)
	GetStrangerInfo(userID int64, noCache bool) (GetStrangerInfoResp, error)
	GetFriendList() (GetFriendListResp, error)
	GetGroupInfo(groupID int64, noCache bool) (GetGroupInfoResp, error)
	GetGroupList() (GetGroupListResp, error)
	GetGroupMemberInfo(groupID int64, userID int64, noCache bool) (GetGroupMemberInfoResp, error)
	GetGroupMemberList(groupID int64) (GetGroupMemberListResp, error)
	GetGroupHonourInfo(groupID int64, Type string) (GetGroupHonourInfoResp, error)
	GetCookies(domain string) (GetCookiesResp, error)
	GetCSRFToken() (GetCSRFTokenResp, error)
	GetCredentials(domain string) (GetCredentialsResp, error)
	GetRecord(file string, outFormat string) (GetRecordResp, error)
	GetImage(file string) (GetImageResp, error)
	CanSendImage() (CanSendImageResp, error)
	CanSendRecord() (CanSendRecordResp, error)
	GetStatus() (GetStatusResp, error)
	GetVersionInfo() (GetVersionInfoResp, error)
	SetRestart()
	CleanCache()
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

type GetForwardMsgReq struct {
	ID string `json:"id"`
}

type GetForwardMsgResp struct {
	Message interface{} `json:"message"`
}

type SendLikeReq struct {
	UserID int64 `json:"user_id"`
	Times  int64 `json:"times"`
}

type SendLikeResp struct {
}

type SetGroupKickReq struct {
	GroupID          int64 `json:"group_id"`
	UserID           int64 `json:"user_id"`
	RejectAddRequest bool  `json:"reject_add_request"`
}

type SetGroupKickResp struct {
}

type SetGroupBanReq struct {
	GroupID  int64 `json:"group_id"`
	UserID   int64 `json:"user_id"`
	Duration int64 `json:"duration"`
}

type SetGroupBanResp struct {
}

type SetGroupAnonymousBanReq struct {
	GroupID       int64       `json:"group_id"`
	Anonymous     interface{} `json:"anonymous"`
	AnonymousFlag string      `json:"anonymous_flag"`
	Flag          string      `json:"flag"`
	Duration      int64       `json:"duration"`
}

type SetGroupAnonymousBanResp struct {
}

type SetGroupWholeBanReq struct {
	GroupID int64 `json:"group_id"`
	Enable  bool  `json:"enable"`
}

type SetGroupWholeBanResp struct {
}

type SetGroupAdminReq struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	Enable  bool  `json:"enable"`
}

type SetGroupAdminResp struct {
}

type SetGroupAnonymousReq struct {
	GroupID int64 `json:"group_id"`
	Enable  bool  `json:"enable"`
}

type SetGroupCardReq struct {
	GroupID int64  `json:"group_id"`
	UserID  int64  `json:"user_id"`
	Card    string `json:"card"`
}

type SetGroupCardResp struct {
}

type SetGroupNameReq struct {
	GroupID   int64  `json:"group_id"`
	GroupName string `json:"group_name"`
}

type SetGroupNameResp struct {
}

type SetGroupLeaveReq struct {
	GroupID   int64 `json:"group_id"`
	IsDismiss bool  `json:"is_dismiss"`
}

type SetGroupLeaveResp struct {
}

type SetGroupSpecialTitleReq struct {
	GroupID      int64  `json:"group_id"`
	UserID       int64  `json:"user_id"`
	SpecialTitle string `json:"special_title"`
	Duration     int64  `json:"duration"`
}

type SetGroupSpecialTitleResp struct{}

type SetFriendAddRequestReq struct {
	Flag    string `json:"flag"`
	Approve bool   `json:"approve"`
	Remark  string `json:"remark"`
}

type SetFriendAddRequestResp struct {
}

type SetGroupAddRequestReq struct {
	Flag    string `json:"flag"`
	Type    string `json:"type"`
	SubType string `json:"sub_type"`
	Approve bool   `json:"approve"`
	Reason  string `json:"reason"`
}

type SetGroupAddRequestResp struct {
}

type GetLoginInfoReq struct {
}

type GetLoginInfoResp struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
}

type GetStrangerInfoReq struct {
	UserID  int64 `json:"user_id"`
	NoCache bool  `json:"no_cache"`
}

type GetStrangerInfoResp struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}

type GetFriendListReq struct {
}

type GetFriendListRespItem struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
}

type GetFriendListResp = []GetFriendListRespItem

type GetGroupInfoReq struct {
	GroupID int64 `json:"group_id"`
	NoCache bool  `json:"no_cache"`
}

type GetGroupInfoResp struct {
	GroupID        int64  `json:"group_id"`
	GroupName      string `json:"group_name"`
	MemberCount    int32  `json:"member_count"`
	MaxMemberCount int32  `json:"max_member_count"`
}

type GetGroupListReq struct {
}

type GetGroupListResp = []GetGroupInfoResp

type GetGroupMemberInfoReq struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	NoCache bool  `json:"no_cache"`
}

type GetGroupMemberInfoResp struct {
	GroupID         int64  `json:"group_id"`
	UserID          int64  `json:"user_id"`
	Nickname        string `json:"nickname"`
	Card            string `json:"card"`
	Sex             string `json:"sex"`
	Age             int32  `json:"age"`
	Area            string `json:"area"`
	JoinTime        int32  `json:"join_time"`
	LastSentTime    int32  `json:"last_sent_time"`
	Level           string `json:"level"`
	Role            string `json:"role"`
	Unfriendly      bool   `json:"unfriendly"`
	Title           string `json:"title"`
	TitleExpireTime int32  `json:"title_expire_time"`
	CardChangeable  bool   `json:"card_changeable"`
}

type GetGroupMemberListReq struct {
	GroupID int64 `json:"group_id"`
}

type GetGroupMemberListResp = []GetGroupMemberInfoResp

type GetGroupHonourInfoReq struct {
	GroupID int64  `json:"group_id"`
	Type    string `json:"type"`
}

type GroupHonourCurrentTalkative struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	DayCount int32  `json:"day_count"`
}

type GroupHonourListItem struct {
	UserID      int64  `json:"user_id"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

type GetGroupHonourInfoResp struct {
	GroupID          int64                       `json:"group_id"`
	CurrentTalkative GroupHonourCurrentTalkative `json:"current_talkative"`
	TalkativeList    []GroupHonourListItem       `json:"talkative_list"`
	PerformerList    []GroupHonourListItem       `json:"performer_list"`
	LegendList       []GroupHonourListItem       `json:"legend_list"`
	StrongNewbieList []GroupHonourListItem       `json:"strong_newbie_list"`
	EmotionList      []GroupHonourListItem       `json:"emotion_list"`
}

type GetCookiesReq struct {
	Domain string `json:"domain"`
}

type GetCookiesResp struct {
	Cookies string `json:"cookies"`
}

type GetCSRFTokenReq struct {
}

type GetCSRFTokenResp struct {
	Token int32 `json:"token"`
}

type GetCredentialsReq struct {
	Domain string `json:"domain"`
}

type GetCredentialsResp struct {
	Cookies   string `json:"cookies"`
	CSRFToken int32  `json:"csrf_token"`
}

type GetRecordReq struct {
	File      string `json:"file"`
	OutFormat string `json:"out_format"`
}

type GetRecordResp struct {
	File string `json:"file"`
}

type GetImageReq struct {
	File string `json:"file"`
}

type GetImageResp struct {
	File string `json:"file"`
}

type CanSendImageReq struct {
}

type CanSendImageResp struct {
	Yes bool `json:"yes"`
}

type CanSendRecordReq struct {
}
type CanSendRecordResp struct {
	Yes bool `json:"yes"`
}

type GetStatusReq struct {
}

type GetStatusResp struct {
	Online bool `json:"online"`
	Good   bool `json:"good"`
}

type GetVersionInfoReq struct {
}

type GetVersionInfoResp struct {
	AppName         string `json:"app_name"`
	AppVersion      string `json:"app_version"`
	ProtocolVersion string `json:"protocol_version"`
}

type SetRestartReq struct {
	Delay int64 `json:"delay"`
}

type SetRestartResp struct {
}

type CleanCacheReq struct {
}

type CleanCacheResp struct {
}

// TODO: implement more APIs
