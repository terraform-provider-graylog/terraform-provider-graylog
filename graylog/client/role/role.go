package role

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
	ctx context.Context, name string,
) (map[string]interface{}, *http.Response, error) {
	if name == "" {
		return nil, nil, errors.New("name is required")
	}

	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/roles/" + name,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Create(
	ctx context.Context, role interface{},
) (map[string]interface{}, *http.Response, error) {
	if role == nil {
		return nil, nil, errors.New("request body is nil")
	}

	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "POST",
		Path:         "/roles",
		RequestBody:  role,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Update(
	ctx context.Context, name string, role interface{},
) (map[string]interface{}, *http.Response, error) {
	if name == "" {
		return nil, nil, errors.New("name is required")
	}
	if role == nil {
		return nil, nil, errors.New("request body is nil")
	}

	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "PUT",
		Path:         "/roles/" + name,
		RequestBody:  role,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Delete(ctx context.Context, name string) (*http.Response, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method: "DELETE",
		Path:   "/roles/" + name,
	})
	return resp, err
}
