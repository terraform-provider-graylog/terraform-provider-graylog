package connection

import (
	"context"
	"errors"
	"net/http"

	"github.com/suzuki-shunsuke/go-httpclient/httpclient"
)

type Client struct {
	Client httpclient.Client
}

func (cl Client) GetConnectionsOfStream(
	ctx context.Context, streamID string,
) (map[string]interface{}, *http.Response, error) {
	if streamID == "" {
		return nil, nil, errors.New("stream id is required")
	}
	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/system/pipelines/connections/" + streamID,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) ConnectPipelinesToStream(
	ctx context.Context, data map[string]interface{},
) (*http.Response, error) {
	if data == nil {
		return nil, errors.New("request body is nil")
	}

	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:      "POST",
		Path:        "/system/pipelines/connections/to_stream",
		RequestBody: data,
	})
	return resp, err
}
