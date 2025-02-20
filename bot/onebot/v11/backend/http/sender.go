package onebot_v11_impl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/arttnba3/Shigure-Bot/api/onebot/v11"
	"io"
	"net/http"
	"sync"
	"time"
)

type V11HTTPSender struct {
	postHost string
	postPort int
	Logger   func(params ...interface{})
	respDB   sync.Map // safe for race condition
	UUIDBase int64
	uuidLock sync.Mutex
}

func (sender *V11HTTPSender) Log(format string, params ...any) {
	if sender.Logger != nil {
		sender.Logger(fmt.Sprintf(format, params))
	}
}

func (sender *V11HTTPSender) NextUUID() string {
	sender.uuidLock.Lock()
	defer sender.uuidLock.Unlock()

	uuid := fmt.Sprintf("%v-%v", sender.UUIDBase, time.Now().UnixNano())
	sender.UUIDBase += 1

	return uuid
}

func (sender *V11HTTPSender) SendRequest(request onebot_v11_api.BotAction) error {
	postUrl := fmt.Sprintf("http://%v:%v/%v", sender.postHost, sender.postPort, request.Action)
	byteData, err := json.Marshal(request.Params)
	if err != nil {
		sender.Log("Error occur while sending request to url: \"%v\", error: %v", postUrl, err.Error())
		return err
	}

	resp, err := http.Post(postUrl, "application/json;charset=UTF-8", bytes.NewReader(byteData))
	if err != nil {
		sender.Log("Error occur while getting response from url: \"%v\", error: %v", postUrl, err.Error())
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			sender.Log("Error occur while closing response body: \"%v\", error: %v", postUrl, err.Error())
		}
	}(resp.Body)

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		sender.Log("Error occur while reading response from url: \"%v\", error: %v", postUrl, err.Error())
		return err
	}

	sender.respDB.Store(request.UUID, respData)
	return nil
}

func (sender *V11HTTPSender) GetRequestResult(uuid string) ([]byte, error) {
	data, ok := sender.respDB.LoadAndDelete(uuid) // We only need to get result for each action ONCE
	if ok {
		return data.([]byte), nil
	} else {
		return nil, errors.New("invalid uuid")
	}
}

func (sender *V11HTTPSender) SendRequestAndGetResult(action string, params interface{}) ([]byte, error) {
	var respBody onebot_v11_api.HTTPResponseBody

	uuid := sender.NextUUID()
	request := onebot_v11_api.BotAction{
		Action: action,
		Params: params,
		UUID:   uuid,
	}

	err := sender.SendRequest(request)
	if err != nil {
		sender.Log(fmt.Sprintf(
			"Error occur while sending request for msg: %v, error: %v", request.UUID, err.Error(),
		))
		return nil, err
	}

	reqResp, err := sender.GetRequestResult(request.UUID)
	if err != nil {
		sender.Log(fmt.Sprintf(
			"Error occur while getting response for msg: %v, error: %v", request.UUID, err.Error(),
		))
		return nil, err
	}

	err = json.Unmarshal(reqResp, &respBody)
	if err != nil {
		sender.Log(fmt.Sprintf(
			"Cannot parse response for msg: %v, original response: %v, error: %v", request.UUID, reqResp, err.Error(),
		))
		return nil, err
	}

	respJson, err := json.Marshal(respBody.Data)
	if err != nil {
		sender.Log(fmt.Sprintf("Cannot marshal resp data: %v, error: %v", respBody.Data, err.Error()))
		return nil, err
	}

	switch respBody.Status {
	case "ok":
		return respJson, nil
	case "async":
		sender.Log(fmt.Sprintf("Got unsupported async result from OneBot backend for msg: %v", request.UUID))
		return nil, errors.New("unsupported async result from OneBot backend")
	case "failed":
		sender.Log(fmt.Sprintf(
			"Failed to get result from OneBot backend for msg: %v, retcode: %v",
			request.UUID,
			respBody.RetCode,
		))
		return nil, errors.New("failed to get result from OneBot backend")
	default:
		sender.Log(fmt.Sprintf(
			"Unknown status \"%v\" from OneBot backend for msg: %v", respBody.Status, request.UUID,
		))
		return nil, errors.New("unknown status from OneBot backend")
	}
}

func NewV11HTTPSender(host string, port int, logger func(params ...any)) (*V11HTTPSender, error) {
	return &V11HTTPSender{
		postHost: host,
		postPort: port,
		Logger:   logger,
		UUIDBase: time.Now().UnixNano(),
	}, nil
}
