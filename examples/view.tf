# TODO View doesn't support yet
# resource "graylog_view" "test" {
#   title     = "test"
#   type      = "DASHBOARD"
#   search_id = ""
#   state {
#     id = ""
#     widgets {
#       widget_id = ""
#       type      = "messages"
#       config    = <<EOF
# {
#   "fields": [
#     "timestamp",
#     "source"
#   ],
#   "show_message_row": true,
#   "decorators": [],
#   "sort": [
#     {
#       "type": "pivot",
#       "field": "timestamp",
#       "direction": "Descending"
#     }
#   ]
# }
# EOF
#       timerange = <<EOF
# {
#   "type": "relative",
# 	"range": 0
# }
# EOF
#       query {
#         type         = "elasticsearch"
#         query_string = ""
#       }
#     }
# 
#     widget_mapping = <<EOF
# {
#   "ed5ec0d9-4dec-4f4a-8418-bbc1a3ef1c5c": [
#     "7741cf47-f074-4c37-be92-757db924bb76"
#   ],
#   "77092ac1-5d54-492f-b7dc-13b9d0933a6f": [
#     "6a45b115-92f3-4b97-afb6-a529bca322ad"
#   ]
# }
# EOF
# 
#     positions = <<EOF
# {
#   "ed5ec0d9-4dec-4f4a-8418-bbc1a3ef1c5c": {
#     "col": 1,
#     "row": 1,
#     "height": 2,
#     "width": "Infinity"
#   },
#   "77092ac1-5d54-492f-b7dc-13b9d0933a6f": {
#     "col": 1,
#     "row": 3,
#     "height": 6,
#     "width": "Infinity"
#   }
# }
# EOF
# 
#     titles {
#       widget = {
#         "ed5ec0d9-4dec-4f4a-8418-bbc1a3ef1c5c" : "Message Count",
#         "77092ac1-5d54-492f-b7dc-13b9d0933a6f" : "All Messages"
#       }
#     }
# 
#   }
# 
#   summary     = ""
#   description = ""
#   owner       = ""
# }
