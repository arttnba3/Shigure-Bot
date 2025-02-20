//
// Message Event
//
// Refer to following link for detailed docs:
//
// 		https://github.com/botuniverse/onebot-11/blob/master/event/message.md
//

package onebot_v11_api

type PrivateMessageSender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}

type GroupMessageSender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Card     string `json:"card"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
	Area     string `json:"area"`
	Level    string `json:"level"`
	Role     string `json:"role"`
	Title    string `json:"title"`
}

type PrivateMessage struct {
	Time        int64                `json:"time"`
	SelfId      int64                `json:"self_id"`
	PostType    string               `json:"post_type"`
	MessageType string               `json:"message_type"`
	SubType     string               `json:"sub_type"`
	MessageId   int32                `json:"message_id"`
	UserId      int64                `json:"user_id"`
	Message     interface{}          `json:"message"`
	RawMessage  string               `json:"raw_message"`
	Font        int32                `json:"font"`
	Sender      PrivateMessageSender `json:"sender"`
}

type GroupMessage struct {
	Time        int64              `json:"time"`
	SelfId      int64              `json:"self_id"`
	PostType    string             `json:"post_type"`
	MessageType string             `json:"message_type"`
	SubType     string             `json:"sub_type"`
	MessageId   int32              `json:"message_id"`
	GroupId     int64              `json:"group_id"`
	UserId      int64              `json:"user_id"`
	Anonymous   interface{}        `json:"anonymous"`
	Message     interface{}        `json:"message"`
	RawMessage  string             `json:"raw_message"`
	Font        int32              `json:"font"`
	Sender      GroupMessageSender `json:"sender"`
}

type MessageSegment struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
