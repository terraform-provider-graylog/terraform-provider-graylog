package sidecar

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
		Path:         "/sidecars/" + id,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) GetAll(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/sidecars/all",
		ResponseBody: &body,
	})
	return body, resp, err
}
