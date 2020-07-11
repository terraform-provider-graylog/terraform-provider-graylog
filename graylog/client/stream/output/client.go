package output

import (
	"context"
	"errors"
	"net/http"

	"github.com/suzuki-shunsuke/go-httpclient/httpclient"
)

type Client struct {
	Client httpclient.Client
}

func (cl Client) GetOutputsOfStream(
	ctx context.Context, streamID string,
) (map[string]interface{}, *http.Response, error) {
	if streamID == "" {
		return nil, nil, errors.New("stream id is required")
	}
	body := map[string]interface{}{}
	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method:       "GET",
		Path:         "/streams/" + streamID + "/outputs",
		ResponseBody: &body,
	})
	return body, resp, err
}

func (cl Client) AssociateOutputsWithStream(
	ctx context.Context, streamID string, outputIDs []string,
) (*http.Response, error) {
	if streamID == "" {
		return nil, errors.New("stream id is required")
	}

	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method: "POST",
		Path:   "/streams/" + streamID + "/outputs",
		RequestBody: map[string]interface{}{
			"outputs": outputIDs,
		},
	})
	return resp, err
}

func (cl Client) Delete(ctx context.Context, streamID, id string) (*http.Response, error) {
	if streamID == "" {
		return nil, errors.New("stream id is required")
	}
	if id == "" {
		return nil, errors.New("output id is required")
	}

	resp, err := cl.Client.Call(ctx, httpclient.CallParams{
		Method: "DELETE",
		Path:   "/streams/" + streamID + "/outputs/" + id,
	})
	return resp, err
}
