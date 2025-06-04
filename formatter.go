package main

import (
	"fmt"
	"reflect"
	"strings"
)

func formatClean(obj interface{}) string {
	return formatValue(reflect.ValueOf(obj))
}

func formatValue(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Struct:
		parts := []string{}
		typeOfT := v.Type()

		for i := range v.NumField() {
			field := typeOfT.Field(i)
			fieldValue := v.Field(i)

			fmt.Println("field", field.Name, "value", formatValue(fieldValue))

			parts = append(parts, fmt.Sprintf("%s: %v", field.Name, formatValue(fieldValue)))
		}

		return fmt.Sprintf("{%s}", strings.Join(parts, ", "))

	case reflect.String:
		return fmt.Sprintf("\"%s\"", v.String())

	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int())

	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", v.Float())

	case reflect.Ptr:
		if v.IsNil() {
			return "nil"
		}
		return formatValue(v.Elem())

	case reflect.Slice, reflect.Array:
		if v.IsNil() {
			return "nil"
		}

		if v.Len() == 0 {
			return "[]"
		}

		elements := []string{}

		for i := range v.Len() {
			elements = append(elements, formatValue(v.Index(i)))
		}

		return fmt.Sprintf("[%s]", strings.Join(elements, ", "))

	case reflect.Map:
		if v.IsNil() {
			return "nil"
		}
		var pairs []string
		for _, key := range v.MapKeys() {
			value := v.MapIndex(key)
			pairs = append(pairs, fmt.Sprintf("%s: %s", formatValue(key), formatValue(value)))
		}
		return fmt.Sprintf("map[%s]", strings.Join(pairs, ", "))
	}

	return ""
}
