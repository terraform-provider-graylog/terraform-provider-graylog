package collector

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccCollector(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	collectorBody := ""
	resourcePath := "/api/sidecar/collectors/5ec65adb2ab79c001226759c"
	resourceName := "graylog_sidecar_collector.test"

	getRoute := flute.Route{
		Name: "get a collector",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         resourcePath,
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(collectorBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a collector",
		Matcher: flute.Matcher{
			Method: "POST",
		},
		Tester: flute.Tester{
			Path:         "/api/sidecar/collectors",
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "name": "foo",
  "service_type": "exec",
  "node_operating_system": "linux",
  "executable_path": "/usr/share/filebeat/bin/filebeat",
  "execute_parameters": "",
  "validation_parameters": "",
  "default_template": ""
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				collectorBody = `{
  "id": "5ec65adb2ab79c001226759c",
  "name": "foo",
  "service_type": "exec",
  "node_operating_system": "linux",
  "executable_path": "/usr/share/filebeat/bin/filebeat",
  "execute_parameters": "",
  "validation_parameters": "",
  "default_template": ""
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "id": "5ec65adb2ab79c001226759c",
  "name": "foo",
  "service_type": "exec",
  "node_operating_system": "linux",
  "executable_path": "/usr/share/filebeat/bin/filebeat",
  "execute_parameters": "",
  "validation_parameters": "",
  "default_template": ""
}`,
		},
	}

	createStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, postRoute)
		},
		Config: `
resource "graylog_sidecar_collector" "test" {
  name                  = "foo"
  service_type          = "exec"
  node_operating_system = "linux"
  executable_path       = "/usr/share/filebeat/bin/filebeat"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "name", "foo"),
			resource.TestCheckResourceAttr(resourceName, "node_operating_system", "linux"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a collector",
		Matcher: flute.Matcher{
			Method: "PUT",
		},
		Tester: flute.Tester{
			Path:         resourcePath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "name": "foo_updated",
  "service_type": "exec",
  "node_operating_system": "linux",
  "executable_path": "/usr/share/filebeat/bin/filebeat",
  "execute_parameters": "",
  "validation_parameters": "",
  "default_template": ""
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				collectorBody = `{
  "id": "5ec65adb2ab79c001226759c",
  "name": "foo_updated",
  "service_type": "exec",
  "node_operating_system": "linux",
  "executable_path": "/usr/share/filebeat/bin/filebeat",
  "execute_parameters": "",
  "validation_parameters": "",
  "default_template": ""
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "id": "5ec65adb2ab79c001226759c",
  "name": "foo_updated",
  "service_type": "exec",
  "node_operating_system": "linux",
  "executable_path": "/usr/share/filebeat/bin/filebeat",
  "execute_parameters": "",
  "validation_parameters": "",
  "default_template": ""
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a collector",
		Matcher: flute.Matcher{
			Method: "DELETE",
		},
		Tester: flute.Tester{
			Path:         resourcePath,
			PartOfHeader: testutil.Header(),
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
resource "graylog_sidecar_collector" "test" {
  name                  = "foo_updated"
  service_type          = "exec"
  node_operating_system = "linux"
  executable_path       = "/usr/share/filebeat/bin/filebeat"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "name", "foo_updated"),
			resource.TestCheckResourceAttr(resourceName, "node_operating_system", "linux"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_sidecar_collector", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
