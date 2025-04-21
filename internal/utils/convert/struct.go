package convert

import (
	"reflect"
	"strings"
)

// StructToMap converts a struct to a map, using json tags as keys
// Only non-nil pointer fields will be included in the map
func StructToMap(input interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(input)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		// Only process pointer fields that are not nil
		if field.Kind() == reflect.Ptr && !field.IsNil() {
			// Get the json tag name
			jsonTag := t.Field(i).Tag.Get("json")
			if jsonTag != "" {
				result[jsonTag] = field.Elem().Interface()
			}
		}
	}

	return result
}

// Helper function to check if string contains substring (case insensitive)
func ContainsIgnoreCase(s, substr string) bool {
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)
	return strings.Contains(s, substr)
}
