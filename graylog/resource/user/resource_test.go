package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccUser(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	userBody := ""

	getRoute := flute.Route{
		Name: "get a user",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         "/api/users/test",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(userBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a user",
		Matcher: flute.Matcher{
			Method: "POST",
		},
		Tester: flute.Tester{
			Path:         "/api/users",
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "username": "test",
  "email": "test@example.com",
  "password": "password",
  "full_name": "test test",
  "timezone": "",
  "session_timeout_ms": 3600000,
  "roles": ["Reader"],
  "permissions": []
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				userBody = `{
  "id": "5ea23d422ab79c001251dbfa",
  "username": "test",
  "email": "test@example.com",
  "full_name": "test test",
  "permissions": [
    "users:edit:test",
    "users:passwordchange:test",
    "users:tokencreate:test",
    "users:tokenlist:test",
    "users:tokenremove:test",
    "clusterconfigentry:read",
    "indexercluster:read",
    "messagecount:read",
    "journal:read",
    "messages:analyze",
    "inputs:read",
    "metrics:read",
    "savedsearches:edit",
    "fieldnames:read",
    "buffers:read",
    "system:read",
    "savedsearches:create",
    "jvmstats:read",
    "decorators:read",
    "throughput:read",
    "savedsearches:read",
    "messages:read"
  ],
  "preferences": {
    "updateUnfocussed": false,
    "enableSmartSearch": true
  },
  "timezone": null,
  "session_timeout_ms": 3600000,
  "external": false,
  "startpage": null,
  "roles": [
    "Reader"
  ],
  "read_only": false,
  "session_active": false,
  "last_activity": null,
  "client_address": null
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
		},
	}

	createStep := resource.TestStep{
		ResourceName: "graylog_user.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, postRoute)
		},
		Config: `
resource "graylog_user" "test" {
  username  = "test"
  email     = "test@example.com"
  password  = "password"
  full_name = "test test"
  roles = [
    "Reader",
  ]
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("graylog_user.test", "username", "test"),
			resource.TestCheckResourceAttr("graylog_user.test", "email", "test@example.com"),
			resource.TestCheckResourceAttr("graylog_user.test", "password", "password"),
			resource.TestCheckResourceAttr("graylog_user.test", "full_name", "test test"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a user",
		Matcher: flute.Matcher{
			Method: "PUT",
		},
		Tester: flute.Tester{
			Path:         "/api/users/test",
			PartOfHeader: testutil.Header(),
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				body := map[string]interface{}{}
				if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
					t.Fatal(err)
				}
				keys := []string{
					"email", "password", "full_name", "roles", "permissions",
					"session_timeout_ms", "username", "timezone",
				}
				if err := testutil.EqualMapKeys(body, keys...); err != nil {
					t.Fatal(err)
				}
				require.Equal(t, "test@example.com", body["email"])
				require.Equal(t, "password", body["password"])
				require.Equal(t, "test test updated", body["full_name"])
				require.Equal(t, []interface{}{"Reader"}, body["roles"])

				userBody = `{
  "id": "5ea23d422ab79c001251dbfa",
  "username": "test",
  "email": "test@example.com",
  "full_name": "test test updated",
  "permissions": [
    "users:edit:test",
    "users:passwordchange:test",
    "users:tokencreate:test",
    "users:tokenlist:test",
    "users:tokenremove:test",
    "clusterconfigentry:read",
    "indexercluster:read",
    "messagecount:read",
    "journal:read",
    "messages:analyze",
    "inputs:read",
    "metrics:read",
    "savedsearches:edit",
    "fieldnames:read",
    "buffers:read",
    "system:read",
    "savedsearches:create",
    "jvmstats:read",
    "decorators:read",
    "throughput:read",
    "savedsearches:read",
    "messages:read"
  ],
  "preferences": {
    "updateUnfocussed": false,
    "enableSmartSearch": true
  },
  "timezone": null,
  "session_timeout_ms": 3600000,
  "external": false,
  "startpage": null,
  "roles": [
    "Reader"
  ],
  "read_only": false,
  "session_active": false,
  "last_activity": null,
  "client_address": null
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a user",
		Matcher: flute.Matcher{
			Method: "DELETE",
		},
		Tester: flute.Tester{
			Path:         "/api/users/test",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 204,
			},
		},
	}

	updateStep := resource.TestStep{
		ResourceName: "graylog_user.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_user" "test" {
  username  = "test"
  email     = "test@example.com"
  password  = "password"
  full_name = "test test updated"
  roles = [
    "Reader",
  ]
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("graylog_user.test", "username", "test"),
			resource.TestCheckResourceAttr("graylog_user.test", "email", "test@example.com"),
			resource.TestCheckResourceAttr("graylog_user.test", "password", "password"),
			resource.TestCheckResourceAttr("graylog_user.test", "full_name", "test test updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_user", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
