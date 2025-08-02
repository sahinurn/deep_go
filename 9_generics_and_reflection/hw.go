package main

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

func Serialize[T any](val T) string {
	t := reflect.TypeOf(val)
	v := reflect.ValueOf(val)

	b := strings.Builder{}

	for i := range t.NumField() {
		fieldType := t.Field(i)
		tag, ok := fieldType.Tag.Lookup("properties")
		if !ok {
			continue
		}

		fieldValue := v.Field(i)

		tagArgs := strings.Split(tag, ",")
		if len(tagArgs) == 0 {
			continue
		}

		omitempty := slices.Contains(tagArgs, "omitempty")
		if omitempty && fieldValue.IsZero() {
			continue
		}

		fieldName := tagArgs[0]

		val, ok := t.FieldByName(t.Field(i).Name)
		if !ok {
			continue
		}

		reflect.ValueOf(val)

		_, _ = b.WriteString(fieldName + "=" + fmt.Sprintf("%v", fieldValue.Interface()))
		if i != t.NumField()-1 {
			_, _ = b.WriteString("\n")
		}
	}

	return b.String()
}
