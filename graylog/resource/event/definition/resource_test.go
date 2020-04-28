package definition

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/suzuki-shunsuke/terraform-provider-graylog/graylog/testutil"
)

func TestAccAmarmCallback(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	definitionBody := ""

	postURLPath := "/api/events/definitions"
	resourceURLPath := postURLPath + "/5ea3c8b42ab79c00127570c4"
	resourceName := "graylog_event_definition.test"

	getRoute := flute.Route{
		Name: "get a event definition",
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
					Body:       ioutil.NopCloser(strings.NewReader(definitionBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a event definition",
		Matcher: &flute.Matcher{
			Method: "POST",
		},
		Tester: &flute.Tester{
			Path:         postURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "title": "new-event-definition",
  "description": "",
  "priority": 1,
  "alert": true,
  "config": {
    "type": "aggregation-v1",
    "query": "test",
    "query_parameters": [],
    "streams": [
      "5e9989962ab79c001156f7e2"
    ],
    "group_by": [],
    "series": [],
    "conditions": null,
    "search_within_ms": 60000,
    "execute_every_ms": 60000
  },
  "field_spec": {
    "test": {
      "data_type": "string",
      "providers": [
        {
          "type": "template-v1",
          "template": "test",
          "require_values": false
        }
      ]
    }
  },
  "key_spec": [
    "test"
  ],
  "notification_settings": {
    "grace_period_ms": 0,
    "backlog_size": 0
  },
  "notifications": [
    {
      "notification_id": "5ea3c1d72ab79c00127567fe"
    }
  ]
}
`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				definitionBody = `{
  "id": "5ea3c8b42ab79c00127570c4",
  "title": "new-event-definition",
  "description": "",
  "priority": 1,
  "alert": true,
  "config": {
    "type": "aggregation-v1",
    "query": "test",
    "query_parameters": [],
    "streams": [
      "5e9989962ab79c001156f7e2"
    ],
    "group_by": [],
    "series": [],
    "conditions": null,
    "search_within_ms": 60000,
    "execute_every_ms": 60000
  },
  "field_spec": {
    "test": {
      "data_type": "string",
      "providers": [
        {
          "type": "template-v1",
          "template": "test",
          "require_values": false
        }
      ]
    }
  },
  "key_spec": [
    "test"
  ],
  "notification_settings": {
    "grace_period_ms": 0,
    "backlog_size": 0
  },
  "notifications": [
    {
      "notification_id": "5ea3c1d72ab79c00127567fe",
      "notification_parameters": null
    }
  ],
  "storage": [
    {
      "type": "persist-to-streams-v1",
      "streams": [
        "000000000000000000000002"
      ]
    }
  ]
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "id": "5ea3c8b42ab79c00127570c4",
  "title": "new-event-definition",
  "description": "",
  "priority": 1,
  "alert": true,
  "config": {
    "type": "aggregation-v1",
    "query": "test",
    "query_parameters": [],
    "streams": [
      "5e9989962ab79c001156f7e2"
    ],
    "group_by": [],
    "series": [],
    "conditions": null,
    "search_within_ms": 60000,
    "execute_every_ms": 60000
  },
  "field_spec": {
    "test": {
      "data_type": "string",
      "providers": [
        {
          "type": "template-v1",
          "template": "test",
          "require_values": false
        }
      ]
    }
  },
  "key_spec": [
    "test"
  ],
  "notification_settings": {
    "grace_period_ms": 0,
    "backlog_size": 0
  },
  "notifications": [
    {
      "notification_id": "5ea3c1d72ab79c00127567fe",
      "notification_parameters": null
    }
  ],
  "storage": [
    {
      "type": "persist-to-streams-v1",
      "streams": [
        "000000000000000000000002"
      ]
    }
  ]
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a event definition",
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

	createStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, postRoute, deleteRoute)
		},
		Config: `
resource "graylog_event_definition" "test" {
  title       = "new-event-definition"
  description = ""
  priority    = 1
  alert       = true
  config      = <<EOF
{
  "type": "aggregation-v1",
  "query": "test",
	"query_parameters": [],
  "streams": [
    "5e9989962ab79c001156f7e2"
  ],
  "search_within_ms": 60000,
  "execute_every_ms": 60000,
  "group_by": [],
  "series": [],
  "conditions": null
}
EOF
  field_spec  = <<EOF
{
  "test": {
    "data_type": "string",
    "providers": [
      {
        "type": "template-v1",
        "template": "test",
        "require_values": false
      }
    ]
  }
}
EOF

  key_spec = ["test"]

  notification_settings {
    grace_period_ms = 0
    backlog_size    = 0
  }

  notifications {
    notification_id = "5ea3c1d72ab79c00127567fe"
  }
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "title", "new-event-definition"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a event definition",
		Matcher: &flute.Matcher{
			Method: "PUT",
		},
		Tester: &flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "id": "5ea3c8b42ab79c00127570c4",
  "title": "new-event-definition",
  "description": "updated",
  "priority": 1,
  "alert": true,
  "config": {
    "type": "aggregation-v1",
    "query": "test",
    "query_parameters": [],
    "streams": [
      "5e9989962ab79c001156f7e2"
    ],
    "group_by": [],
    "series": [],
    "conditions": null,
    "search_within_ms": 60000,
    "execute_every_ms": 60000
  },
  "field_spec": {
    "test": {
      "data_type": "string",
      "providers": [
        {
          "type": "template-v1",
          "template": "test",
          "require_values": false
        }
      ]
    }
  },
  "key_spec": [
    "test"
  ],
  "notification_settings": {
    "grace_period_ms": 0,
    "backlog_size": 0
  },
  "notifications": [
    {
      "notification_id": "5ea3c1d72ab79c00127567fe"
    }
  ]
}`,
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				definitionBody = `{
  "id": "5ea3c8b42ab79c00127570c4",
  "title": "new-event-definition",
  "description": "updated",
  "priority": 1,
  "alert": true,
  "config": {
    "type": "aggregation-v1",
    "query": "test",
    "query_parameters": [],
    "streams": [
      "5e9989962ab79c001156f7e2"
    ],
    "group_by": [],
    "series": [],
    "conditions": null,
    "search_within_ms": 60000,
    "execute_every_ms": 60000
  },
  "field_spec": {
    "test": {
      "data_type": "string",
      "providers": [
        {
          "type": "template-v1",
          "template": "test",
          "require_values": false
        }
      ]
    }
  },
  "key_spec": [
    "test"
  ],
  "notification_settings": {
    "grace_period_ms": 0,
    "backlog_size": 0
  },
  "notifications": [
    {
      "notification_id": "5ea3c1d72ab79c00127567fe",
      "notification_parameters": null
    }
  ],
  "storage": [
    {
      "type": "persist-to-streams-v1",
      "streams": [
        "000000000000000000000002"
      ]
    }
  ]
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "id": "5ea3c8b42ab79c00127570c4",
  "title": "new-event-definition",
  "description": "updated",
  "priority": 1,
  "alert": true,
  "config": {
    "type": "aggregation-v1",
    "query": "test",
    "query_parameters": [],
    "streams": [
      "5e9989962ab79c001156f7e2"
    ],
    "group_by": [],
    "series": [],
    "conditions": null,
    "search_within_ms": 60000,
    "execute_every_ms": 60000
  },
  "field_spec": {
    "test": {
      "data_type": "string",
      "providers": [
        {
          "type": "template-v1",
          "template": "test",
          "require_values": false
        }
      ]
    }
  },
  "key_spec": [
    "test"
  ],
  "notification_settings": {
    "grace_period_ms": 0,
    "backlog_size": 0
  },
  "notifications": [
    {
      "notification_id": "5ea3c1d72ab79c00127567fe",
      "notification_parameters": null
    }
  ],
  "storage": [
    {
      "type": "persist-to-streams-v1",
      "streams": [
        "000000000000000000000002"
      ]
    }
  ]
}`,
		},
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_event_definition" "test" {
  title       = "new-event-definition"
  description = "updated"
  priority    = 1
  alert       = true
  config      = <<EOF
{
  "type": "aggregation-v1",
  "query": "test",
	"query_parameters": [],
  "streams": [
    "5e9989962ab79c001156f7e2"
  ],
  "search_within_ms": 60000,
  "execute_every_ms": 60000,
  "group_by": [],
  "series": [],
  "conditions": null
}
EOF
  field_spec  = <<EOF
{
  "test": {
    "data_type": "string",
    "providers": [
      {
        "type": "template-v1",
        "template": "test",
        "require_values": false
      }
    ]
  }
}
EOF

  key_spec = ["test"]

  notification_settings {
    grace_period_ms = 0
    backlog_size    = 0
  }

  notifications {
    notification_id = "5ea3c1d72ab79c00127567fe"
  }
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "description", "updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_event_definition", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
