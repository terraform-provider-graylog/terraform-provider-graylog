package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/suzuki-shunsuke/go-dataeq/dataeq"
)

func HandleGetResourceError(
	d *schema.ResourceData, resp *http.Response, err error, codes ...int,
) error {
	if resp == nil {
		return err
	}
	if resp.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	for _, code := range codes {
		if resp.StatusCode == code {
			d.SetId("")
			return nil
		}
	}
	return err
}

func SchemaDiffSuppressJSONString(k, oldV, newV string, d *schema.ResourceData) bool {
	b, err := dataeq.JSON.Equal([]byte(oldV), []byte(newV))
	if err != nil {
		return false
	}
	return b
}

func GenStateFunc(keys ...string) schema.StateFunc {
	return func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
		a := strings.Split(d.Id(), "/")
		size := len(keys)
		if len(a) != size {
			return nil, errors.New("format of import argument should be " + strings.Join(keys, "/"))
		}
		for i, k := range keys {
			if err := d.Set(k, a[i]); err != nil {
				return nil, err
			}
		}
		return []*schema.ResourceData{d}, nil
	}
}

func RenameKey(data map[string]interface{}, oldKey, newKey string) (interface{}, bool) {
	v, ok := data[oldKey]
	if !ok {
		return nil, false
	}
	delete(data, oldKey)
	data[newKey] = v
	return v, true
}

var ValidateIsJSON = WrapValidateFunc(func(value interface{}, key string) error {
	var a interface{}
	if err := json.Unmarshal([]byte(value.(string)), &a); err != nil {
		return fmt.Errorf("the value of the field '%s' must be JSON string: %w", key, err)
	}
	return nil
})

var ValidateIsMapJSON = WrapValidateFunc(func(value interface{}, key string) error {
	var a interface{}
	if err := json.Unmarshal([]byte(value.(string)), &a); err != nil {
		return fmt.Errorf("the value of the field '%s' must be JSON string: %w", key, err)
	}
	if _, ok := a.(map[string]interface{}); !ok {
		return errors.New("the value of the field '" + key + "' must be JSON string of map")
	}
	return nil
})

func WrapValidateFunc(f func(v interface{}, k string) error) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (s []string, es []error) {
		if err := f(v, k); err != nil {
			es = append(es, err)
		}
		return
	}
}

func SetDefaultValue(data map[string]interface{}, key string, value interface{}) {
	if _, ok := data[key]; !ok {
		data[key] = value
	}
}
