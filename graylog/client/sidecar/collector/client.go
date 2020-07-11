package collector

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
		Path:         "/sidecar/collectors/" + id,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Create(
	ctx context.Context, collector interface{},
) (map[string]interface{}, *http.Response, error) {
	if collector == nil {
		return nil, nil, errors.New("request body is nil")
	}

	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "POST",
		Path:         "/sidecar/collectors",
		RequestBody:  collector,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Update(
	ctx context.Context, id string, collector interface{},
) (map[string]interface{}, *http.Response, error) {
	if id == "" {
		return nil, nil, errors.New("id is required")
	}
	if collector == nil {
		return nil, nil, errors.New("request body is nil")
	}

	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "PUT",
		Path:         "/sidecar/collectors/" + id,
		RequestBody:  collector,
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
		Path:   "/sidecar/collectors/" + id,
	})
	return resp, err
}
