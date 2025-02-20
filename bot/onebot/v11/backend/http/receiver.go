package onebot_v11_impl

import (
	"fmt"
	"io"
	"net/http"
)

type V11HTTPReceiver struct {
	listenPort int
	Logger     func(params ...interface{})
	Handler    func(rawData []byte)
}

func (receiver *V11HTTPReceiver) Log(format string, params ...any) {
	if receiver.Logger != nil {
		receiver.Logger(fmt.Sprintf(format, params))
	}
}

func (receiver *V11HTTPReceiver) HTTPServer(resp http.ResponseWriter, request *http.Request) {
	data, err := io.ReadAll(request.Body)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			receiver.Log("Failed to close HTTP request body, error:%s", err.Error())
		}
	}(request.Body)

	if receiver.Handler != nil {
		go receiver.Handler(data)
	}

	resp.WriteHeader(http.StatusOK)
}

func NewV11HTTPReceiver(port int, logger func(params ...any), handler func(rawData []byte)) (*V11HTTPReceiver, error) {
	receiver := &V11HTTPReceiver{
		listenPort: port,
		Logger:     logger,
		Handler:    handler,
	}
	httpHandler := func(writer http.ResponseWriter, request *http.Request) { receiver.HTTPServer(writer, request) }
	httpHost := fmt.Sprintf(":%d", receiver.listenPort)

	http.HandleFunc("/", httpHandler)
	go func() {
		_ = http.ListenAndServe(httpHost, nil)
	}()

	if logger != nil {
		logger(fmt.Sprintf("Bot starts serving HTTP on %v", httpHost))
	}

	return receiver, nil
}
