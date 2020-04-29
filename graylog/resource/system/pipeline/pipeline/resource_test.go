package pipeline

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccPipeline(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	ruleBody := ""

	postURLPath := "/api/system/pipelines/pipeline"
	resourceURLPath := postURLPath + "/5ea3e4122ab79c001275832c"
	resourceName := "graylog_pipeline.test"

	getRoute := flute.Route{
		Name: "get a pipeline",
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
		Name: "create a pipeline",
		Matcher: &flute.Matcher{
			Method: "POST",
		},
		Tester: &flute.Tester{
			Path:         postURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "description": "test",
  "source": "pipeline \"test\"\n  stage 0 match either\nend\n"
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				ruleBody = `{
  "id": "5ea3e4122ab79c001275832c",
  "title": "test",
  "description": "test",
  "source": "pipeline \"test\"\n  stage 0 match either\nend\n",
  "created_at": "2020-04-25T07:17:38.490Z",
  "modified_at": "2020-04-25T07:17:38.490Z",
  "stages": [
    {
      "stage": 0,
      "match_all": false,
      "rules": []
    }
  ],
  "errors": null
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "id": "5ea3e4122ab79c001275832c",
  "title": "test",
  "description": "test",
  "source": "pipeline \"test\"\n  stage 0 match either\nend\n",
  "created_at": "2020-04-25T07:17:38.490Z",
  "modified_at": "2020-04-25T07:17:38.490Z",
  "stages": [
    {
      "stage": 0,
      "match_all": false,
      "rules": []
    }
  ],
  "errors": null
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a pipeline",
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
resource "graylog_pipeline" "test" {
  source = <<EOF
pipeline "test"
  stage 0 match either
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
		Name: "update a pipeline",
		Matcher: &flute.Matcher{
			Method: "PUT",
		},
		Tester: &flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "description": "test updated",
  "source": "pipeline \"test\"\n  stage 0 match either\nend\n"
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				ruleBody = `{
  "id": "5ea3e4122ab79c001275832c",
  "title": "test",
  "description": "test updated",
  "source": "pipeline \"test\"\n  stage 0 match either\nend\n",
  "created_at": "2020-04-25T07:17:38.490Z",
  "modified_at": "2020-04-25T07:19:08.164Z",
  "stages": [
    {
      "stage": 0,
      "match_all": false,
      "rules": []
    }
  ],
  "errors": null
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "id": "5ea3e4122ab79c001275832c",
  "title": "test",
  "description": "test updated",
  "source": "pipeline \"test\"\n  stage 0 match either\nend\n",
  "created_at": "2020-04-25T07:17:38.490Z",
  "modified_at": "2020-04-25T07:19:08.164Z",
  "stages": [
    {
      "stage": 0,
      "match_all": false,
      "rules": []
    }
  ],
  "errors": null
}`,
		},
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_pipeline" "test" {
  source = <<EOF
pipeline "test"
  stage 0 match either
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
		Providers: testutil.SingleResourceProviders("graylog_pipeline", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
