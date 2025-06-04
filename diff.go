package main

import (
	"fmt"
	"reflect"

	"github.com/seu-usuario/meu-projeto/models"
)

func FindDifferences(expected, actual interface{}) []models.FieldDiff {
	var diffs []models.FieldDiff
	compare(expected, actual, "", &diffs)
	return diffs
}

func compare(expected, actual interface{}, path string, diffs *[]models.FieldDiff) {
	expectedValue := reflect.ValueOf(expected)
	actualValue := reflect.ValueOf(actual)

	if expectedValue.Kind() != actualValue.Kind() {
		*diffs = append(*diffs, models.FieldDiff{
			Path:     path,
			Expected: expectedValue.Kind(),
			Actual:   actualValue.Kind(),
		})
		return
	}

	switch expectedValue.Kind() {
	case reflect.Struct:
		typeOfT := expectedValue.Type()
		for i := range expectedValue.NumField() {
			field := typeOfT.Field(i)
			newPath := buildPath(path, field.Name)

			expectedField := expectedValue.Field(i).Interface()
			actualField := actualValue.Field(i).Interface()

			compare(expectedField, actualField, newPath, diffs)
		}

	case reflect.String:
		if expectedValue.String() != actualValue.String() {
			*diffs = append(*diffs, models.FieldDiff{
				Path:     path,
				Expected: expectedValue.String(),
				Actual:   actualValue.String(),
			})
		}

	case reflect.Bool:
		if expectedValue.Bool() != actualValue.Bool() {
			*diffs = append(*diffs, models.FieldDiff{
				Path:     path,
				Expected: expectedValue.Bool(),
				Actual:   actualValue.Bool(),
			})
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if expectedValue.Int() != actualValue.Int() {
			*diffs = append(*diffs, models.FieldDiff{
				Path:     path,
				Expected: expectedValue.Interface(),
				Actual:   actualValue.Interface(),
			})
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if expectedValue.Uint() != actualValue.Uint() {
			*diffs = append(*diffs, models.FieldDiff{
				Path:     path,
				Expected: expectedValue.Interface(),
				Actual:   actualValue.Interface(),
			})
		}

	case reflect.Float32, reflect.Float64:
		if expectedValue.Float() != actualValue.Float() {
			*diffs = append(*diffs, models.FieldDiff{
				Path:     path,
				Expected: expectedValue.Interface(),
				Actual:   actualValue.Interface(),
			})
		}

	case reflect.Ptr:
		if expectedValue.IsNil() != actualValue.IsNil() {
			*diffs = append(*diffs, models.FieldDiff{
				Path:     path,
				Expected: expectedValue.Interface(),
				Actual:   actualValue.Interface(),
			})
			return
		}
		if !expectedValue.IsNil() {
			compare(expectedValue.Elem().Interface(), actualValue.Elem().Interface(), path, diffs)
		}

	case reflect.Slice, reflect.Array:
		if expectedValue.IsNil() != actualValue.IsNil() {
			*diffs = append(*diffs, models.FieldDiff{
				Path:     path,
				Expected: expectedValue.Interface(),
				Actual:   actualValue.Interface(),
			})
			return
		}

		if expectedValue.Len() == 0 && actualValue.Len() == 0 {
			return
		}

		if expectedValue.Len() != actualValue.Len() {
			*diffs = append(*diffs, models.FieldDiff{
				Path:     path,
				Expected: expectedValue.Interface(),
				Actual:   actualValue.Interface(),
			})
			return
		}

		//compare elements one by one
		for i := range expectedValue.Len() {
			if !reflect.DeepEqual(expectedValue.Index(i).Interface(), actualValue.Index(i).Interface()) {
				elementPath := buildPath(path, fmt.Sprintf("[%d]", i))
				compare(expectedValue.Index(i).Interface(), actualValue.Index(i).Interface(), elementPath, diffs)
			}
		}

	case reflect.Map:
		if expectedValue.IsNil() != actualValue.IsNil() {
			*diffs = append(*diffs, models.FieldDiff{
				Path:     path,
				Expected: expectedValue.Interface(),
				Actual:   actualValue.Interface(),
			})
			return
		}

		if expectedValue.IsNil() {
			return
		}

		if expectedValue.Len() != actualValue.Len() {
			*diffs = append(*diffs, models.FieldDiff{
				Path:     path,
				Expected: expectedValue.Interface(),
				Actual:   actualValue.Interface(),
			})
			return
		}

		for _, key := range expectedValue.MapKeys() {
			actualVal := actualValue.MapIndex(key)
			if !actualVal.IsValid() {
				*diffs = append(*diffs, models.FieldDiff{
					Path:     buildPath(path, fmt.Sprintf("[%v]", key.Interface())),
					Expected: expectedValue.MapIndex(key).Interface(),
					Actual:   nil,
				})
				continue
			}

			keyPath := buildPath(path, fmt.Sprintf("[%v]", key.Interface()))
			compare(expectedValue.MapIndex(key).Interface(), actualVal.Interface(), keyPath, diffs)
		}
	}
}

func buildPath(parent, field string) string {
	if parent == "" {
		return field
	}
	return parent + "." + field
}
