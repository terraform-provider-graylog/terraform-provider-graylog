package condition

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccAmarmCallback(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	conditionBody := ""

	resourceURLPath := "/api/streams/5ea26bb42ab79c0012521287/alerts/conditions/a9a29806-e788-4e18-863b-64270a085500"
	resourceName := "graylog_alert_condition.test"

	getRoute := flute.Route{
		Name: "get a alert condition",
		Matcher: &flute.Matcher{
			Method: "GET",
		},
		Tester: &flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
		},
		Response: &flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(conditionBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a alert condition",
		Matcher: &flute.Matcher{
			Method: "POST",
		},
		Tester: &flute.Tester{
			Path:         "/api/streams/5ea26bb42ab79c0012521287/alerts/conditions",
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "type": "field_content_value",
  "title": "test",
  "parameters": {
    "backlog": 3,
    "repeat_notifications": false,
    "field": "message",
    "query": "*",
    "grace": 1,
    "value": "hoge hoge"
  }
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				conditionBody = `{
  "id": "a9a29806-e788-4e18-863b-64270a085500",
  "type": "field_content_value",
  "creator_user_id": "admin",
  "created_at": "2020-04-24T12:16:31.275+0000",
  "parameters": {
    "backlog": 3,
    "repeat_notifications": false,
    "field": "message",
    "query": "*",
    "grace": 1,
    "value": "hoge hoge"
  },
  "in_grace": false,
  "title": "test"
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
			BodyString: `{
  "alert_condition_id": "a9a29806-e788-4e18-863b-64270a085500"
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a alert condition",
		Matcher: &flute.Matcher{
			Method: "DELETE",
		},
		Tester: &flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
		},
		Response: &flute.Response{
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
resource "graylog_alert_condition" "test" {
  type      = "field_content_value"
  stream_id = "5ea26bb42ab79c0012521287"
  in_grace  = false
  title     = "test"

  parameters = <<EOF
{
  "backlog": 3,
  "repeat_notifications": false,
  "field": "message",
  "query": "*",
  "grace": 1,
  "value": "hoge hoge"
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "title", "test"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a alert condition",
		Matcher: &flute.Matcher{
			Method: "PUT",
		},
		Tester: &flute.Tester{
			Path:           resourceURLPath,
			PartOfHeader:   testutil.Header(),
			BodyJSONString: ``,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				conditionBody = `{
  "id": "a9a29806-e788-4e18-863b-64270a085500",
  "type": "field_content_value",
  "creator_user_id": "admin",
  "created_at": "2020-04-24T12:16:31.275+0000",
  "parameters": {
    "backlog": 4,
    "repeat_notifications": false,
    "field": "message",
    "query": "*",
    "grace": 1,
    "value": "hoge hoge"
  },
  "in_grace": false,
  "title": "test updated"
}`
			},
		},
		Response: &flute.Response{
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
resource "graylog_alert_condition" "test" {
  type      = "field_content_value"
  stream_id = "5ea26bb42ab79c0012521287"
  in_grace  = false
  title     = "test updated"

  parameters = <<EOF
{
  "backlog": 4,
  "repeat_notifications": false,
  "field": "message",
  "query": "*",
  "grace": 1,
  "value": "hoge hoge"
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "title", "test updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_alert_condition", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
