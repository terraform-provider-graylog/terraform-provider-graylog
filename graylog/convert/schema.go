package convert

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetSchema(data interface{}, sc *schema.Schema) (interface{}, error) {
	// ResourceData => Graylog API (Post/Update)
	switch sc.Type { //nolint:exhaustive
	case schema.TypeList:
		if sc.MinItems == 1 && sc.MaxItems == 1 {
			return data.([]interface{})[0], nil
		}
	case schema.TypeSet:
		list := data.(*schema.Set).List()
		if rsc, ok := sc.Elem.(*schema.Resource); ok {
			arr := make([]interface{}, len(list))
			for i, a := range list {
				elem, err := GetResource(a.(map[string]interface{}), rsc)
				if err != nil {
					return nil, err
				}
				arr[i] = elem
			}
			return arr, nil
		}
		return list, nil
	case schema.TypeMap:
	}
	return data, nil
}

func SetSchema(data interface{}, sc *schema.Schema) (interface{}, error) {
	// Graylog API (Get) => ResourceData
	switch sc.Type { //nolint:exhaustive
	case schema.TypeList:
		return SetTypeList(data, sc)
	case schema.TypeSet:
		return SetTypeList(data, sc)
	case schema.TypeMap:
	}
	return data, nil
}

func SetTypeList(data interface{}, sc *schema.Schema) (interface{}, error) {
	switch t := sc.Elem.(type) {
	case *schema.Resource:
		switch v := data.(type) { //nolint:gocritic
		case []interface{}:
			ret := make([]interface{}, len(v))
			for i, a := range v {
				b, err := SetResource(a.(map[string]interface{}), t)
				if err != nil {
					return nil, err
				}
				ret[i] = b
			}
			return ret, nil
		}
	case *schema.Schema:
	}
	if sc.MinItems == 1 && sc.MaxItems == 1 {
		return []interface{}{data}, nil
	}
	return data, nil
}
