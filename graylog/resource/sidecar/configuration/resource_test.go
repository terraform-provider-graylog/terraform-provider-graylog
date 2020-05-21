package configuration

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccConfiguration(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	configurationBody := ""
	resourcePath := "/api/sidecar/configurations/5ec7027f2ab79c001226e584"
	resourceName := "graylog_sidecar_configuration.test"

	getRoute := flute.Route{
		Name: "get a configuration",
		Matcher: &flute.Matcher{
			Method: "GET",
		},
		Tester: &flute.Tester{
			Path:         resourcePath,
			PartOfHeader: testutil.Header(),
		},
		Response: &flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(configurationBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a configuration",
		Matcher: &flute.Matcher{
			Method: "POST",
		},
		Tester: &flute.Tester{
			Path:         "/api/sidecar/configurations",
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "collector_id": "5ec65adb2ab79c001226759c",
  "name": "foo",
  "color": "#00796b",
  "template": "fields_under_root: true"
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				configurationBody = `{
  "id": "5ec7027f2ab79c001226e584",
  "collector_id": "5ec65adb2ab79c001226759c",
  "name": "foo",
  "color": "#00796b",
  "template": "fields_under_root: true"
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "id": "5ec7027f2ab79c001226e584",
  "collector_id": "5ec65adb2ab79c001226759c",
  "name": "foo",
  "color": "#00796b",
  "template": "fields_under_root: true"
}`,
		},
	}

	createStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, postRoute)
		},
		Config: `
resource "graylog_sidecar_configuration" "test" {
  name         = "foo"
  color        = "#00796b"
  collector_id = "5ec65adb2ab79c001226759c"
  template     = "fields_under_root: true"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "name", "foo"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a configuration",
		Matcher: &flute.Matcher{
			Method: "PUT",
		},
		Tester: &flute.Tester{
			Path:         resourcePath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "collector_id": "5ec65adb2ab79c001226759c",
  "name": "foo_updated",
  "color": "#00796b",
  "template": "fields_under_root: true"
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				configurationBody = `{
  "id": "5ec7027f2ab79c001226e584",
  "collector_id": "5ec65adb2ab79c001226759c",
  "name": "foo_updated",
  "color": "#00796b",
  "template": "fields_under_root: true"
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "id": "5ec7027f2ab79c001226e584",
  "collector_id": "5ec65adb2ab79c001226759c",
  "name": "foo_updated",
  "color": "#00796b",
  "template": "fields_under_root: true"
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a configuration",
		Matcher: &flute.Matcher{
			Method: "DELETE",
		},
		Tester: &flute.Tester{
			Path:         resourcePath,
			PartOfHeader: testutil.Header(),
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
resource "graylog_sidecar_configuration" "test" {
  name         = "foo_updated"
  color        = "#00796b"
  collector_id = "5ec65adb2ab79c001226759c"
  template     = "fields_under_root: true"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "name", "foo_updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_sidecar_configuration", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
