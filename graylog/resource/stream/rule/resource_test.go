package rule

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccStreamRule(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	ruleBody := ""

	postURLPath := "/api/streams/5ea26bb42ab79c0012521287/rules"
	resourceURLPath := postURLPath + "/5ea3de782ab79c0012757c27"
	resourceName := "graylog_stream_rule.test"

	getRoute := flute.Route{
		Name: "get a stream rule",
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
					Body:       ioutil.NopCloser(strings.NewReader(ruleBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a stream rule",
		Matcher: flute.Matcher{
			Method: "POST",
		},
		Tester: flute.Tester{
			Path:         postURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "type": 1,
  "field": "tag",
  "value": "4",
  "description": "test",
  "inverted": false
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				ruleBody = `{
  "field": "tag",
  "stream_id": "5ea26bb42ab79c0012521287",
  "description": "test",
  "id": "5ea3de782ab79c0012757c27",
  "type": 1,
  "inverted": false,
  "value": "4"
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
			BodyString: `{
  "streamrule_id": "5ea3de782ab79c0012757c27"
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a stream rule",
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
resource "graylog_stream_rule" "test" {
  field       = "tag"
  value       = "4"
  stream_id   = "5ea26bb42ab79c0012521287"
  description = "test"
  type        = 1
  inverted    = false
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "description", "test"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a stream rule",
		Matcher: flute.Matcher{
			Method: "PUT",
		},
		Tester: flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "type": 1,
  "field": "tag",
  "value": "4",
  "description": "test updated",
  "inverted": false
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				ruleBody = `{
  "field": "tag",
  "stream_id": "5ea26bb42ab79c0012521287",
  "description": "test updated",
  "id": "5ea3de782ab79c0012757c27",
  "type": 1,
  "inverted": false,
  "value": "4"
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "streamrule_id": "5ea3de782ab79c0012757c27"
}`,
		},
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_stream_rule" "test" {
  field       = "tag"
  value       = "4"
  stream_id   = "5ea26bb42ab79c0012521287"
  description = "test updated"
  type        = 1
  inverted    = false
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "description", "test updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_stream_rule", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
