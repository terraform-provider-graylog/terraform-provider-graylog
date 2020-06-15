package connection

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccPipelineConnection(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	connectionBody := ""

	postURLPath := "/api/system/pipelines/connections/to_stream"
	resourceName := "graylog_pipeline_connection.test"

	getRoute := flute.Route{
		Name: "get a pipeline connection",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         "/api/system/pipelines/connections/5ea26bb42ab79c0012521287",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(connectionBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a pipeline connection",
		Matcher: flute.Matcher{
			Method: "POST",
			BodyJSONString: `{
			  "stream_id": "5ea26bb42ab79c0012521287",
			  "pipeline_ids": ["5ea3e4122ab79c001275832c"]
			}`,
		},
		Tester: flute.Tester{
			Path:         postURLPath,
			PartOfHeader: testutil.Header(),
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				connectionBody = `{
  "stream_id": "5ea26bb42ab79c0012521287",
  "pipeline_ids": [
    "5ea3e4122ab79c001275832c"
  ],
  "id": "5ea3f2302ab79c0012759085"
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "stream_id": "5ea26bb42ab79c0012521287",
  "pipeline_ids": [
    "5ea3e4122ab79c001275832c"
  ],
  "id": "5ea3f2302ab79c0012759085"
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a pipeline connection",
		Matcher: flute.Matcher{
			Method: "POST",
			// BodyJSONString: `{
			//   "stream_id": "5ea26bb42ab79c0012521287",
			//   "pipeline_ids": []
			// }`,
		},
		Tester: flute.Tester{
			Path:         postURLPath,
			PartOfHeader: testutil.Header(),
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				connectionBody = `{
  "stream_id": "5ea26bb42ab79c0012521287",
  "pipeline_ids": [],
  "id": "5ea3f2302ab79c0012759085"
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "stream_id": "5ea26bb42ab79c0012521287",
  "pipeline_ids": [],
  "id": "5ea3f2302ab79c0012759085"
}`,
		},
	}

	createStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, postRoute, deleteRoute)
		},
		Config: `
resource "graylog_pipeline_connection" "test" {
  stream_id    = "5ea26bb42ab79c0012521287"
  pipeline_ids = ["5ea3e4122ab79c001275832c"]
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "stream_id", "5ea26bb42ab79c0012521287"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a pipeline connection",
		Matcher: flute.Matcher{
			Method: "POST",
			BodyJSONString: `{
  "stream_id": "5ea26bb42ab79c0012521287",
  "pipeline_ids": ["5ea3e4122ab79c001275832c","5ea3f3442ab79c00127591d6"]
}`,
		},
		Tester: flute.Tester{
			Path:         postURLPath,
			PartOfHeader: testutil.Header(),
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				connectionBody = `{
  "stream_id": "5ea26bb42ab79c0012521287",
  "pipeline_ids": [
    "5ea3e4122ab79c001275832c",
    "5ea3f3442ab79c00127591d6"
  ],
  "id": "5ea3f2302ab79c0012759085"
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "stream_id": "5ea26bb42ab79c0012521287",
  "pipeline_ids": [
    "5ea3e4122ab79c001275832c",
    "5ea3f3442ab79c00127591d6"
  ],
  "id": "5ea3f2302ab79c0012759085"
}`,
		},
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_pipeline_connection" "test" {
  stream_id    = "5ea26bb42ab79c0012521287"
  pipeline_ids = ["5ea3e4122ab79c001275832c", "5ea3f3442ab79c00127591d6"]
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "stream_id", "5ea26bb42ab79c0012521287"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_pipeline_connection", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
