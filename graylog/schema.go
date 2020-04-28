package graylog

// import "github.com/hashicorp/terraform/helper/schema"
//
// type Schema struct {
// 	Schema *schema.Schema
// }
//
// func Get(sc *schema.Schema, data interface{}) (interface{}, error) {
// 	// json -> schema
// 	switch sc.Type {
// 	case schema.TypeList:
// 		a, ok := data.([]interface{})
// 		if ok {
// 			list := make([]interface{}, len(a))
// 			for i, b := range a {
// 				for k, t := range sc.Elem.Schema {
//
// 				}
//
// 			}
// 		}
//
// 		switch t := sc.Elem.(type) {
// 		case *schema.Schema:
// 			a, ok := data.([]interface{})
// 			if ok {
// 				return a, nil
// 			}
// 			// TODO convert map to list
// 		case map[string]*schema.Schema:
// 		}
//
// 		if sc.MaxItems == 1 && sc.MinItems == 1 {
//
// 		}
// 	case schema.TypeMap:
//
// 	}
// 	return data, nil
// }
