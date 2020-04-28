package graylog

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/dashboard"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/dashboard/position"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/dashboard/widget"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/event/definition"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/event/notification"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/role"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/stream"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/stream/alarmcallback"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/stream/alert/condition"
	streamOutput "github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/stream/output"
	streamRule "github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/stream/rule"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/system/grok"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/system/indices/indexset"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/system/input"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/system/input/extractor"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/system/input/staticfield"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/system/ldap/setting"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/system/output"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/system/pipeline/connection"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/system/pipeline/pipeline"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/system/pipeline/rule"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/user"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/resource/view"
)

var resourceMap = map[string]*schema.Resource{
	"graylog_alarm_callback":             alarmcallback.Resource(),
	"graylog_alert_condition":            condition.Resource(),
	"graylog_dashboard":                  dashboard.Resource(),
	"graylog_dashboard_widget":           widget.Resource(),
	"graylog_dashboard_widget_positions": position.Resource(),
	"graylog_event_definition":           definition.Resource(),
	"graylog_event_notification":         notification.Resource(),
	"graylog_extractor":                  extractor.Resource(),
	"graylog_grok_pattern":               grok.Resource(),
	"graylog_index_set":                  indexset.Resource(),
	"graylog_input":                      input.Resource(),
	"graylog_input_static_fields":        staticfield.Resource(),
	"graylog_ldap_setting":               setting.Resource(),
	"graylog_output":                     output.Resource(),
	"graylog_pipeline":                   pipeline.Resource(),
	"graylog_pipeline_connection":        connection.Resource(),
	"graylog_pipeline_rule":              rule.Resource(),
	"graylog_role":                       role.Resource(),
	"graylog_stream":                     stream.Resource(),
	"graylog_stream_output":              streamOutput.Resource(),
	"graylog_stream_rule":                streamRule.Resource(),
	"graylog_user":                       user.Resource(),
	"graylog_view":                       view.Resource(),
}
