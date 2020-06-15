package notification

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccEventNotification(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	notificationBody := ""

	postURLPath := "/api/events/notifications"
	resourceURLPath := postURLPath + "/5ea3c1d72ab79c00127567fe"
	resourceName := "graylog_event_notification.http"

	getRoute := flute.Route{
		Name: "get a event notification",
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
					Body:       ioutil.NopCloser(strings.NewReader(notificationBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a event notification",
		Matcher: flute.Matcher{
			Method: "POST",
		},
		Tester: flute.Tester{
			Path:         postURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "title": "http",
  "description": "http notification",
  "config": {
    "type": "http-notification-v1",
    "url": "http://example.com"
  }
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				notificationBody = `{
  "id": "5ea3c1d72ab79c00127567fe",
  "title": "http",
  "description": "http notification",
  "config": {
    "type": "http-notification-v1",
    "url": "http://example.com"
  }
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "id": "5ea3c1d72ab79c00127567fe",
  "title": "http",
  "description": "http notification",
  "config": {
    "type": "http-notification-v1",
    "url": "http://example.com"
  }
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a event notification",
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
resource "graylog_event_notification" "http" {
  title       = "http"
  description = "http notification"

  config = <<EOF
{
  "type": "http-notification-v1",
  "url": "http://example.com"
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "description", "http notification"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a event notification",
		Matcher: flute.Matcher{
			Method: "PUT",
		},
		Tester: flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "id": "5ea3c1d72ab79c00127567fe",
  "title": "http",
  "description": "http notification updated",
  "config": {
    "type": "http-notification-v1",
    "url": "http://example.com"
  }
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				notificationBody = `{
  "id": "5ea3c1d72ab79c00127567fe",
  "title": "http",
  "description": "http notification updated",
  "config": {
    "type": "http-notification-v1",
    "url": "http://example.com"
  }
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "id": "5ea3c1d72ab79c00127567fe",
  "title": "http",
  "description": "http notification updated",
  "config": {
    "type": "http-notification-v1",
    "url": "http://example.com"
  }
}`,
		},
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_event_notification" "http" {
  title       = "http"
  description = "http notification updated"

  config = <<EOF
{
  "type": "http-notification-v1",
  "url": "http://example.com"
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "description", "http notification updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_event_notification", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
