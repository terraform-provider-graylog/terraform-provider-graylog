package input

import (
	"context"
	"errors"
	"net/http"

	"github.com/suzuki-shunsuke/go-httpclient/httpclient"
)

type Client struct {
	Client httpclient.Client
}

func (cl Client) Get(ctx context.Context, id string) (map[string]interface{}, *http.Response, error) {
	if id == "" {
		return nil, nil, errors.New("id is required")
	}

	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/system/inputs/" + id,
		ResponseBody: &body,
	})
	return body, resp, err
}

func renameFieldAttributesToConfiguration(input map[string]interface{}) {
	// change attributes to configuration
	// https://github.com/Graylog2/graylog2-server/issues/3480
	if attr, ok := input["attributes"]; ok {
		input["configuration"] = attr
		delete(input, "attributes")
	}
}

func (cl Client) Create(
	ctx context.Context, input map[string]interface{},
) (map[string]interface{}, *http.Response, error) {
	if input == nil {
		return nil, nil, errors.New("request body is nil")
	}

	body := map[string]interface{}{}
	renameFieldAttributesToConfiguration(input)
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "POST",
		Path:         "/system/inputs",
		RequestBody:  input,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Update(
	ctx context.Context, id string, input map[string]interface{},
) (map[string]interface{}, *http.Response, error) {
	if id == "" {
		return nil, nil, errors.New("id is required")
	}
	if input == nil {
		return nil, nil, errors.New("request body is nil")
	}

	body := map[string]interface{}{}
	renameFieldAttributesToConfiguration(input)
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "PUT",
		Path:         "/system/inputs/" + id,
		RequestBody:  input,
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
		Path:   "/system/inputs/" + id,
	})
	return resp, err
}
