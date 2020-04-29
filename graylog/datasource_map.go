package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/datasource/dashboard"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/datasource/stream"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/datasource/system/indices/indexset"
)

var dataSourcesMap = map[string]*schema.Resource{
	"graylog_dashboard": dashboard.DataSource(),
	"graylog_index_set": indexset.DataSource(),
	"graylog_stream":    stream.DataSource(),
}
