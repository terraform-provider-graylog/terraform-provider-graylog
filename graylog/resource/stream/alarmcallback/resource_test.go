package alarmcallback

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccAlarmCallback(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	alarmcallbackBody := ""

	resourceURLPath := "/api/streams/5ea26bb42ab79c0012521287/alarmcallbacks/5ea2bc0a2ab79c001274e26f"
	resourceName := "graylog_alarm_callback.http"

	getRoute := flute.Route{
		Name: "get a alarmcallback",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(alarmcallbackBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a alarmcallback",
		Matcher: flute.Matcher{
			Method: "POST",
		},
		Tester: flute.Tester{
			Path:         "/api/streams/5ea26bb42ab79c0012521287/alarmcallbacks",
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "type": "org.graylog2.alarmcallbacks.HTTPAlarmCallback",
  "title": "test",
  "configuration": {
    "url": "https://example.com"
  }
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				alarmcallbackBody = `{
  "id": "5ea2bc0a2ab79c001274e26f",
  "type": "org.graylog2.alarmcallbacks.HTTPAlarmCallback",
  "configuration": {
    "url": "https://example.com"
  },
  "stream_id": "5ea26bb42ab79c0012521287",
  "title": "test",
  "created_at": "2020-04-24T10:14:34.389+0000",
  "creator_user_id": "admin"
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
			BodyString: `{
  "alarmcallback_id": "5ea2bc0a2ab79c001274e26f"
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a alarmcallback",
		Matcher: flute.Matcher{
			Method: "DELETE",
		},
		Tester: flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 204,
			},
		},
	}

	createStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, postRoute, deleteRoute)
		},
		Config: `
resource "graylog_alarm_callback" "http" {
  type      = "org.graylog2.alarmcallbacks.HTTPAlarmCallback"
  stream_id = "5ea26bb42ab79c0012521287"
  title     = "test"

  configuration = <<EOF
{
	"url": "https://example.com"
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "title", "test"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a alarmcallback",
		Matcher: flute.Matcher{
			Method: "PUT",
		},
		Tester: flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "type": "org.graylog2.alarmcallbacks.HTTPAlarmCallback",
  "title": "test updated",
  "configuration": {
    "url": "https://example.com/updated"
  }
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				alarmcallbackBody = `{
  "id": "5ea2bc0a2ab79c001274e26f",
  "type": "org.graylog2.alarmcallbacks.HTTPAlarmCallback",
  "configuration": {
    "url": "https://example.com/updated"
  },
  "stream_id": "5ea26bb42ab79c0012521287",
  "title": "test updated",
  "created_at": "2020-04-24T10:14:34.389+0000",
  "creator_user_id": "admin"
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 204,
			},
		},
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_alarm_callback" "http" {
  type      = "org.graylog2.alarmcallbacks.HTTPAlarmCallback"
  stream_id = "5ea26bb42ab79c0012521287"
  title     = "test updated"

  configuration = <<EOF
{
	"url": "https://example.com/updated"
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "title", "test updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_alarm_callback", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
