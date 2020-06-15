package output

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccOutput(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	outputBody := ""

	resourceURLPath := "/api/system/outputs/5ea2a4442ab79c001274d9dc"
	resourceName := "graylog_output.stdout"

	getRoute := flute.Route{
		Name: "get a output",
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
					Body:       ioutil.NopCloser(strings.NewReader(outputBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a output",
		Matcher: flute.Matcher{
			Method: "POST",
		},
		Tester: flute.Tester{
			Path:         "/api/system/outputs",
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "title": "stdout",
  "type": "org.graylog2.outputs.LoggingOutput",
  "configuration": {
    "prefix": "Writing message: "
  }
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				outputBody = `{
  "id": "5ea2a4442ab79c001274d9dc",
  "title": "stdout",
  "type": "org.graylog2.outputs.LoggingOutput",
  "creator_user_id": "admin",
  "created_at": "2020-04-24T08:33:08.136Z",
  "configuration": {
    "prefix": "Writing message: "
  },
  "content_pack": null
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
			BodyString: `{
  "id": "5ea2a4442ab79c001274d9dc",
  "title": "stdout",
  "type": "org.graylog2.outputs.LoggingOutput",
  "creator_user_id": "admin",
  "created_at": "2020-04-24T08:33:08.136Z",
  "configuration": {
    "prefix": "Writing message: "
  },
  "content_pack": null
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a output",
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
resource "graylog_output" "stdout" {
  title = "stdout"
  type  = "org.graylog2.outputs.LoggingOutput"

  configuration = <<EOF
{
  "prefix": "Writing message: "
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "title", "stdout"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a output",
		Matcher: flute.Matcher{
			Method: "PUT",
		},
		Tester: flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "title": "stdout updated",
  "type": "org.graylog2.outputs.LoggingOutput",
  "configuration": {
    "prefix": "Writing message (updated): "
  }
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				outputBody = `{
  "id": "5ea2a4442ab79c001274d9dc",
  "title": "stdout updated",
  "type": "org.graylog2.outputs.LoggingOutput",
  "creator_user_id": "admin",
  "created_at": "2020-04-24T08:33:08.136Z",
  "configuration": {
    "prefix": "Writing message (updated): "
  },
  "content_pack": null
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "_id": "5ea2a4442ab79c001274d9dc",
  "title": "stdout updated",
  "type": "org.graylog2.outputs.LoggingOutput",
  "creator_user_id": "admin",
  "configuration": {
    "prefix": "Writing message (updated): "
  },
  "created_at": "2020-04-24T08:33:08.136+0000",
  "content_pack": null
}`,
		},
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_output" "stdout" {
  title = "stdout updated"
  type  = "org.graylog2.outputs.LoggingOutput"

  configuration = <<EOF
{
  "prefix": "Writing message (updated): "
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "title", "stdout updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_output", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
