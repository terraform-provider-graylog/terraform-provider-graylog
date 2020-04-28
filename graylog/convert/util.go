package convert

func ConvertListToMap(data []interface{}, key string) map[string]interface{} {
	m := make(map[string]interface{}, len(data))
	for _, d := range data {
		elem := d.(map[string]interface{})
		a := elem[key].(string)
		delete(elem, key)
		m[a] = elem
	}
	return m
}

func ConvertMapToList(data map[string]interface{}, key string) []interface{} {
	list := make([]interface{}, len(data))
	i := 0
	for k, d := range data {
		elem := d.(map[string]interface{})
		elem[key] = k
		list[i] = elem
		i++
	}
	return list
}

func ConvertInterfaceListToStringList(data []interface{}) []string {
	list := make([]string, len(data))
	for i, a := range data {
		list[i] = a.(string)
	}
	return list
}
