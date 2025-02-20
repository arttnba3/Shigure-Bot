package onebot_v11_api

type V11SenderAPI interface {
	SendRequest(request BotAction) error
	GetRequestResult(uuid string) ([]byte, error)
	SendRequestAndGetResult(action string, params interface{}) ([]byte, error)
}
