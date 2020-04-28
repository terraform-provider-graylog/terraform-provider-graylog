package util

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/zclconf/go-cty/cty"
)

func UpgraderType() cty.Type {
	rsc := schema.Resource{}
	return rsc.CoreConfigSchema().ImpliedType()
}
