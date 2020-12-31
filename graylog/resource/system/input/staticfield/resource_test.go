package staticfield

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccInputStaticFields(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	inputBody := map[string]interface{}{}

	if err := json.Unmarshal([]byte(`{
      "title": "gelf udp 2",
      "global": true,
      "name": "GELF UDP",
      "content_pack": null,
      "created_at": "2020-04-17T10:48:53.922Z",
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
      "id": "5e9989952ab79c001156f7d2"
    }`), &inputBody); err != nil {
		t.Fatal(err)
	}

	getURLPath := "/api/system/inputs/5e9989952ab79c001156f7d2"
	postURLPath := getURLPath + "/staticfields"
	resourceName := "graylog_input_static_fields.gelf_udp"

	getRoute := flute.Route{
		Name: "get a input staticfield",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         getURLPath,
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				b := &bytes.Buffer{}
				if err := json.NewEncoder(b).Encode(inputBody); err != nil {
					return nil, err
				}
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(b),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a input staticfield",
		Matcher: flute.Matcher{
			Method: "POST",
		},
		Tester: flute.Tester{
			Path:         postURLPath,
			PartOfHeader: testutil.Header(),
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				v := struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				}{}
				if err := json.NewDecoder(req.Body).Decode(&v); err != nil {
					t.Fatal(err)
				}
				fields := inputBody[KeyStaticFields].(map[string]interface{})
				fields[v.Key] = v.Value
				inputBody[KeyStaticFields] = fields
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a input staticfield",
		Matcher: flute.Matcher{
			Method: "DELETE",
		},
		Tester: flute.Tester{
			PartOfHeader: testutil.Header(),
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				a := strings.LastIndex(req.URL.Path, "/")
				key := req.URL.Path[a+1:]
				fields := inputBody[KeyStaticFields].(map[string]interface{})
				delete(fields, key)
				inputBody[KeyStaticFields] = fields
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
resource "graylog_input_static_fields" "gelf_udp" {
  input_id = "5e9989952ab79c001156f7d2"
  fields = {
    foo = "foo"
  }
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "input_id", "5e9989952ab79c001156f7d2"),
		),
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, postRoute, deleteRoute)
		},
		Config: `
resource "graylog_input_static_fields" "gelf_udp" {
  input_id = "5e9989952ab79c001156f7d2"
  fields = {
    foo = "bar"
  }
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "input_id", "5e9989952ab79c001156f7d2"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_input_static_fields", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
