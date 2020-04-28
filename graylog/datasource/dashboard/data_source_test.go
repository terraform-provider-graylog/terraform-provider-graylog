package dashboard

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccDashboard(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	getRoute := flute.Route{
		Name: "get dashboards",
		Matcher: &flute.Matcher{
			Method: "GET",
		},
		Tester: &flute.Tester{
			Path:         "/api/dashboards",
			PartOfHeader: testutil.Header(),
		},
		Response: &flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{
  "total": 2,
  "dashboards": [
    {
      "creator_user_id": "admin",
      "description": "zoo",
      "created_at": "2020-04-28T13:36:27.559Z",
      "positions": {},
      "id": "5ea8315b2ab79c00129dcba2",
      "title": "zoo",
      "widgets": []
    },
    {
      "creator_user_id": "admin",
      "description": "test description",
      "created_at": "2020-04-28T13:36:27.559Z",
      "positions": {},
      "id": "5ea8315b2ab79c00129dcba2",
      "title": "test",
      "widgets": []
    }
  ]
}`)),
				}, nil
			},
		},
	}

	readStep := resource.TestStep{
		ResourceName: "data.graylog_dashboard.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute)
		},
		Config: `
data "graylog_dashboard" "test" {
  title = "test"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.graylog_dashboard.test", "title", "test"),
			resource.TestCheckResourceAttr("data.graylog_dashboard.test", "description", "test description"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleDataSourceProviders("graylog_dashboard", DataSource()),
		Steps: []resource.TestStep{
			readStep,
		},
	})
}
