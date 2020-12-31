package grok

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccGrokPattern(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	grokBody := ""

	postURLPath := "/api/system/grok"
	resourceURLPath := postURLPath + "/5ea3ce6d2ab79c0012757673"
	resourceName := "graylog_grok_pattern.test"

	getRoute := flute.Route{
		Name: "get a event grok",
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
					Body:       ioutil.NopCloser(strings.NewReader(grokBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a event grok",
		Matcher: flute.Matcher{
			Method: "POST",
		},
		Tester: flute.Tester{
			Path:         postURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "name": "test",
  "pattern": "test"
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				grokBody = `{
  "name": "test",
  "pattern": "test",
  "content_pack": null,
  "id": "5ea3ce6d2ab79c0012757673"
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
			BodyString: `{
  "name": "test",
  "pattern": "test",
  "content_pack": null,
  "id": "5ea3ce6d2ab79c0012757673"
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a event grok",
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
resource "graylog_grok_pattern" "test" {
  name    = "test"
  pattern = "test"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "name", "test"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a event grok",
		Matcher: flute.Matcher{
			Method: "PUT",
		},
		Tester: flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "name": "test_updated",
  "pattern": "test"
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				grokBody = `{
  "name": "test_updated",
  "pattern": "test",
  "content_pack": null,
  "id": "5ea3ce6d2ab79c0012757673"
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "name": "test_updated",
  "pattern": "test",
  "content_pack": null,
  "id": "5ea3ce6d2ab79c0012757673"
}`,
		},
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_grok_pattern" "test" {
  name    = "test_updated"
  pattern = "test"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "name", "test_updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_grok_pattern", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
