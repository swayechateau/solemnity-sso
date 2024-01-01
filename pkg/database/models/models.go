package models

import (
	"fmt"
	"reflect"
)

func ProcessStruct(s interface{}) {
	val := reflect.ValueOf(s).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		dbTag := typ.Field(i).Tag.Get("db")

		if dbTag != "-" {
			// Process fields that are not excluded from DB
			fmt.Printf("DB field: %s, Value: %v\n", dbTag, field.Interface())
		}
	}
}
