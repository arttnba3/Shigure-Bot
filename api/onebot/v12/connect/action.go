package onebot_v12_api_connect

type ActionRequest struct {
	Action string         `json:"action"`
	Params map[string]any `json:"params"`
	Echo   string         `json:"echo"`
}

type ActionRequestMultiConn struct {
	ActionRequest
	Self any `json:"self"`
}

type ActionResponse struct {
	Status  string      `json:"status"`
	Retcode int64       `json:"retcode"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Echo    string      `json:"echo"`
}

const (
	RETCODE_BAD_REQUEST              = 10001
	RETCODE_UNSUPPORTED_ACTION       = 10002
	RETCODE_BAD_PARAM                = 10003
	RETCODE_UNSUPPORTED_PARAM        = 10004
	RETCODE_UNSUPPORTED_SEGMENT      = 10005
	RETCODE_BAD_SEGMENT_DATA         = 10006
	RETCODE_UNSUPPORTED_SEGMENT_DATA = 10007
	RETCODE_WHO_AM_I                 = 10101
	RETCODE_UNKNOWN_SELF             = 10102
)

const (
	RETCODE_BAD_HANDLER            = 20001
	RETCODE_INTERNAL_HANDLER_ERROR = 20002
)

const (
	RETCODE_DATABASE_ERROR   = 31000
	RETCODE_FILESYSTEM_ERROR = 32000
	RETCODE_NETWORK_ERROR    = 33000
	RETCODE_PLATFORM_ERROR   = 34000
	RETCODE_LOGIC_ERROR      = 35000
	RETCODE_I_AM_TIRED       = 36000
)
