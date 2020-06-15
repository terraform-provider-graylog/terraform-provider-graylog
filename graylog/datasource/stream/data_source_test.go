package stream

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccStream(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	getRoute := flute.Route{
		Name: "get streams",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         "/api/streams",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{
  "total": 3,
  "streams": [
    {
      "id": "000000000000000000000001",
      "creator_user_id": "local:admin",
      "outputs": [],
      "matching_type": "AND",
      "description": "Stream containing all messages",
      "created_at": "2020-04-28T12:08:21.010Z",
      "disabled": false,
      "rules": [],
      "alert_conditions": [],
      "alert_receivers": {
        "emails": [],
        "users": []
      },
      "title": "All messages",
      "content_pack": null,
      "remove_matches_from_default_stream": false,
      "index_set_id": "5ea81cb42ab79c00129dbe58",
      "is_default": true
    },
    {
      "id": "000000000000000000000003",
      "creator_user_id": "admin",
      "outputs": [],
      "matching_type": "AND",
      "description": "Stream containing all system events created by Graylog",
      "created_at": "2020-04-28T12:08:32.193Z",
      "disabled": false,
      "rules": [],
      "alert_conditions": [],
      "alert_receivers": {
        "emails": [],
        "users": []
      },
      "title": "All system events",
      "content_pack": null,
      "remove_matches_from_default_stream": true,
      "index_set_id": "5ea81cc02ab79c00129dbf1f",
      "is_default": false
    },
    {
      "id": "000000000000000000000002",
      "creator_user_id": "admin",
      "outputs": [],
      "matching_type": "AND",
      "description": "Stream containing all events created by Graylog",
      "created_at": "2020-04-28T12:08:32.186Z",
      "disabled": false,
      "rules": [],
      "alert_conditions": [],
      "alert_receivers": {
        "emails": [],
        "users": []
      },
      "title": "All events",
      "content_pack": null,
      "remove_matches_from_default_stream": true,
      "index_set_id": "5ea81cc02ab79c00129dbf1c",
      "is_default": false
    }
  ]
}`)),
				}, nil
			},
		},
	}

	readStep := resource.TestStep{
		ResourceName: "data.graylog_stream.all_events",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute)
		},
		Config: `
data "graylog_stream" "all_events" {
  title = "All events"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.graylog_stream.all_events", "title", "All events"),
			resource.TestCheckResourceAttr("data.graylog_stream.all_events", "description", "Stream containing all events created by Graylog"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleDataSourceProviders("graylog_stream", DataSource()),
		Steps: []resource.TestStep{
			readStep,
		},
	})
}

func TestAccStream_byStreamID(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	getRoute := flute.Route{
		Name: "get streams",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         "/api/streams/000000000000000000000003",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{
  "id": "000000000000000000000003",
  "creator_user_id": "admin",
  "outputs": [],
  "matching_type": "AND",
  "description": "Stream containing all system events created by Graylog",
  "created_at": "2020-04-28T12:08:32.193Z",
  "disabled": false,
  "rules": [],
  "alert_conditions": [],
  "alert_receivers": {
    "emails": [],
    "users": []
  },
  "title": "All system events",
  "content_pack": null,
  "remove_matches_from_default_stream": true,
  "index_set_id": "5ea81cc02ab79c00129dbf1f",
  "is_default": false
}
`)),
				}, nil
			},
		},
	}

	readStep := resource.TestStep{
		ResourceName: "data.graylog_stream.all_system_events",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute)
		},
		Config: `
data "graylog_stream" "all_system_events" {
  stream_id = "000000000000000000000003"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.graylog_stream.all_system_events", "title", "All system events"),
			resource.TestCheckResourceAttr("data.graylog_stream.all_system_events", "description", "Stream containing all system events created by Graylog"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleDataSourceProviders("graylog_stream", DataSource()),
		Steps: []resource.TestStep{
			readStep,
		},
	})
}
