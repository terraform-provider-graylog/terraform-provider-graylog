package staticfield

import (
	"context"
	"net/http"

	"github.com/suzuki-shunsuke/go-httpclient/httpclient"
)

type Client struct {
	Client httpclient.Client
}

func (cl Client) Create(
	ctx context.Context, inputID, key, value string,
) (*http.Response, error) {
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method: "POST",
		Path:   "/system/inputs/" + inputID + "/staticfields",
		RequestBody: map[string]string{
			"key":   key,
			"value": value,
		},
	})
	return resp, err
}

func (cl Client) Delete(ctx context.Context, inputID, id string) (*http.Response, error) {
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method: "DELETE",
		Path:   "/system/inputs/" + inputID + "/staticfields/" + id,
	})
	return resp, err
}
