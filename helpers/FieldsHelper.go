package helpers

import (
	"reflect"
)

func GetFields(structName interface{}) map[string]interface{} {
	typeStruct := reflect.TypeOf(structName)
	fields := make(map[string]interface{})

	var fieldName, fieldType string

	for i := 0; i < typeStruct.NumField(); i++ {
		fieldName = typeStruct.Field(i).Tag.Get("json")
		fieldType = typeStruct.Field(i).Type.Name()
		fields[fieldName] = fieldType
	}
	return fields
}
