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
  config = jsonencode({
    timerange = {
      type  = "relative"
      range = 400
    }
    lower_is_better = true
    trend           = true
    stream_id       = graylog_stream.test.id
    query           = ""
  })
}

resource "graylog_dashboard_widget" "test2" {
  description  = "Quick values"
  dashboard_id = graylog_dashboard.test.id
  type         = "QUICKVALUES"
  cache_time   = 10

  config = jsonencode({
    timerange = {
      type  = "relative"
      range = 300
    }
    stream_id        = graylog_stream.test.id
    query            = ""
    field            = "status"
    show_data_table  = true
    show_pie_chart   = true
    limit            = 5
    sort_order       = "desc"
    data_table_limit = 60
  })
}

resource "graylog_dashboard_widget_positions" "test" {
  dashboard_id = graylog_dashboard_widget.test.dashboard_id

  positions = jsonencode({
    (graylog_dashboard_widget.test.widget_id) = {
      row    = 0
      col    = 0
      height = 1
      width  = 1
    }
    (graylog_dashboard_widget.test2.widget_id) = {
      row    = 0
      col    = 1
      height = 2
      width  = 2
    }
  })
}

resource "graylog_dashboard_widget" "stacked_chart" {
  description  = "stacked chart"
  dashboard_id = graylog_dashboard.test.id
  type         = "STACKED_CHART"
  cache_time   = 10

  config = jsonencode({
    interval = "hour"
    timerange = {
      type  = "relative"
      range = 86400
    },
    renderer      = "bar"
    interpolation = "linear"
    series = [
      {
        query                = ""
        field                = "AccessMask"
        statistical_function = "count"
      }
    ]
  })
}
