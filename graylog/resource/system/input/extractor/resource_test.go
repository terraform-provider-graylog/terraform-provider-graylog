package extractor

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccExtractor(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	conditionBody := ""

	postURLPath := "/api/system/inputs/5e9989952ab79c001156f7d2/extractors"
	resourceURLPath := postURLPath + "/553e37b0-86a8-11ea-a7d4-0242ac120004"
	resourceName := "graylog_extractor.test_json"

	getRoute := flute.Route{
		Name: "get a extractor",
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
					Body:       ioutil.NopCloser(strings.NewReader(conditionBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a extractor",
		Matcher: flute.Matcher{
			Method: "POST",
		},
		Tester: flute.Tester{
			Path:         postURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "title": "test",
  "extractor_type": "json",
  "converters": {},
  "order": 0,
  "cut_or_copy": "copy",
  "source_field": "message",
  "target_field": "none",
  "extractor_config": {
    "list_separator": ", ",
    "kv_separator": "=",
    "key_prefix": "visit_",
    "key_separator": "_",
    "replace_key_whitespace": false,
    "key_whitespace_replacement": "_"
  },
  "condition_type": "none",
  "condition_value": ""
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				conditionBody = `{
  "id": "553e37b0-86a8-11ea-a7d4-0242ac120004",
  "title": "test",
  "type": "json",
  "converters": [],
  "order": 0,
  "exceptions": 0,
  "metrics": {
    "total": {
      "time": {
        "min": 0,
        "max": 0,
        "mean": 0,
        "std_dev": 0,
        "95th_percentile": 0,
        "98th_percentile": 0,
        "99th_percentile": 0
      },
      "rate": {
        "total": 0,
        "mean": 0,
        "one_minute": 0,
        "five_minute": 0,
        "fifteen_minute": 0
      },
      "duration_unit": "microseconds",
      "rate_unit": "events/second"
    },
    "condition": {
      "time": {
        "min": 0,
        "max": 0,
        "mean": 0,
        "std_dev": 0,
        "95th_percentile": 0,
        "98th_percentile": 0,
        "99th_percentile": 0
      },
      "rate": {
        "total": 0,
        "mean": 0,
        "one_minute": 0,
        "five_minute": 0,
        "fifteen_minute": 0
      },
      "duration_unit": "microseconds",
      "rate_unit": "events/second"
    },
    "execution": {
      "time": {
        "min": 0,
        "max": 0,
        "mean": 0,
        "std_dev": 0,
        "95th_percentile": 0,
        "98th_percentile": 0,
        "99th_percentile": 0
      },
      "rate": {
        "total": 0,
        "mean": 0,
        "one_minute": 0,
        "five_minute": 0,
        "fifteen_minute": 0
      },
      "duration_unit": "microseconds",
      "rate_unit": "events/second"
    },
    "converters": {
      "time": {
        "min": 0,
        "max": 0,
        "mean": 0,
        "std_dev": 0,
        "95th_percentile": 0,
        "98th_percentile": 0,
        "99th_percentile": 0
      },
      "rate": {
        "total": 0,
        "mean": 0,
        "one_minute": 0,
        "five_minute": 0,
        "fifteen_minute": 0
      },
      "duration_unit": "microseconds",
      "rate_unit": "events/second"
    },
    "condition_hits": 0,
    "condition_misses": 0
  },
  "cursor_strategy": "copy",
  "source_field": "message",
  "target_field": "none",
  "extractor_config": {
    "list_separator": ", ",
    "kv_separator": "=",
    "key_prefix": "visit_",
    "key_separator": "_",
    "replace_key_whitespace": false,
    "key_whitespace_replacement": "_"
  },
  "creator_user_id": "admin",
  "condition_type": "none",
  "condition_value": "",
  "converter_exceptions": 0
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
			BodyString: `{
  "extractor_id": "553e37b0-86a8-11ea-a7d4-0242ac120004"
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a extractor",
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
resource "graylog_extractor" "test_json" {
  input_id        = "5e9989952ab79c001156f7d2"
  title           = "test"
  type            = "json"
  cursor_strategy = "copy"
  source_field    = "message"
  target_field    = "none"
  condition_type  = "none"
  condition_value = ""
  order           = 0

  extractor_config = <<EOF
{
  "list_separator": ", ",
  "kv_separator": "=",
  "key_prefix": "visit_",
  "key_separator": "_",
  "replace_key_whitespace": false,
  "key_whitespace_replacement": "_"
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "title", "test"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a extractor",
		Matcher: flute.Matcher{
			Method: "PUT",
		},
		Tester: flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "title": "test updated",
  "extractor_type": "json",
  "converters": {},
  "order": 0,
  "cut_or_copy": "copy",
  "source_field": "message",
  "target_field": "none",
  "extractor_config": {
    "list_separator": ", ",
    "kv_separator": "=",
    "key_prefix": "visit_updated_",
    "key_separator": "_",
    "replace_key_whitespace": false,
    "key_whitespace_replacement": "_"
  },
  "condition_type": "none",
  "condition_value": ""
}
`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				conditionBody = `{
  "id": "553e37b0-86a8-11ea-a7d4-0242ac120004",
  "title": "test updated",
  "type": "json",
  "converters": [],
  "order": 0,
  "exceptions": 0,
  "metrics": {
    "total": {
      "time": {
        "min": 0,
        "max": 0,
        "mean": 0,
        "std_dev": 0,
        "95th_percentile": 0,
        "98th_percentile": 0,
        "99th_percentile": 0
      },
      "rate": {
        "total": 0,
        "mean": 0,
        "one_minute": 0,
        "five_minute": 0,
        "fifteen_minute": 0
      },
      "duration_unit": "microseconds",
      "rate_unit": "events/second"
    },
    "condition": {
      "time": {
        "min": 0,
        "max": 0,
        "mean": 0,
        "std_dev": 0,
        "95th_percentile": 0,
        "98th_percentile": 0,
        "99th_percentile": 0
      },
      "rate": {
        "total": 0,
        "mean": 0,
        "one_minute": 0,
        "five_minute": 0,
        "fifteen_minute": 0
      },
      "duration_unit": "microseconds",
      "rate_unit": "events/second"
    },
    "execution": {
      "time": {
        "min": 0,
        "max": 0,
        "mean": 0,
        "std_dev": 0,
        "95th_percentile": 0,
        "98th_percentile": 0,
        "99th_percentile": 0
      },
      "rate": {
        "total": 0,
        "mean": 0,
        "one_minute": 0,
        "five_minute": 0,
        "fifteen_minute": 0
      },
      "duration_unit": "microseconds",
      "rate_unit": "events/second"
    },
    "converters": {
      "time": {
        "min": 0,
        "max": 0,
        "mean": 0,
        "std_dev": 0,
        "95th_percentile": 0,
        "98th_percentile": 0,
        "99th_percentile": 0
      },
      "rate": {
        "total": 0,
        "mean": 0,
        "one_minute": 0,
        "five_minute": 0,
        "fifteen_minute": 0
      },
      "duration_unit": "microseconds",
      "rate_unit": "events/second"
    },
    "condition_hits": 0,
    "condition_misses": 0
  },
  "cursor_strategy": "copy",
  "source_field": "message",
  "target_field": "none",
  "extractor_config": {
    "list_separator": ", ",
    "kv_separator": "=",
    "key_prefix": "visit_updated_",
    "key_separator": "_",
    "replace_key_whitespace": false,
    "key_whitespace_replacement": "_"
  },
  "creator_user_id": "admin",
  "condition_type": "none",
  "condition_value": "",
  "converter_exceptions": 0
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "id": "553e37b0-86a8-11ea-a7d4-0242ac120004",
  "title": "test updated",
  "type": "json",
  "converters": [],
  "order": 0,
  "exceptions": 0,
  "metrics": {
    "total": {
      "time": {
        "min": 0,
        "max": 0,
        "mean": 0,
        "std_dev": 0,
        "95th_percentile": 0,
        "98th_percentile": 0,
        "99th_percentile": 0
      },
      "rate": {
        "total": 0,
        "mean": 0,
        "one_minute": 0,
        "five_minute": 0,
        "fifteen_minute": 0
      },
      "duration_unit": "microseconds",
      "rate_unit": "events/second"
    },
    "condition": {
      "time": {
        "min": 0,
        "max": 0,
        "mean": 0,
        "std_dev": 0,
        "95th_percentile": 0,
        "98th_percentile": 0,
        "99th_percentile": 0
      },
      "rate": {
        "total": 0,
        "mean": 0,
        "one_minute": 0,
        "five_minute": 0,
        "fifteen_minute": 0
      },
      "duration_unit": "microseconds",
      "rate_unit": "events/second"
    },
    "execution": {
      "time": {
        "min": 0,
        "max": 0,
        "mean": 0,
        "std_dev": 0,
        "95th_percentile": 0,
        "98th_percentile": 0,
        "99th_percentile": 0
      },
      "rate": {
        "total": 0,
        "mean": 0,
        "one_minute": 0,
        "five_minute": 0,
        "fifteen_minute": 0
      },
      "duration_unit": "microseconds",
      "rate_unit": "events/second"
    },
    "converters": {
      "time": {
        "min": 0,
        "max": 0,
        "mean": 0,
        "std_dev": 0,
        "95th_percentile": 0,
        "98th_percentile": 0,
        "99th_percentile": 0
      },
      "rate": {
        "total": 0,
        "mean": 0,
        "one_minute": 0,
        "five_minute": 0,
        "fifteen_minute": 0
      },
      "duration_unit": "microseconds",
      "rate_unit": "events/second"
    },
    "condition_hits": 0,
    "condition_misses": 0
  },
  "cursor_strategy": "copy",
  "source_field": "message",
  "target_field": "none",
  "extractor_config": {
    "list_separator": ", ",
    "kv_separator": "=",
    "key_prefix": "visit_updated_",
    "key_separator": "_",
    "replace_key_whitespace": false,
    "key_whitespace_replacement": "_"
  },
  "creator_user_id": "admin",
  "condition_type": "none",
  "condition_value": "",
  "converter_exceptions": 0
}`,
		},
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_extractor" "test_json" {
  input_id        = "5e9989952ab79c001156f7d2"
  title           = "test updated"
  type            = "json"
  cursor_strategy = "copy"
  source_field    = "message"
  target_field    = "none"
  condition_type  = "none"
  condition_value = ""
  order           = 0

  extractor_config = <<EOF
{
  "list_separator": ", ",
  "kv_separator": "=",
  "key_prefix": "visit_updated_",
  "key_separator": "_",
  "replace_key_whitespace": false,
  "key_whitespace_replacement": "_"
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "title", "test updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_extractor", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
