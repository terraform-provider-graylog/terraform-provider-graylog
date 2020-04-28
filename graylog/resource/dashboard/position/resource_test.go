package position

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccDashboardWidgetPositions(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	conditionBody := ""

	dashboardURLPath := "/api/dashboards/5ea24b8c2ab79c001251ee46"
	resourceURLPath := dashboardURLPath + "/positions"
	resourceName := "graylog_dashboard_widget_positions.test"

	getRoute := flute.Route{
		Name: "get a dashboard widget position",
		Matcher: &flute.Matcher{
			Method: "GET",
		},
		Tester: &flute.Tester{
			Path:         dashboardURLPath,
			PartOfHeader: testutil.Header(),
		},
		Response: &flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(conditionBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a dashboard widget position",
		Matcher: &flute.Matcher{
			Method: "PUT",
		},
		Tester: &flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "positions": [
    {
		  "id": "5ef230f3-cf95-4be5-bd2f-eb3a5a64ac65",
      "width": 2,
      "col": 1,
      "row": 0,
      "height": 2
    },
    {
      "id": "e19ef574-173b-4159-ab87-4b9b5bb0aa27",
      "width": 1,
      "col": 0,
      "row": 0,
      "height": 1
    }
  ]
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				conditionBody = `{
  "creator_user_id": "admin",
  "description": "description",
  "created_at": "2020-04-24T02:14:36.078Z",
  "positions": {
    "5ef230f3-cf95-4be5-bd2f-eb3a5a64ac65": {
      "width": 2,
      "col": 1,
      "row": 0,
      "height": 2
    },
    "e19ef574-173b-4159-ab87-4b9b5bb0aa27": {
      "width": 1,
      "col": 0,
      "row": 0,
      "height": 1
    }
  },
  "id": "5ea24b8c2ab79c001251ee46",
  "title": "title",
  "widgets": [
    {
      "creator_user_id": "admin",
      "cache_time": 20,
      "description": "Stream search result count",
      "id": "51e7238a-f73d-4a5c-a4cb-9a91b5560c1d",
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
    }
  ]
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 204,
			},
		},
	}

	createStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, postRoute)
		},
		Config: `
resource "graylog_dashboard_widget_positions" "test" {
  dashboard_id = "5ea24b8c2ab79c001251ee46"

	positions = <<EOF
{
  "5ef230f3-cf95-4be5-bd2f-eb3a5a64ac65": {
    "width": 2,
    "col": 1,
    "row": 0,
    "height": 2
  },
  "e19ef574-173b-4159-ab87-4b9b5bb0aa27": {
    "width": 1,
    "col": 0,
    "row": 0,
    "height": 1
  }
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "dashboard_id", "5ea24b8c2ab79c001251ee46"),
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
  "positions": [
    {
		  "id": "5ef230f3-cf95-4be5-bd2f-eb3a5a64ac65",
      "width": 2,
      "col": 2,
      "row": 0,
      "height": 2
    },
    {
      "id": "e19ef574-173b-4159-ab87-4b9b5bb0aa27",
      "width": 1,
      "col": 1,
      "row": 1,
      "height": 1
    }
  ]
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				conditionBody = `{
  "creator_user_id": "admin",
  "description": "description",
  "created_at": "2020-04-24T02:14:36.078Z",
  "positions": {
    "5ef230f3-cf95-4be5-bd2f-eb3a5a64ac65": {
      "width": 2,
      "col": 2,
      "row": 0,
      "height": 2
    },
    "e19ef574-173b-4159-ab87-4b9b5bb0aa27": {
      "width": 1,
      "col": 1,
      "row": 1,
      "height": 1
    }
  },
  "id": "5ea24b8c2ab79c001251ee46",
  "title": "title",
  "widgets": [
    {
      "creator_user_id": "admin",
      "cache_time": 20,
      "description": "Stream search result count",
      "id": "51e7238a-f73d-4a5c-a4cb-9a91b5560c1d",
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
    }
  ]
}`
			},
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
			testutil.SetHTTPClient(t, getRoute, updateRoute)
		},
		Config: `
resource "graylog_dashboard_widget_positions" "test" {
  dashboard_id = "5ea24b8c2ab79c001251ee46"

	positions = <<EOF
{
  "5ef230f3-cf95-4be5-bd2f-eb3a5a64ac65": {
    "width": 2,
    "col": 2,
    "row": 0,
    "height": 2
  },
  "e19ef574-173b-4159-ab87-4b9b5bb0aa27": {
    "width": 1,
    "col": 1,
    "row": 1,
    "height": 1
  }
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "dashboard_id", "5ea24b8c2ab79c001251ee46"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_dashboard_widget_positions", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
