package main

import (
	"fmt"
	"reflect"
	"strings"
)

func formatTestOutput(obj interface{}) string {
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

// formatComparisonValue formats objects with improved handling of types and exported fields only
func formatComparisonValue(obj interface{}) string {
	return formatValueComparison(reflect.ValueOf(obj))
}

// formatValueComparison handles the formatting logic for different reflect.Value types
func formatValueComparison(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Struct:
		var parts []string
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			if !field.IsExported() {
				continue
			}
			fieldValue := v.Field(i)
			parts = append(parts, fmt.Sprintf("%s: %s", field.Name, formatValueComparison(fieldValue)))
		}
		return fmt.Sprintf("{%s}", strings.Join(parts, ", "))

	case reflect.String:
		return fmt.Sprintf(`"%s"`, v.String())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%g", v.Float())

	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())

	case reflect.Ptr:
		if v.IsNil() {
			return "nil"
		}
		return formatValueComparison(v.Elem())

	case reflect.Slice, reflect.Array:
		if v.IsNil() {
			return "nil"
		}
		var elements []string
		for i := 0; i < v.Len(); i++ {
			elements = append(elements, formatValueComparison(v.Index(i)))
		}
		return fmt.Sprintf("[%s]", strings.Join(elements, ", "))

	case reflect.Map:
		if v.IsNil() {
			return "nil"
		}
		var pairs []string
		for _, key := range v.MapKeys() {
			value := v.MapIndex(key)
			pairs = append(pairs, fmt.Sprintf("%s: %s", formatValueComparison(key), formatValueComparison(value)))
		}
		return fmt.Sprintf("map[%s]", strings.Join(pairs, ", "))

	default:
		return fmt.Sprintf("%v", v.Interface())
	}
}
