package indexset

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccIndexSet(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	indexsetBody := ""

	resourceURLPath := "/api/system/indices/index_sets/5ea25a282ab79c00125200b9"
	resourceName := "graylog_index_set.test"

	getRoute := flute.Route{
		Name: "get a indexset",
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
					Body:       ioutil.NopCloser(strings.NewReader(indexsetBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a indexset",
		Matcher: &flute.Matcher{
			Method: "POST",
		},
		Tester: &flute.Tester{
			Path:         "/api/system/indices/index_sets",
			PartOfHeader: testutil.Header(),
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				body := map[string]interface{}{}
				if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
					t.Fatal(err)
				}
				keys := []string{
					"title", "description", "index_prefix", "shards", "replicas",
					"rotation_strategy", "rotation_strategy_class",
					"retention_strategy", "retention_strategy_class",
					"index_analyzer", "index_optimization_max_num_segments",
					"index_optimization_disabled", "field_type_refresh_interval",
					"creation_date", "writable",
				}
				if err := testutil.EqualMapKeys(body, keys...); err != nil {
					t.Fatal(err)
				}
				require.Equal(t, "test", body["title"])
				require.Equal(t, "test", body["description"])
				require.Equal(t, 4.0, body["shards"])
				require.Equal(t, 1.0, body["replicas"])
				require.Equal(t, "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy", body["rotation_strategy_class"])
				require.Equal(t, "terraform-provider-graylog-test", body["index_prefix"])
				require.Equal(t, map[string]interface{}{
					"type":               "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
					"max_docs_per_index": 30000000.0,
				}, body["rotation_strategy"])

				indexsetBody = `{
  "id": "5ea25a282ab79c00125200b9",
  "title": "test",
  "description": "test",
  "index_prefix": "terraform-provider-graylog-test",
  "shards": 4,
  "replicas": 1,
  "rotation_strategy_class": "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy",
  "rotation_strategy": {
    "type": "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
    "max_docs_per_index": 30000000
  },
  "retention_strategy_class": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
  "retention_strategy": {
    "type": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
    "max_number_of_indices": 30
  },
  "creation_date": "2020-04-21T12:56:41.97Z",
  "index_analyzer": "standard",
  "index_optimization_max_num_segments": 1,
  "index_optimization_disabled": true,
  "field_type_refresh_interval": 5000,
  "writable": true,
  "default": false
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 200,
			},
			BodyString: `{
  "id": "5ea25a282ab79c00125200b9",
  "title": "test",
  "description": "test",
  "index_prefix": "terraform-provider-graylog-test",
  "shards": 4,
  "replicas": 1,
  "rotation_strategy_class": "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy",
  "rotation_strategy": {
    "type": "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
    "max_docs_per_index": 30000000
  },
  "retention_strategy_class": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
  "retention_strategy": {
    "type": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
    "max_number_of_indices": 30
  },
  "creation_date": "2020-04-21T12:56:41.97Z",
  "index_analyzer": "standard",
  "index_optimization_max_num_segments": 1,
  "index_optimization_disabled": true,
  "field_type_refresh_interval": 5000,
  "writable": true,
  "default": false
}`,
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a indexset",
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
resource "graylog_index_set" "test" {
  title                               = "test"
  index_prefix                        = "terraform-provider-graylog-test"
  rotation_strategy_class             = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy"
  retention_strategy_class            = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy"
  description                         = "test"
  index_analyzer                      = "standard"
  index_optimization_disabled         = true
  writable                            = true
  shards                              = 4
  replicas                            = 1
  index_optimization_max_num_segments = 1
  field_type_refresh_interval         = 5000

  retention_strategy = <<EOF
{
  "max_number_of_indices": 30,
  "type": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig"
}
EOF

  rotation_strategy = <<EOF
{
  "max_docs_per_index": 30000000,
  "type": "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig"
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "title", "test"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a indexset",
		Matcher: &flute.Matcher{
			Method: "PUT",
		},
		Tester: &flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			Test: func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
				body := map[string]interface{}{}
				if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
					t.Fatal(err)
				}
				keys := []string{
					"title", "description", "index_prefix", "shards", "replicas",
					"rotation_strategy", "rotation_strategy_class",
					"retention_strategy", "retention_strategy_class",
					"index_analyzer", "index_optimization_max_num_segments",
					"index_optimization_disabled", "field_type_refresh_interval",
					"creation_date", "writable",
				}
				if err := testutil.EqualMapKeys(body, keys...); err != nil {
					t.Fatal(err)
				}
				require.Equal(t, "test updated", body["title"])
				require.Equal(t, "test", body["description"])
				require.Equal(t, 5.0, body["shards"])
				require.Equal(t, 1.0, body["replicas"])
				require.Equal(t, "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy", body["rotation_strategy_class"])
				require.Equal(t, "terraform-provider-graylog-test", body["index_prefix"])
				require.Equal(t, map[string]interface{}{
					"type":               "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
					"max_docs_per_index": 40000000.0,
				}, body["rotation_strategy"])

				indexsetBody = `{
  "id": "5ea25a282ab79c00125200b9",
  "title": "test updated",
  "description": "test",
  "index_prefix": "terraform-provider-graylog-test",
  "shards": 5,
  "replicas": 1,
  "rotation_strategy_class": "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy",
  "rotation_strategy": {
    "type": "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
    "max_docs_per_index": 40000000
  },
  "retention_strategy_class": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
  "retention_strategy": {
    "type": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
    "max_number_of_indices": 40
  },
  "creation_date": "2020-04-21T12:56:41.97Z",
  "index_analyzer": "standard",
  "index_optimization_max_num_segments": 1,
  "index_optimization_disabled": true,
  "field_type_refresh_interval": 5000,
  "writable": true,
  "default": false
}`
			},
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 201,
			},
			BodyString: `{
  "id": "5ea25a282ab79c00125200b9",
  "title": "test updated",
  "description": "test",
  "index_prefix": "terraform-provider-graylog-test",
  "shards": 5,
  "replicas": 1,
  "rotation_strategy_class": "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy",
  "rotation_strategy": {
    "type": "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
    "max_docs_per_index": 40000000
  },
  "retention_strategy_class": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
  "retention_strategy": {
    "type": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
    "max_number_of_indices": 40
  },
  "creation_date": "2020-04-21T12:56:41.97Z",
  "index_analyzer": "standard",
  "index_optimization_max_num_segments": 1,
  "index_optimization_disabled": true,
  "field_type_refresh_interval": 5000,
  "writable": true,
  "default": false
}`,
		},
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_index_set" "test" {
  title                               = "test updated"
  index_prefix                        = "terraform-provider-graylog-test"
  rotation_strategy_class             = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy"
  retention_strategy_class            = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy"
  description                         = "test"
  index_analyzer                      = "standard"
  index_optimization_disabled         = true
  writable                            = true
  shards                              = 5
  replicas                            = 1
  index_optimization_max_num_segments = 1
  field_type_refresh_interval         = 5000

  retention_strategy = <<EOF
{
  "max_number_of_indices": 40,
  "type": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig"
}
EOF

  rotation_strategy = <<EOF
{
  "max_docs_per_index": 40000000,
  "type": "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig"
}
EOF
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "title", "test updated"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_index_set", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
