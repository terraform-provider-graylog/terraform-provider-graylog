package stream

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccStream(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	streamBody := ""

	resourceURLPath := "/api/streams/5ea26bb42ab79c0012521287"
	resourceName := "graylog_stream.test"

	getRoute := flute.Route{
		Name: "get a stream",
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
					Body:       ioutil.NopCloser(strings.NewReader(streamBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a stream",
		Matcher: flute.Matcher{
			Method: "POST",
		},
		Tester: flute.Tester{
			Path:         "/api/streams",
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "matching_type": "AND",
  "description": "test",
  "title": "test",
	"remove_matches_from_default_stream": false,
  "index_set_id": "5e9861442ab79c0012e7d1c4"
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				streamBody = `{
  "id": "5ea26bb42ab79c0012521287",
  "creator_user_id": "admin",
  "outputs": [],
  "matching_type": "AND",
  "description": "test",
  "created_at": "2020-04-24T04:31:48.481Z",
  "disabled": true,
  "rules": [],
  "alert_conditions": [],
  "alert_receivers": {
    "emails": [],
    "users": []
  },
  "title": "test",
  "content_pack": null,
  "remove_matches_from_default_stream": false,
  "index_set_id": "5e9861442ab79c0012e7d1c4",
  "is_default": false
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "stream_id": "5ea26bb42ab79c0012521287"
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a stream",
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
resource "graylog_stream" "test" {
  title         = "test"
  index_set_id  = "5e9861442ab79c0012e7d1c4"
  disabled      = true
  matching_type = "AND"
  description   = "test"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "title", "test"),
			resource.TestCheckResourceAttr(resourceName, "description", "test"),
			resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
			resource.TestCheckResourceAttr(resourceName, "matching_type", "AND"),
			resource.TestCheckResourceAttr(resourceName, "index_set_id", "5e9861442ab79c0012e7d1c4"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a stream",
		Matcher: flute.Matcher{
			Method: "PUT",
		},
		Tester: flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "matching_type": "AND",
  "description": "test updated",
  "title": "test updated",
	"remove_matches_from_default_stream": false,
  "index_set_id": "5e9861442ab79c0012e7d1c4"
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				streamBody = `{
  "id": "5ea26bb42ab79c0012521287",
  "creator_user_id": "admin",
  "outputs": [],
  "matching_type": "AND",
  "description": "test updated",
  "created_at": "2020-04-24T04:31:48.481Z",
  "disabled": true,
  "rules": [],
  "alert_conditions": [],
  "alert_receivers": {
    "emails": [],
    "users": []
  },
  "title": "test updated",
  "content_pack": null,
  "remove_matches_from_default_stream": false,
  "index_set_id": "5e9861442ab79c0012e7d1c4",
  "is_default": false
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "id": "5ea26bb42ab79c0012521287",
  "creator_user_id": "admin",
  "outputs": [],
  "matching_type": "AND",
  "description": "test updated",
  "created_at": "2020-04-24T04:31:48.481Z",
  "disabled": true,
  "rules": [],
  "alert_conditions": [],
  "alert_receivers": {
    "emails": [],
    "users": []
  },
  "title": "test updated",
  "content_pack": null,
  "remove_matches_from_default_stream": false,
  "index_set_id": "5e9861442ab79c0012e7d1c4",
  "is_default": false
}`,
		},
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_stream" "test" {
  title         = "test updated"
  index_set_id  = "5e9861442ab79c0012e7d1c4"
  disabled      = true
  matching_type = "AND"
  description   = "test updated"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "title", "test updated"),
			resource.TestCheckResourceAttr(resourceName, "description", "test updated"),
			resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
			resource.TestCheckResourceAttr(resourceName, "matching_type", "AND"),
			resource.TestCheckResourceAttr(resourceName, "index_set_id", "5e9861442ab79c0012e7d1c4"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_stream", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
