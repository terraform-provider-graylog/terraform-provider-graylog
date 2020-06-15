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

func TestAccStreamOutput(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	outputBody := ""

	postURLPath := "/api/streams/5ea26bb42ab79c0012521287/outputs"
	resourceName := "graylog_stream_output.test"

	getRoute := flute.Route{
		Name: "get a stream output",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         postURLPath,
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
		Name: "create a stream output",
		Matcher: flute.Matcher{
			Method: "POST",
			BodyJSONString: `{
  "outputs": ["5ea2a4442ab79c001274d9dc"]
}`,
		},
		Tester: flute.Tester{
			Path:         postURLPath,
			PartOfHeader: testutil.Header(),
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				outputBody = `{
  "total": 1,
  "outputs": [
    {
      "id": "5ea2a4442ab79c001274d9dc",
      "title": "stdout",
      "type": "org.graylog2.outputs.LoggingOutput",
      "creator_user_id": "admin",
      "created_at": "2020-04-24T08:33:08.136Z",
      "configuration": {
        "prefix": "Writing message: "
      },
      "content_pack": null
    }
  ]
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 202,
			},
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a stream output",
		Matcher: flute.Matcher{
			Method: "DELETE",
		},
		Tester: flute.Tester{
			PartOfHeader: testutil.Header(),
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				outputBody = ``
			},
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
resource "graylog_stream_output" "test" {
  stream_id = "5ea26bb42ab79c0012521287"
  output_ids = [
	  "5ea2a4442ab79c001274d9dc"
  ]
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "stream_id", "5ea26bb42ab79c0012521287"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a stream output",
		Matcher: flute.Matcher{
			Method: "POST",
			//			BodyJSONString: `{
			//  "outputs": ["5ea2a4442ab79c001274d9dc", "5e9989962ab79c001156f7db"]
			//}`,
		},
		Tester: flute.Tester{
			Path:         postURLPath,
			PartOfHeader: testutil.Header(),
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				outputBody = `{
  "total": 2,
  "outputs": [
    {
      "id": "5ea2a4442ab79c001274d9dc",
      "title": "stdout updated",
      "type": "org.graylog2.outputs.LoggingOutput",
      "creator_user_id": "admin",
      "created_at": "2020-04-24T08:33:08.136Z",
      "configuration": {
        "prefix": "Writing message (updated): "
      },
      "content_pack": null
    },
    {
      "id": "5e9989962ab79c001156f7db",
      "title": "gelf",
      "type": "org.graylog2.outputs.GelfOutput",
      "creator_user_id": "admin",
      "created_at": "2020-04-17T10:48:54.009Z",
      "configuration": {
        "hostname": "localhost",
        "protocol": "TCP",
        "connect_timeout": 1000,
        "queue_size": 512,
        "reconnect_delay": 500,
        "port": 12201,
        "max_inflight_sends": 512,
        "tcp_keep_alive": false,
        "tcp_no_delay": false,
        "tls_trust_cert_chain": "",
        "tls_verification_enabled": false
      },
      "content_pack": null
    }
  ]
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 202,
			},
		},
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_stream_output" "test" {
  stream_id = "5ea26bb42ab79c0012521287"
  output_ids = [
	  "5ea2a4442ab79c001274d9dc",
    "5e9989962ab79c001156f7db"
  ]
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "stream_id", "5ea26bb42ab79c0012521287"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_stream_output", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
