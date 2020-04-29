package widget

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccDashboardWidget(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	dashboardBody := ""

	resourceName := "graylog_dashboard_widget.test"
	resourceURLPath := "/api/dashboards/5ea24b8c2ab79c001251ee46/widgets/51e7238a-f73d-4a5c-a4cb-9a91b5560c1d"

	getRoute := flute.Route{
		Name: "get a dashboard widget",
		Matcher: &flute.Matcher{
			Method: "GET",
		},
		Tester: &flute.Tester{
			Path:         resourceURLPath,
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
		Name: "create a dashboard widget",
		Matcher: &flute.Matcher{
			Method: "POST",
		},
		Tester: &flute.Tester{
			Path:         "/api/dashboards/5ea24b8c2ab79c001251ee46/widgets",
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "cache_time": 10,
  "description": "Stream search result count",
  "type": "STREAM_SEARCH_RESULT_COUNT",
  "config": {
    "timerange": {
      "type": "relative",
      "range": 400
    },
    "lower_is_better": true,
    "stream_id": "5e9989962ab79c001156f7e2",
    "trend": true,
    "query": ""
  }
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				dashboardBody = `{
  "id": "51e7238a-f73d-4a5c-a4cb-9a91b5560c1d",
  "description": "Stream search result count",
  "type": "STREAM_SEARCH_RESULT_COUNT",
  "config": {
    "timerange": {
      "type": "relative",
      "range": 400
    },
    "lower_is_better": true,
    "stream_id": "5e9989962ab79c001156f7e2",
    "trend": true,
    "query": ""
  },
  "cache_time": 10,
  "creator_user_id": "admin"
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
			BodyString: `{
  "widget_id": "51e7238a-f73d-4a5c-a4cb-9a91b5560c1d"
}`,
		},
	}

	createStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, postRoute)
		},
		Config: `
resource "graylog_dashboard_widget" "test" {
  description = "Stream search result count"
  dashboard_id = "5ea24b8c2ab79c001251ee46"
  type         = "STREAM_SEARCH_RESULT_COUNT"
  cache_time   = 10
  config       = <<EOF
{
  "timerange": {
    "type": "relative",
    "range": 400
  },
  "lower_is_better": true,
  "stream_id": "5e9989962ab79c001156f7e2",
  "trend": true,
  "query": ""
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("graylog_dashboard_widget.test", "description", "Stream search result count"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a dashboard widget",
		Matcher: &flute.Matcher{
			Method: "PUT",
		},
		Tester: &flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "cache_time": 20,
  "description": "Stream search result count updated",
  "type": "STREAM_SEARCH_RESULT_COUNT",
  "config": {
    "timerange": {
      "type": "relative",
      "range": 500
    },
    "lower_is_better": true,
    "stream_id": "5e9989962ab79c001156f7e2",
    "trend": true,
    "query": ""
  }
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				dashboardBody = `{
  "id": "51e7238a-f73d-4a5c-a4cb-9a91b5560c1d",
  "description": "Stream search result count updated",
  "type": "STREAM_SEARCH_RESULT_COUNT",
  "config": {
    "timerange": {
      "type": "relative",
      "range": 500
    },
    "lower_is_better": true,
    "stream_id": "5e9989962ab79c001156f7e2",
    "trend": true,
    "query": ""
  },
  "cache_time": 20,
  "creator_user_id": "admin"
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
		Name: "delete a dashboard widget",
		Matcher: &flute.Matcher{
			Method: "DELETE",
		},
		Tester: &flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
		},
		Response: &flute.Response{
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
resource "graylog_dashboard_widget" "test" {
  description = "Stream search result count updated"
  dashboard_id = "5ea24b8c2ab79c001251ee46"
  type         = "STREAM_SEARCH_RESULT_COUNT"
  cache_time   = 20
  config       = <<EOF
{
  "timerange": {
    "type": "relative",
    "range": 500
  },
  "lower_is_better": true,
  "stream_id": "5e9989962ab79c001156f7e2",
  "trend": true,
  "query": ""
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("graylog_dashboard_widget.test", "description", "Stream search result count updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_dashboard_widget", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
