package role

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccRole(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	roleBody := ""

	getRoute := flute.Route{
		Name: "get a role",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         "/api/roles/terraform",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(roleBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a role",
		Matcher: flute.Matcher{
			Method: "POST",
		},
		Tester: flute.Tester{
			Path:         "/api/roles",
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "name": "terraform",
  "description": "role description",
  "permissions": ["dashboards:*"]
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				roleBody = `{
  "name": "terraform",
  "description": "role description",
  "permissions": [
    "dashboards:*"
  ],
  "read_only": false
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
			BodyString: `{
  "name": "terraform",
  "description": "role description",
  "permissions": [
    "dashboards:*"
  ],
  "read_only": false
}`,
		},
	}

	createStep := resource.TestStep{
		ResourceName: "graylog_role.terraform",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, postRoute)
		},
		Config: `
resource "graylog_role" "terraform" {
  name        = "terraform"
  description = "role description"

  permissions = [
    "dashboards:*",
  ]
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("graylog_role.terraform", "name", "terraform"),
			resource.TestCheckResourceAttr("graylog_role.terraform", "description", "role description"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a role",
		Matcher: flute.Matcher{
			Method: "PUT",
		},
		Tester: flute.Tester{
			Path:         "/api/roles/terraform",
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "name": "terraform",
  "description": "role description updated",
  "permissions": ["dashboards:*"]
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				roleBody = `{
  "name": "terraform",
  "description": "role description updated",
  "permissions": [
    "dashboards:*"
  ],
  "read_only": false
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
			BodyString: `{
  "name": "terraform",
  "description": "role description updated",
  "permissions": [
    "dashboards:*"
  ],
  "read_only": false
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a role",
		Matcher: flute.Matcher{
			Method: "DELETE",
		},
		Tester: flute.Tester{
			Path:         "/api/roles/terraform",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 204,
			},
		},
	}

	updateStep := resource.TestStep{
		ResourceName: "graylog_role.terraform",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_role" "terraform" {
  name        = "terraform"
  description = "role description updated"

  permissions = [
    "dashboards:*",
  ]
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("graylog_role.terraform", "name", "terraform"),
			resource.TestCheckResourceAttr("graylog_role.terraform", "description", "role description updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_role", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
