package stream

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
	ctx context.Context, id string,
) (map[string]interface{}, *http.Response, error) {
	if id == "" {
		return nil, nil, errors.New("id is required")
	}

	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/streams/" + id,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Gets(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/streams",
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Create(
	ctx context.Context, data map[string]interface{},
) (map[string]interface{}, *http.Response, error) {
	if data == nil {
		return nil, nil, errors.New("request body is nil")
	}

	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "POST",
		Path:         "/streams",
		RequestBody:  data,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Update(
	ctx context.Context, id string, data map[string]interface{},
) (map[string]interface{}, *http.Response, error) {
	if id == "" {
		return nil, nil, errors.New("id is required")
	}
	if data == nil {
		return nil, nil, errors.New("request body is nil")
	}
	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "PUT",
		Path:         "/streams/" + id,
		RequestBody:  data,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Delete(ctx context.Context, id string) (*http.Response, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method: "DELETE",
		Path:   "/streams/" + id,
	})
	return resp, err
}

func (cl Client) Resume(ctx context.Context, id string) (*http.Response, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method: "POST",
		Path:   "/streams/" + id + "/resume",
	})
	return resp, err
}

func (cl Client) Pause(ctx context.Context, id string) (*http.Response, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method: "POST",
		Path:   "/streams/" + id + "/pause",
	})
	return resp, err
}
