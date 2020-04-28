package input

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/testutil"
)

func TestAccUser(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	inputBody := ""

	getRoute := flute.Route{
		Name: "get a input",
		Matcher: &flute.Matcher{
			Method: "GET",
		},
		Tester: &flute.Tester{
			Path:         "/api/system/inputs/5ea252212ab79c001251f682",
			PartOfHeader: testutil.Header(),
		},
		Response: &flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(inputBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a input",
		Matcher: &flute.Matcher{
			Method: "POST",
		},
		Tester: &flute.Tester{
			Path:         "/api/system/inputs",
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "title": "gelf udp",
  "global": true,
  "type": "org.graylog2.inputs.gelf.udp.GELFUDPInput",
  "configuration": {
    "recv_buffer_size": 262144,
    "decompress_size_limit": 8388608,
    "bind_address": "0.0.0.0",
    "port": 12201
  }
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				inputBody = `{
  "title": "gelf udp",
  "global": true,
  "name": "GELF UDP",
  "content_pack": null,
  "created_at": "2020-04-24T02:42:41.927Z",
  "type": "org.graylog2.inputs.gelf.udp.GELFUDPInput",
  "creator_user_id": "admin",
  "attributes": {
    "recv_buffer_size": 262144,
    "decompress_size_limit": 8388608,
    "bind_address": "0.0.0.0",
    "port": 12201
  },
  "static_fields": {},
  "node": null,
  "id": "5ea252212ab79c001251f682"
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
			BodyString: `{
  "id": "5ea252212ab79c001251f682"
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a input",
		Matcher: &flute.Matcher{
			Method: "DELETE",
		},
		Tester: &flute.Tester{
			Path:         "/api/system/inputs/5ea252212ab79c001251f682",
			PartOfHeader: testutil.Header(),
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 204,
			},
		},
	}

	createStep := resource.TestStep{
		ResourceName: "graylog_input.gelf_udp",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, postRoute, deleteRoute)
		},
		Config: `
resource "graylog_input" "gelf_udp" {
  title  = "gelf udp"
  type   = "org.graylog2.inputs.gelf.udp.GELFUDPInput"
  global = true

  attributes = <<EOF
{
  "bind_address": "0.0.0.0",
	"port": 12201,
	"recv_buffer_size": 262144,
	"decompress_size_limit": 8388608
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("graylog_input.gelf_udp", "title", "gelf udp"),
			resource.TestCheckResourceAttr("graylog_input.gelf_udp", "type", "org.graylog2.inputs.gelf.udp.GELFUDPInput"),
			resource.TestCheckResourceAttr("graylog_input.gelf_udp", "global", "true"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a input",
		Matcher: &flute.Matcher{
			Method: "PUT",
		},
		Tester: &flute.Tester{
			Path:         "/api/system/inputs/5ea252212ab79c001251f682",
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "title": "gelf udp updated",
  "global": true,
  "type": "org.graylog2.inputs.gelf.udp.GELFUDPInput",
  "configuration": {
    "recv_buffer_size": 262144,
    "decompress_size_limit": 8388608,
    "bind_address": "0.0.0.0",
    "port": 12202
  }
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				inputBody = `{
  "title": "gelf udp updated",
  "global": true,
  "name": "GELF UDP",
  "content_pack": null,
  "created_at": "2020-04-24T02:46:17.976Z",
  "type": "org.graylog2.inputs.gelf.udp.GELFUDPInput",
  "creator_user_id": "admin",
  "attributes": {
    "recv_buffer_size": 262144,
    "decompress_size_limit": 8388608,
    "bind_address": "0.0.0.0",
    "port": 12202
  },
  "static_fields": {},
  "node": null,
  "id": "5ea252212ab79c001251f682"
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
			BodyString: `{
  "id": "5ea252212ab79c001251f682"
}`,
		},
	}

	updateStep := resource.TestStep{
		ResourceName: "graylog_input.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_input" "gelf_udp" {
  title  = "gelf udp updated"
  type   = "org.graylog2.inputs.gelf.udp.GELFUDPInput"
  global = true

  attributes = <<EOF
{
  "bind_address": "0.0.0.0",
	"port": 12202,
	"recv_buffer_size": 262144,
	"decompress_size_limit": 8388608
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("graylog_input.gelf_udp", "title", "gelf udp updated"),
			resource.TestCheckResourceAttr("graylog_input.gelf_udp", "type", "org.graylog2.inputs.gelf.udp.GELFUDPInput"),
			resource.TestCheckResourceAttr("graylog_input.gelf_udp", "global", "true"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_input", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
