package widget

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
	ctx context.Context, dashboardID, id string,
) (map[string]interface{}, *http.Response, error) {
	if dashboardID == "" {
		return nil, nil, errors.New("dashboard id is required")
	}
	if id == "" {
		return nil, nil, errors.New("dashboard widget id is required")
	}

	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/dashboards/" + dashboardID + "/widgets/" + id,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Create(
	ctx context.Context, dashboardID string, data map[string]interface{},
) (map[string]interface{}, *http.Response, error) {
	if dashboardID == "" {
		return nil, nil, errors.New("dashboard id is required")
	}
	if data == nil {
		return nil, nil, errors.New("request body is nil")
	}

	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "POST",
		Path:         "/dashboards/" + dashboardID + "/widgets",
		RequestBody:  data,
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Update(
	ctx context.Context, dashboardID, id string, data map[string]interface{},
) (*http.Response, error) {
	if dashboardID == "" {
		return nil, errors.New("dashboard id is required")
	}
	if id == "" {
		return nil, errors.New("dashboard widget id is required")
	}
	if data == nil {
		return nil, errors.New("request body is nil")
	}

	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:      "PUT",
		Path:        "/dashboards/" + dashboardID + "/widgets/" + id,
		RequestBody: data,
	})
	return resp, err
}

func (cl Client) Delete(ctx context.Context, dashboardID, id string) (*http.Response, error) {
	if dashboardID == "" {
		return nil, errors.New("dashboard id is required")
	}
	if id == "" {
		return nil, errors.New("dashboard widget id is required")
	}

	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method: "DELETE",
		Path:   "/dashboards/" + dashboardID + "/widgets/" + id,
	})
	return resp, err
}
