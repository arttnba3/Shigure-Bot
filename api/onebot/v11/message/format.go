//
// Message Format
//
// Note that type of some fields are not defined in docs, we have to guess they're all string here:(
//
// Refer to following link for detailed docs:
//
// 		https://github.com/botuniverse/onebot-11/tree/master/message
//

package onebot_v11_api_message

// Message types

type MessageString struct {
	UserID  int64  `json:"user_id"`
	Message string `json:"message"`
}

type MessageSegment struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type MessageArray = []MessageSegment

// message segment details

type MessageSegmentDataText struct {
	Text string `json:"text"`
}

type MessageSegmentDataEmoticon struct {
	ID string `json:"id"`
}

type MessageSegmentDataImage struct {
	File    string `json:"file"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Cache   string `json:"cache"`
	Proxy   string `json:"proxy"`
	Timeout string `json:"timeout"`
}

type MessageSegmentDataVoice struct {
	File    string `json:"file"`
	Magic   string `json:"magic"`
	URL     string `json:"url"`
	Cache   string `json:"cache"`
	Proxy   string `json:"proxy"`
	Timeout string `json:"timeout"`
}

type MessageSegmentDataVideo struct {
	File    string `json:"file"`
	URL     string `json:"url"`
	Cache   string `json:"cache"`
	Proxy   string `json:"proxy"`
	Timeout string `json:"timeout"`
}

type MessageSegmentDataAt struct {
	QQ string `json:"qq"`
}

type MessageSegmentDataRPS struct {
}

type MessageSegmentDataDice struct {
}

type MessageSegmentDataShake struct {
}

type MessageSegmentDataChuoyichuo struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MessageSegmentDataAnonymousMessage struct {
	Ignore string `json:"ignore"`
}

type MessageSegmentDataShareLink struct {
	URL     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

type MessageSegmentDataRecommendFriend struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type MessageSegmentDataRecommendGroup struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type MessageSegmentDataLocation struct {
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type MessageSegmentDataShareMusic struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type MessageSegmentDataShareCustomMusic struct {
	Type    string `json:"type"`
	URL     string `json:"url"`
	Audio   string `json:"audio"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

type MessageSegmentDataReply struct {
	ID string `json:"id"`
}

type MessageSegmentDataCombinedForward struct {
	ID string `json:"id"`
}

type MessageSegmentDataCombinedForwardNode struct {
	ID string `json:"id"`
}

type MessageSegmentDataCombinedForwardCustomNode struct {
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`
	Content  string `json:"content"`
}

type MessageSegmentDataXML struct {
	Data string `json:"data"`
}

type MessageSegmentDataJSON struct {
	Data string `json:"data"`
}
