//
// Event base and sub-events
//
// Refer to following link for detailed docs:
//
//		https://github.com/botuniverse/onebot-11/blob/master/event/README.md
// 		https://github.com/botuniverse/onebot-11/blob/master/event/meta.md
//

package onebot_v11_api

// EventBase
// This struct ONLY indicates what an event should have
type EventBase struct {
	Time     int64  `json:"time"`
	SelfId   int64  `json:"self_id"`
	PostType string `json:"post_type"`
}

type LifeCycle struct {
	Time          int64  `json:"time"`
	SelfId        int64  `json:"self_id"`
	PostType      string `json:"post_type"`
	MetaEventType string `json:"meta_event_type"`
	SubType       string `json:"sub_type"`
}

type HeartBeat struct {
	Time          int64       `json:"time"`
	SelfId        int64       `json:"self_id"`
	PostType      string      `json:"post_type"`
	MetaEventType string      `json:"meta_event_type"`
	Status        interface{} `json:"status"`
	Interval      int64       `json:"interval"`
}
