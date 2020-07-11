package position

import (
	"context"
	"errors"
	"net/http"

	"github.com/suzuki-shunsuke/go-httpclient/httpclient"
)

type Client struct {
	Client httpclient.Client
}

func (cl Client) Update(
	ctx context.Context, dashboardID string, data map[string]interface{},
) (*http.Response, error) {
	if dashboardID == "" {
		return nil, errors.New("dashboard id is required")
	}
	if data == nil {
		return nil, errors.New("request body is nil")
	}

	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:      "PUT",
		Path:        "/dashboards/" + dashboardID + "/positions",
		RequestBody: data,
	})
	return resp, err
}
