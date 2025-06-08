package main

import (
	"testing"
)

func TestFormatComparisonValue_String(t *testing.T) {
	input := "hello world"
	expected := `"hello world"`
	result := formatComparisonValue(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestFormatComparisonValue_Int(t *testing.T) {
	input := 42
	expected := "42"
	result := formatComparisonValue(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestFormatComparisonValue_Uint(t *testing.T) {
	input := uint(42)
	expected := "42"
	result := formatComparisonValue(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestFormatComparisonValue_Float(t *testing.T) {
	input := 3.14159
	expected := "3.14159"
	result := formatComparisonValue(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestFormatComparisonValue_Bool(t *testing.T) {
	tests := []struct {
		input    bool
		expected string
	}{
		{true, "true"},
		{false, "false"},
	}

	for _, test := range tests {
		result := formatComparisonValue(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}

func TestFormatComparisonValue_Slice(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"empty slice", []int{}, "[]"},
		{"int slice", []int{1, 2, 3}, "[1, 2, 3]"},
		{"string slice", []string{"a", "b"}, `["a", "b"]`},
		{"nil slice", []int(nil), "nil"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := formatComparisonValue(test.input)
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}

func TestFormatComparisonValue_Map(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"string map", map[string]int{"a": 1, "b": 2}, `map["a": 1, "b": 2]`},
		{"nil map", map[string]int(nil), "nil"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := formatComparisonValue(test.input)
			// Note: map iteration order is not guaranteed, so we might need to adjust this test
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}

func TestFormatComparisonValue_Pointer(t *testing.T) {
	value := 42
	ptr := &value
	var nilPtr *int

	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"valid pointer", ptr, "42"},
		{"nil pointer", nilPtr, "nil"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := formatComparisonValue(test.input)
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}

func TestFormatComparisonValue_Struct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	type PrivateFields struct {
		Name     string
		age      int // private field
		Internal string
	}

	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			"simple struct",
			Person{Name: "John", Age: 30},
			`{Name: "John", Age: 30}`,
		},
		{
			"struct with private fields",
			PrivateFields{Name: "John", age: 30, Internal: "test"},
			`{Name: "John", Internal: "test"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := formatComparisonValue(test.input)
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}

func TestFormatComparisonValue_ComplexTypes(t *testing.T) {
	type Address struct {
		Street string
		City   string
	}

	type Person struct {
		Name    string
		Age     int
		Address *Address
		Hobbies []string
	}

	person := Person{
		Name: "John",
		Age:  30,
		Address: &Address{
			Street: "123 Main St",
			City:   "New York",
		},
		Hobbies: []string{"reading", "gaming"},
	}

	result := formatComparisonValue(person)
	expected := `{Name: "John", Age: 30, Address: {Street: "123 Main St", City: "New York"}, Hobbies: ["reading", "gaming"]}`

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
