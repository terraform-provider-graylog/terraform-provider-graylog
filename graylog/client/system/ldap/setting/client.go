package setting

import (
	"context"
	"errors"
	"net/http"

	"github.com/suzuki-shunsuke/go-httpclient/httpclient"
)

type Client struct {
	Client httpclient.Client
}

func (cl Client) Get(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/system/ldap/settings",
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) Update(
	ctx context.Context, data map[string]interface{},
) (*http.Response, error) {
	if data == nil {
		return nil, errors.New("request body is nil")
	}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:      "PUT",
		Path:        "/system/ldap/settings",
		RequestBody: data,
	})
	return resp, err
}

func (cl Client) Delete(ctx context.Context) (*http.Response, error) {
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method: "DELETE",
		Path:   "/system/ldap/settings",
	})
	return resp, err
}
