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
		fmt.Printf("string one: %v, string two: %v\n", expectedValue.String(), actualValue.String())
		if expectedValue.String() != actualValue.String() {
			*diffs = append(*diffs, models.FieldDiff{
				Path:     path,
				Expected: expectedValue.String(),
				Actual:   actualValue.String(),
			})
		}

	case reflect.Slice, reflect.Array:
		//verify if one of the values is nil
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

		for i := range min(expectedValue.Len(), actualValue.Len()) {
			fmt.Printf("slice one: %v, slice two: %v\n", expectedValue.Index(i), actualValue.Index(i))
			if !reflect.DeepEqual(expectedValue.Index(i).Interface(), actualValue.Index(i).Interface()) {
				*diffs = append(*diffs, models.FieldDiff{
					Path:     path,
					Expected: expectedValue.Index(i).Interface(),
					Actual:   actualValue.Index(i).Interface(),
				})
			}
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

	case reflect.Float32, reflect.Float64:
		if expectedValue.Float() != actualValue.Float() {
			*diffs = append(*diffs, models.FieldDiff{
				Path:     path,
				Expected: expectedValue.Float(),
				Actual:   actualValue.Float(),
			})
		}
	}
}
func buildPath(parent, field string) string {
	if parent == "" {
		return field
	}
	return parent + "." + field
}
