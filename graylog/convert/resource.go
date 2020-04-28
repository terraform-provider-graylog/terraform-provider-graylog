package convert

import "github.com/hashicorp/terraform/helper/schema"

func GetResource(
	data map[string]interface{}, rsc *schema.Resource,
) (map[string]interface{}, error) {
	ret := make(map[string]interface{}, len(data))
	for k, sc := range rsc.Schema {
		v, ok := data[k]
		if !ok {
			continue
		}
		a, err := GetSchema(v, sc)
		if err != nil {
			return nil, err
		}
		ret[k] = a
	}
	return ret, nil
}

func SetResource(data map[string]interface{}, rsc *schema.Resource) (interface{}, error) {
	ret := make(map[string]interface{}, len(data))
	for k, sc := range rsc.Schema {
		v, ok := data[k]
		if !ok {
			continue
		}
		a, err := SetSchema(v, sc)
		if err != nil {
			return nil, err
		}
		ret[k] = a
	}
	return ret, nil
}
