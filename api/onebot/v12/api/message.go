package onebot_v12_api_interface

type MessageSegment struct {
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}

type MessageSegmentDataText struct {
	Text string `json:"text"`
}

type MessageSegmentDataMention struct {
	UserID string `json:"user_id"`
}

type MessageSegmentDataMentionAll struct {
}

type MessageSegmentDataImage struct {
	FileID string `json:"file_id"`
}

type MessageSegmentDataVoice struct {
	FileID string `json:"file_id"`
}

type MessageSegmentDataAudio struct {
	FileID string `json:"file_id"`
}

type MessageSegmentDaVideo struct {
	FileID string `json:"file_id"`
}

type MessageSegmentDataFile struct {
	FileID string `json:"file_id"`
}

type MessageSegmentDataLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Title     string  `json:"title"`
	Content   string  `json:"content"`
}

type MessageSegmentDataReply struct {
	MessageID string `json:"message_id"`
	UserID    string `json:"user_id"`
}
