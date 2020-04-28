package convert

import "github.com/hashicorp/terraform/helper/schema"

func GetFromResourceData(
	d *schema.ResourceData, rsc *schema.Resource,
) (map[string]interface{}, error) {
	ret := make(map[string]interface{}, len(rsc.Schema))
	for k, sc := range rsc.Schema {
		a, err := GetSchema(d.Get(k), sc)
		if err != nil {
			return nil, err
		}
		ret[k] = a
	}
	return ret, nil
}

func SetResourceData(
	d *schema.ResourceData, rsc *schema.Resource, data map[string]interface{},
) error {
	for k, sc := range rsc.Schema {
		v, ok := data[k]
		if !ok {
			continue
		}
		val, err := SetSchema(v, sc)
		if err != nil {
			return err
		}
		if err := d.Set(k, val); err != nil {
			return err
		}
	}
	return nil
}
