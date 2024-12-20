package factory

import (
	"fmt"
	"time"

	"github.com/dstgo/kratosx"
	"github.com/dstgo/kratosx/config"
	"github.com/dstgo/kratosx/library/http"
	json "github.com/json-iterator/go"

	"github.com/dstgo/cron/internal/client/service"
)

type HttpRequest struct {
	URL      string            `json:"url"`
	Params   map[string]any    `json:"params"`
	Method   string            `json:"method"`
	Header   map[string]string `json:"header"`
	Body     string            `json:"body"`
	BodyJson map[string]any    `json:"bodyJson"`
}

func (f *Factory) http(ctx kratosx.Context, task *service.ExecTaskRequest) (int, error) {
	var (
		data HttpRequest
		code = defaultErrorCode
		body string
		err  error
	)
	if err = json.Unmarshal([]byte(task.Value), &data); err != nil {
		return code, err
	}
	if data.Body != "" {
		body = data.Body
	}
	if len(data.BodyJson) != 0 {
		body, err = json.MarshalToString(data.BodyJson)
		if err != nil {
			return code, err
		}
	}

	request := http.New(&config.Http{
		EnableLog:        true,
		RetryCount:       int(task.RetryCount),
		RetryWaitTime:    1 * time.Second,
		MaxRetryWaitTime: time.Duration(task.MaxExecTime) * time.Second,
		Timeout:          60 * time.Second,
	}, ctx.Logger())

	response, err := request.Option(func(req *resty.Request) {
		req.URL = data.URL
		req.Method = data.Method
		for key, val := range data.Header {
			req.Header.Set(key, val)
		}
		for key, val := range data.Params {
			req.QueryParam.Set(key, fmt.Sprint(val))
		}
		req.Body = body
	}).Do()

	if err != nil {
		return defaultErrorCode, err
	}

	f.reply(task.Uuid, logInfo, string(response.Body()))
	return 0, nil
}
