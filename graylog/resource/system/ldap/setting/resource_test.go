package setting

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/suzuki-shunsuke/flute/v2/flute"
	"github.com/terraform-provider-graylog/terraform-provider-graylog/graylog/testutil"
)

func TestAccLDAPSetting(t *testing.T) {
	if err := testutil.SetEnv(); err != nil {
		t.Fatal(err)
	}

	settingBody := ""

	resourceURLPath := "/api/system/ldap/settings"
	resourceName := "graylog_ldap_setting.system"

	getRoute := flute.Route{
		Name: "get a setting",
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
					Body:       ioutil.NopCloser(strings.NewReader(settingBody)),
				}, nil
			},
		},
	}

	postRoute := flute.Route{
		Name: "create a setting",
		Matcher: flute.Matcher{
			Method: "PUT",
		},
		Tester: flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "enabled": false,
  "system_username": "",
  "system_password": "",
  "ldap_uri": "ldap://localhost:389",
  "use_start_tls": false,
  "trust_all_certificates": false,
  "active_directory": false,
  "search_base": "",
  "search_pattern": "",
  "display_name_attribute": "",
  "default_group": "Reader",
  "group_mapping": {},
  "group_search_base": "",
  "group_id_attribute": "",
  "additional_default_groups": [],
  "group_search_pattern": ""
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				settingBody = `{
  "enabled": false,
  "system_username": "",
  "ldap_uri": "ldap://localhost:389",
  "use_start_tls": false,
  "trust_all_certificates": false,
  "active_directory": false,
  "search_base": "",
  "search_pattern": "",
  "display_name_attribute": "",
  "default_group": "Reader",
  "group_mapping": {},
  "group_search_base": "",
  "group_id_attribute": "",
  "additional_default_groups": [],
  "group_search_pattern": "",
  "system_password_set": false
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 204,
			},
		},
	}

	deleteRoute := flute.Route{
		Name: "delete a setting",
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
resource "graylog_ldap_setting" "system" {
  enabled                   = false
  system_username           = ""
  ldap_uri                  = "ldap://localhost:389"
  use_start_tls             = false
  trust_all_certificates    = false
  active_directory          = false
  search_base               = ""
  search_pattern            = ""
  display_name_attribute    = ""
  default_group             = "Reader"
  group_mapping             = null
  group_search_base         = null
  group_id_attribute        = null
  additional_default_groups = null
  group_search_pattern      = null
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "ldap_uri", "ldap://localhost:389"),
		),
	}

	updateRoute := flute.Route{
		Name: "update a setting",
		Matcher: flute.Matcher{
			Method: "PUT",
		},
		Tester: flute.Tester{
			Path:         resourceURLPath,
			PartOfHeader: testutil.Header(),
			BodyJSONString: `{
  "enabled": false,
  "system_username": "",
  "system_password": "",
  "ldap_uri": "ldap://localhost:489",
  "use_start_tls": false,
  "trust_all_certificates": false,
  "active_directory": false,
  "search_base": "",
  "search_pattern": "",
  "display_name_attribute": "",
  "default_group": "Reader",
  "group_mapping": {},
  "group_search_base": "",
  "group_id_attribute": "",
  "additional_default_groups": [],
  "group_search_pattern": ""
}`,
			Test: func(t *testing.T, req *http.Request, svc flute.Service, route flute.Route) {
				settingBody = `{
  "enabled": false,
  "system_username": "",
  "ldap_uri": "ldap://localhost:489",
  "use_start_tls": false,
  "trust_all_certificates": false,
  "active_directory": false,
  "search_base": "",
  "search_pattern": "",
  "display_name_attribute": "",
  "default_group": "Reader",
  "group_mapping": {},
  "group_search_base": "",
  "group_id_attribute": "",
  "additional_default_groups": [],
  "group_search_pattern": "",
  "system_password_set": false
}`
			},
		},
		Response: flute.Response{
			Base: http.Response{
				StatusCode: 204,
			},
		},
	}

	updateStep := resource.TestStep{
		ResourceName: resourceName,
		PreConfig: func() {
			testutil.SetHTTPClient(t, getRoute, updateRoute, deleteRoute)
		},
		Config: `
resource "graylog_ldap_setting" "system" {
  enabled                   = false
  system_username           = ""
  ldap_uri                  = "ldap://localhost:489"
  use_start_tls             = false
  trust_all_certificates    = false
  active_directory          = false
  search_base               = ""
  search_pattern            = ""
  display_name_attribute    = ""
  default_group             = "Reader"
  group_mapping             = null
  group_search_base         = null
  group_id_attribute        = null
  additional_default_groups = null
  group_search_pattern      = null
}
`,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "ldap_uri", "ldap://localhost:489"),
		),
	}

	resource.Test(t, resource.TestCase{
		Providers: testutil.SingleResourceProviders("graylog_ldap_setting", Resource()),
		Steps: []resource.TestStep{
			createStep,
			updateStep,
		},
	})
}
