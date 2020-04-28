# data "graylog_dashboard" "test" {
#   title = "test"
# }

resource "graylog_dashboard" "test" {
  title       = "test"
  description = "test"
}

resource "graylog_dashboard_widget" "test" {
  description = "Stream search result count change"
  # dashboard_id = data.graylog_dashboard.test.id
  dashboard_id = graylog_dashboard.test.id
  type         = "STREAM_SEARCH_RESULT_COUNT"
  cache_time   = 10
  config       = <<EOF
{
  "timerange": {
    "type": "relative",
    "range": 400
  },
  "lower_is_better": true,
  "stream_id": "${graylog_stream.test.id}",
  "trend": true,
  "query": ""
}
EOF
}

resource "graylog_dashboard_widget" "test2" {
  description  = "Quick values"
  dashboard_id = graylog_dashboard.test.id
  type         = "QUICKVALUES"
  cache_time   = 10

  config = <<EOF
{
  "timerange": {
    "type": "relative",
    "range": 300
  },
  "field": "status",
  "stream_id": "${graylog_stream.test.id}",
  "query": "",
  "show_data_table": true,
  "limit": 5,
  "sort_order": "desc",
  "show_pie_chart": true,
  "data_table_limit": 60
}
EOF
}

resource "graylog_dashboard_widget_positions" "test" {
  dashboard_id = graylog_dashboard_widget.test.dashboard_id

  positions = <<EOF
{
  "${graylog_dashboard_widget.test.widget_id}": {
    "width": 1,
    "col": 0,
    "row": 0,
    "height": 1
  },
  "${graylog_dashboard_widget.test2.widget_id}": {
    "width": 2,
    "col": 1,
    "row": 0,
    "height": 2
  }
}
EOF
}

resource "graylog_dashboard_widget" "stacked_chart" {
  description  = "stacked chart"
  dashboard_id = graylog_dashboard.test.id
  type         = "STACKED_CHART"
  cache_time   = 10
  config       = <<EOF
{
  "interval": "hour",
  "timerange": {
    "type": "relative",
    "range": 86400
  },
  "renderer": "bar",
  "interpolation": "linear",
  "series": [
    {
      "query": "",
      "field": "AccessMask",
      "statistical_function": "count"
    }
  ]
}
EOF
}
