package sidecar

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

const createdBody = `{
  "query": "",
  "total": 1,
  "only_active": false,
  "sort": null,
  "order": null,
  "sidecars": [
    {
      "active": true,
      "node_id": "10517347-8fad-4f25-835f-95d1943fa338",
      "node_name": "cd406329ee4f",
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
}`

const updatedBody = `{
  "query": "",
  "total": 1,
  "only_active": false,
  "sort": null,
  "order": null,
  "sidecars": [
    {
      "active": true,
      "node_id": "10517347-8fad-4f25-835f-95d1943fa338",
      "node_name": "cd406329ee4f",
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
          "collector_id": "5ec661032ab79c0012267555",
          "configuration_id": "5ec709642ab79c001226e555"
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
}`

const deletedBody = `{
  "query": "",
  "total": 1,
  "only_active": false,
  "sort": null,
  "order": null,
  "sidecars": [
    {
      "active": true,
      "node_id": "10517347-8fad-4f25-835f-95d1943fa338",
      "node_name": "cd406329ee4f",
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
      "assignments": [],
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
}`

func TestAccSidecar(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	sidecarBody := ""
	resourcePath := "/api/sidecars/all"
	resourceName := "graylog_sidecars.test"

	getRoute := flute.Route{
		Name: "get a sidecar",
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
					Body:       ioutil.NopCloser(strings.NewReader(sidecarBody)),
				}, nil
			},
		},
	}

	var cnt = 0

	updateRoute := flute.Route{
		Name: "update a sidecar",
		Matcher: &flute.Matcher{
			Method: "PUT",
		},
		Tester: &flute.Tester{
			Path:         "/api/sidecars/configurations",
			PartOfHeader: testutil.Header(),
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				body := map[string]interface{}{}
				if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
					t.Fatal("request body should be JSON", err)
				}
				switch cnt {
				case 0:
					require.Equal(t, map[string]interface{}{
						"nodes": []interface{}{
							map[string]interface{}{
								"node_id": "10517347-8fad-4f25-835f-95d1943fa338",
								"assignments": []interface{}{
									map[string]interface{}{
										"collector_id":     "5ec661032ab79c0012267d29",
										"configuration_id": "5ec709642ab79c001226edf9",
									},
								},
							},
						},
					}, body)
					sidecarBody = createdBody
				case 1:
					require.Equal(t, map[string]interface{}{
						"nodes": []interface{}{},
					}, body)
					sidecarBody = deletedBody
				case 2:
					require.Equal(t, map[string]interface{}{
						"nodes": []interface{}{
							map[string]interface{}{
								"node_id": "10517347-8fad-4f25-835f-95d1943fa338",
								"assignments": []interface{}{
									map[string]interface{}{
										"collector_id":     "5ec661032ab79c0012267555",
										"configuration_id": "5ec709642ab79c001226e555",
									},
								},
							},
						},
					}, body)
					sidecarBody = updatedBody
				}
				cnt++
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 202,
			},
		},
	}

	createStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute)
		},
		Config: `
resource "graylog_sidecars" "all" {
  sidecars {
    node_id = "10517347-8fad-4f25-835f-95d1943fa338"
    assignments {
      collector_id     = "5ec661032ab79c0012267d29"
      configuration_id = "5ec709642ab79c001226edf9"
    }
  }
}
`,
		Check: resource.ComposeTestCheckFunc(),
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute)
		},
		Config: `
resource "graylog_sidecars" "all" {
  sidecars {
    node_id = "10517347-8fad-4f25-835f-95d1943fa338"
    assignments {
      collector_id     = "5ec661032ab79c0012267555"
      configuration_id = "5ec709642ab79c001226e555"
    }
  }
}
`,
		Check: resource.ComposeTestCheckFunc(),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_sidecars", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
