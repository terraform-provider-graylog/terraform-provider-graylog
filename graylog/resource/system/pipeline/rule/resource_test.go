package rule

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/testutil"
)

func TestAccPipelineRule(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	ruleBody := ""

	postURLPath := "/api/system/pipelines/rule"
	resourceURLPath := postURLPath + "/5ea3e60f2ab79c00127585ac"
	resourceName := "graylog_pipeline_rule.test"

	getRoute := flute.Route{
		Name: "get a pipeline rule",
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
					Body:       ioutil.NopCloser(strings.NewReader(ruleBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a pipeline rule",
		Matcher: &flute.Matcher{
			Method: "POST",
		},
		Tester: &flute.Tester{
			Path:         postURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `  {
  "description": "test",
  "source": "rule \"test\"\nwhen\n    to_long($message.status) < 500\nthen\n    set_field(\"status_01\", 1);\nend\n"
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				ruleBody = `{
  "title": "test",
  "description": "test",
  "source": "rule \"test\"\nwhen\n    to_long($message.status) < 500\nthen\n    set_field(\"status_01\", 1);\nend\n",
  "created_at": "2020-04-25T07:26:07.322Z",
  "modified_at": "2020-04-25T07:26:07.322Z",
  "errors": null,
  "id": "5ea3e60f2ab79c00127585ac"
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "title": "test",
  "description": "test",
  "source": "rule \"test\"\nwhen\n    to_long($message.status) < 500\nthen\n    set_field(\"status_01\", 1);\nend\n",
  "created_at": "2020-04-25T07:26:07.322Z",
  "modified_at": "2020-04-25T07:26:07.322Z",
  "errors": null,
  "id": "5ea3e60f2ab79c00127585ac"
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a pipeline rule",
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
resource "graylog_pipeline_rule" "test" {
  source = <<EOF
rule "test"
when
    to_long($message.status) < 500
then
    set_field("status_01", 1);
end
EOF

  description = "test"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "description", "test"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a pipeline rule",
		Matcher: &flute.Matcher{
			Method: "PUT",
		},
		Tester: &flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `  {
  "description": "test updated",
  "source": "rule \"test\"\nwhen\n    to_long($message.status) < 500\nthen\n    set_field(\"status_01\", 1);\nend\n"
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				ruleBody = `{
  "title": "test",
  "description": "test updated",
  "source": "rule \"test\"\nwhen\n    to_long($message.status) < 500\nthen\n    set_field(\"status_01\", 1);\nend\n",
  "created_at": "2020-04-25T07:26:07.322Z",
  "modified_at": "2020-04-25T07:28:00.035Z",
  "errors": null,
  "id": "5ea3e60f2ab79c00127585ac"
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "title": "test",
  "description": "test updated",
  "source": "rule \"test\"\nwhen\n    to_long($message.status) < 500\nthen\n    set_field(\"status_01\", 1);\nend\n",
  "created_at": "2020-04-25T07:26:07.322Z",
  "modified_at": "2020-04-25T07:28:00.035Z",
  "errors": null,
  "id": "5ea3e60f2ab79c00127585ac"
}`,
		},
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_pipeline_rule" "test" {
  source = <<EOF
rule "test"
when
    to_long($message.status) < 500
then
    set_field("status_01", 1);
end
EOF

  description = "test updated"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "description", "test updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_pipeline_rule", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
