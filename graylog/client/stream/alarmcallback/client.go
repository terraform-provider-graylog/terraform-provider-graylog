package alarmcallback

import (
	"context"
	"errors"
	"net/http"

	"github.com/suzuki-shunsuke/go-httpclient/httpclient"
)

type Client struct {
	Client httpclient.Client
}

func (cl Client) Get(
	ctx context.Context, streamID, id string,
) (map[string]interface{}, *http.Response, error) {
	if streamID == "" {
		return nil, nil, errors.New("stream id is required")
	}
	if id == "" {
		return nil, nil, errors.New("id is required")
	}

	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/streams/" + streamID + "/alarmcallbacks/" + id,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Create(
	ctx context.Context, streamID string, data map[string]interface{},
) (map[string]interface{}, *http.Response, error) {
	if streamID == "" {
		return nil, nil, errors.New("stream id is required")
	}
	if data == nil {
		return nil, nil, errors.New("request body is nil")
	}

	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "POST",
		Path:         "/streams/" + streamID + "/alarmcallbacks",
		RequestBody:  data,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Update(
	ctx context.Context, streamID, id string, data map[string]interface{},
) (*http.Response, error) {
	if streamID == "" {
		return nil, errors.New("stream id is required")
	}
	if id == "" {
		return nil, errors.New("id is required")
	}
	if data == nil {
		return nil, errors.New("request body is nil")
	}

	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:      "PUT",
		Path:        "/streams/" + streamID + "/alarmcallbacks/" + id,
		RequestBody: data,
	})
	return resp, err
}

func (cl Client) Delete(ctx context.Context, streamID, id string) (*http.Response, error) {
	if streamID == "" {
		return nil, errors.New("stream id is required")
	}
	if id == "" {
		return nil, errors.New("id is required")
	}

	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method: "DELETE",
		Path:   "/streams/" + streamID + "/alarmcallbacks/" + id,
	})
	return resp, err
}
