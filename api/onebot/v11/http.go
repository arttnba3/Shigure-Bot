//
// HTTP API
//
// Refer to following link for detailed docs:
//
// https://github.com/botuniverse/onebot-11/blob/master/communication/http.md
//

package onebot_v11_api

type HTTPResponseBody struct {
	Status  string      `json:"status"`
	RetCode int         `json:"retcode"`
	Data    interface{} `json:"data"`
}
