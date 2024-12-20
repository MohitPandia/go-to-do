package utils

import "reflect"

func StructHasField(s interface{}, fieldName string) bool {
	v := reflect.ValueOf(s)
	// Ensure we have the right kind of input
	if v.Kind() != reflect.Struct {
		return false
	}

	field := v.FieldByName(fieldName)
	return field.IsValid()
}
