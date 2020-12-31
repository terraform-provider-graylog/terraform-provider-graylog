package sidecar

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccSidecar(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	getRoute := flute.Route{
		Name: "get sidecars",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         "/api/sidecars/all",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{
  "query": "",
  "total": 1,
  "only_active": false,
  "sort": null,
  "order": null,
  "sidecars": [
    {
      "active": true,
      "node_id": "10517347-8fad-4f25-835f-95d1943fa338",
      "node_name": "test",
      "node_details": {
        "operating_system": "Linux",
        "ip": "172.18.0.2",
        "metrics": {
          "disks_75": [],
          "cpu_idle": 99.64,
          "load_1": 0.07
        },
        "log_file_list": null,
        "status": {
          "status": 2,
          "message": "0 running / 0 stopped / 1 failing",
          "collectors": [
            {
              "collector_id": "5ec661032ab79c0012267d29",
              "status": 2,
              "message": "Failed to find collector executable /usr/share/filebeat/bin/filebeat: exec: \"/usr/share/filebeat/bin/filebeat\": stat /usr/share/filebeat/bin/filebeat: no such file or directory",
              "verbose_message": ""
            }
          ]
        }
      },
      "assignments": [
        {
          "collector_id": "5ec661032ab79c0012267d29",
          "configuration_id": "5ec709642ab79c001226edf9"
        }
      ],
      "last_seen": "2020-06-06T02:15:36.318Z",
      "sidecar_version": "1.0.2",
      "collectors": null
    }
  ],
  "filters": null,
  "pagination": {
    "total": 1,
    "count": 1,
    "page": 1,
    "per_page": 1
  }
}`)),
				}, nil
			},
		},
	}

	readStep := resource.TestStep{
		ResourceName: "data.graylog_sidecar.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute)
		},
		Config: `
data "graylog_sidecar" "test" {
  node_name = "test"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.graylog_sidecar.test", "node_name", "test"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleDataSourceProviders("graylog_sidecar", DataSource()),
		Steps: []resource.TestStep{
			readStep,
		},
	})
}

func TestAccSidecar_BySidecarID(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	getRoute := flute.Route{
		Name: "get sidecar",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         "/api/sidecars/10517347-8fad-4f25-835f-95d1943fa338",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{
  "active": true,
  "node_id": "10517347-8fad-4f25-835f-95d1943fa338",
  "node_name": "cd406329ee4f",
  "node_details": {
    "operating_system": "Linux",
    "ip": "172.18.0.2",
    "metrics": {
      "disks_75": [],
      "cpu_idle": 99.55,
      "load_1": 0.04
    },
    "log_file_list": null,
    "status": {
      "status": 0,
      "message": "0 running / 0 stopped / 0 failing",
      "collectors": []
    }
  },
  "assignments": [
    {
      "collector_id": "5ec661032ab79c0012267d29",
      "configuration_id": "5ec709642ab79c001226edf9"
    }
  ],
  "last_seen": "2020-06-06T02:04:05.198Z",
  "sidecar_version": "1.0.2",
  "collectors": null
}`)),
				}, nil
			},
		},
	}

	readStep := resource.TestStep{
		ResourceName: "data.graylog_sidecar.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute)
		},
		Config: `
data "graylog_sidecar" "test" {
  node_id = "10517347-8fad-4f25-835f-95d1943fa338"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.graylog_sidecar.test", "node_name", "cd406329ee4f"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleDataSourceProviders("graylog_sidecar", DataSource()),
		Steps: []resource.TestStep{
			readStep,
		},
	})
}
