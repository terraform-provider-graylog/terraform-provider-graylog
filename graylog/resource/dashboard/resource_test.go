package dashboard

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

	dashboardBody := ""

	getRoute := flute.Route{
		Name: "get a dashboard",
		Matcher: &flute.Matcher{
			Method: "GET",
		},
		Tester: &flute.Tester{
			Path:         "/api/dashboards/5ea24b8c2ab79c001251ee46",
			PartOfHeader: testutil.Header(),
		},
		Response: &flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(dashboardBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a dashboard",
		Matcher: &flute.Matcher{
			Method: "POST",
		},
		Tester: &flute.Tester{
			Path:         "/api/dashboards",
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "title": "test",
  "description": "test"
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				dashboardBody = `{
  "creator_user_id": "admin",
  "description": "test",
  "created_at": "2020-04-24T02:14:36.078Z",
  "positions": {},
  "id": "5ea24b8c2ab79c001251ee46",
  "title": "test",
  "widgets": []
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
			BodyString: `{
  "dashboard_id": "5ea24b8c2ab79c001251ee46"
}`,
		},
	}

	createStep := resource.TestStep{
		ResourceName: "graylog_dashboard.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, postRoute)
		},
		Config: `
resource "graylog_dashboard" "test" {
  title       = "test"
  description = "test"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("graylog_dashboard.test", "title", "test"),
			resource.TestCheckResourceAttr("graylog_dashboard.test", "description", "test"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a dashboard",
		Matcher: &flute.Matcher{
			Method: "PUT",
		},
		Tester: &flute.Tester{
			Path:         "/api/dashboards/5ea24b8c2ab79c001251ee46",
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "title": "title updated",
  "description": "description updated"
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				dashboardBody = `{
  "creator_user_id": "admin",
  "description": "description updated",
  "created_at": "2020-04-24T02:14:36.078Z",
  "positions": {},
  "id": "5ea24b8c2ab79c001251ee46",
  "title": "title updated",
  "widgets": []
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 204,
			},
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a dashboard",
		Matcher: &flute.Matcher{
			Method: "DELETE",
		},
		Tester: &flute.Tester{
			Path:         "/api/dashboards/5ea24b8c2ab79c001251ee46",
			PartOfHeader: testutil.Header(),
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 204,
			},
		},
	}

	updateStep := resource.TestStep{
		ResourceName: "graylog_dashboard.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_dashboard" "test" {
  title       = "title updated"
  description = "description updated"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("graylog_dashboard.test", "title", "title updated"),
			resource.TestCheckResourceAttr("graylog_dashboard.test", "description", "description updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_dashboard", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
