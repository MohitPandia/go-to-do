package utils

import (
	"strings"
)

func ExtractJSONValueWithPrefix(data interface{}, prefix string) []string {
	var values []string
	switch v := data.(type) {
	case map[string]interface{}:
		// Iterate over the map
		for _, value := range v {
			// Check if the value is a string that starts with prefix
			if strValue, ok := value.(string); ok {
				if strings.HasPrefix(strValue, prefix) {
					values = append(values, strValue)
				}
			} else {
				// Recurse if the value is another map or array
				subValues := ExtractJSONValueWithPrefix(value, prefix)
				values = append(values, subValues...)
			}
		}
	case []interface{}:
		// Iterate over the array
		for _, item := range v {
			subValues := ExtractJSONValueWithPrefix(item, prefix)
			values = append(values, subValues...)
		}
	}
	return values
}
