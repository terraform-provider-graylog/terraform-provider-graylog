package indexset

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccIndexSet(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	getRoute := flute.Route{
		Name: "get indexsets",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         "/api/system/indices/index_sets",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{
  "total": 3,
  "index_sets": [
    {
      "id": "5ea81cb42ab79c00129dbe58",
      "title": "Default index set",
      "description": "The Graylog default index set",
      "index_prefix": "graylog",
      "shards": 4,
      "replicas": 0,
      "rotation_strategy_class": "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy",
      "rotation_strategy": {
        "type": "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
        "max_docs_per_index": 20000000
      },
      "retention_strategy_class": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
      "retention_strategy": {
        "type": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
        "max_number_of_indices": 20
      },
      "creation_date": "2020-04-28T12:08:20.994Z",
      "index_analyzer": "standard",
      "index_optimization_max_num_segments": 1,
      "index_optimization_disabled": false,
      "field_type_refresh_interval": 5000,
      "writable": true,
      "default": true
    },
    {
      "id": "5ea81cc02ab79c00129dbf1c",
      "title": "Graylog Events",
      "description": "Stores Graylog events.",
      "index_prefix": "gl-events",
      "shards": 4,
      "replicas": 0,
      "rotation_strategy_class": "org.graylog2.indexer.rotation.strategies.TimeBasedRotationStrategy",
      "rotation_strategy": {
        "type": "org.graylog2.indexer.rotation.strategies.TimeBasedRotationStrategyConfig",
        "rotation_period": "P1M"
      },
      "retention_strategy_class": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
      "retention_strategy": {
        "type": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
        "max_number_of_indices": 12
      },
      "creation_date": "2020-04-28T12:08:32.18Z",
      "index_analyzer": "standard",
      "index_optimization_max_num_segments": 1,
      "index_optimization_disabled": false,
      "field_type_refresh_interval": 60000,
      "writable": true,
      "default": false
    },
    {
      "id": "5ea81cc02ab79c00129dbf1f",
      "title": "Graylog System Events",
      "description": "Stores Graylog system events.",
      "index_prefix": "gl-system-events",
      "shards": 4,
      "replicas": 0,
      "rotation_strategy_class": "org.graylog2.indexer.rotation.strategies.TimeBasedRotationStrategy",
      "rotation_strategy": {
        "type": "org.graylog2.indexer.rotation.strategies.TimeBasedRotationStrategyConfig",
        "rotation_period": "P1M"
      },
      "retention_strategy_class": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
      "retention_strategy": {
        "type": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
        "max_number_of_indices": 12
      },
      "creation_date": "2020-04-28T12:08:32.188Z",
      "index_analyzer": "standard",
      "index_optimization_max_num_segments": 1,
      "index_optimization_disabled": false,
      "field_type_refresh_interval": 60000,
      "writable": true,
      "default": false
    }
  ],
  "stats": {}
}`)),
				}, nil
			},
		},
	}

	readStep := resource.TestStep{
		ResourceName: "data.graylog_index_set.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute)
		},
		Config: `
data "graylog_index_set" "test" {
  title = "Default index set"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.graylog_index_set.test", "title", "Default index set"),
			resource.TestCheckResourceAttr("data.graylog_index_set.test", "description", "The Graylog default index set"),
		),
	}

	readStepByPrefix := resource.TestStep{
		ResourceName: "data.graylog_index_set.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute)
		},
		Config: `
data "graylog_index_set" "test" {
  index_prefix = "graylog"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.graylog_index_set.test", "title", "Default index set"),
			resource.TestCheckResourceAttr("data.graylog_index_set.test", "description", "The Graylog default index set"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleDataSourceProviders("graylog_index_set", DataSource()),
		Steps: []resource.TestStep{
			readStep,
			readStepByPrefix,
		},
	})
}

func TestAccIndexSet_byIndexSetID(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	getRoute := flute.Route{
		Name: "get indexset",
		Matcher: flute.Matcher{
			Method: "GET",
		},
		Tester: flute.Tester{
			Path:         "/api/system/indices/index_sets/5ea81cc02ab79c00129dbf1c",
			PartOfHeader: testutil.Header(),
		},
		Response: flute.Response{
			Response: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(strings.NewReader(`{
  "id": "5ea81cc02ab79c00129dbf1c",
  "title": "Graylog Events",
  "description": "Stores Graylog events.",
  "index_prefix": "gl-events",
  "shards": 4,
  "replicas": 0,
  "rotation_strategy_class": "org.graylog2.indexer.rotation.strategies.TimeBasedRotationStrategy",
  "rotation_strategy": {
    "type": "org.graylog2.indexer.rotation.strategies.TimeBasedRotationStrategyConfig",
    "rotation_period": "P1M"
  },
  "retention_strategy_class": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
  "retention_strategy": {
    "type": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
    "max_number_of_indices": 12
  },
  "creation_date": "2020-04-28T12:08:32.18Z",
  "index_analyzer": "standard",
  "index_optimization_max_num_segments": 1,
  "index_optimization_disabled": false,
  "field_type_refresh_interval": 60000,
  "writable": true,
  "default": false
}`)),
				}, nil
			},
		},
	}

	readStep := resource.TestStep{
		ResourceName: "data.graylog_index_set.test",
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute)
		},
		Config: `
data "graylog_index_set" "test" {
  index_set_id = "5ea81cc02ab79c00129dbf1c"
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.graylog_index_set.test", "title", "Graylog Events"),
			resource.TestCheckResourceAttr("data.graylog_index_set.test", "description", "Stores Graylog events."),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleDataSourceProviders("graylog_index_set", DataSource()),
		Steps: []resource.TestStep{
			readStep,
		},
	})
}
