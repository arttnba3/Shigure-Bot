//
// Request Event
//
// Refer to following link for detailed docs:
//
// 		https://github.com/botuniverse/onebot-11/blob/master/event/request.md
//

package onebot_v11_api_event

type FriendAddRequest struct {
	Time        int64  `json:"time"`
	SelfId      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	RequestType string `json:"request_type"`
	UserId      int64  `json:"user_id"`
	Comment     string `json:"comment"`
	Flag        string `json:"flag"`
}

type GroupAddRequest struct {
	Time        int64  `json:"time"`
	SelfId      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	RequestType string `json:"request_type"`
	SubType     string `json:"sub_type"`
	GroupId     int64  `json:"group_id"`
	UserId      int64  `json:"user_id"`
	Comment     string `json:"comment"`
	Flag        string `json:"flag"`
}
