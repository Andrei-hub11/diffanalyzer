package main

import (
	"testing"

	"github.com/seu-usuario/meu-projeto/models"
	"github.com/stretchr/testify/assert"
)

func TestFindDifferences_IdenticalStructs_ShouldReturnNoDifferences(t *testing.T) {
	// Arrange
	person1 := models.Person{
		ID:     1,
		Name:   "Alice",
		Emails: []string{"alice@company.com", "alice@personal.com"},
		Profile: models.Profile{
			Bio:  "Engineer",
			Tags: []string{"go", "backend"},
			Address: models.Address{
				City:    "São Paulo",
				Country: "Brasil",
			},
		},
	}

	person2 := models.Person{
		ID:     1,
		Name:   "Alice",
		Emails: []string{"alice@company.com", "alice@personal.com"},
		Profile: models.Profile{
			Bio:  "Engineer",
			Tags: []string{"go", "backend"},
			Address: models.Address{
				City:    "São Paulo",
				Country: "Brasil",
			},
		},
	}

	// Act
	diffs := FindDifferences(person1, person2)

	// Assert
	assert.Empty(t, diffs, "Identical structs should have no differences")
}

func TestFindDifferences_DifferentStringFields_ShouldReturnCorrectDifferences(t *testing.T) {
	// Arrange
	person1 := models.Person{
		ID:   1,
		Name: "Alice",
	}

	person2 := models.Person{
		ID:   1,
		Name: "Bob",
	}

	// Act
	diffs := FindDifferences(person1, person2)

	// Assert
	assert.Len(t, diffs, 1, "Should find exactly one difference")
	assert.Equal(t, "Name", diffs[0].Path)
	assert.Equal(t, "Alice", diffs[0].Expected)
	assert.Equal(t, "Bob", diffs[0].Actual)
}

func TestFindDifferences_DifferentIntegerFields_ShouldReturnCorrectDifferences(t *testing.T) {
	// Arrange
	person1 := models.Person{
		ID:   1,
		Name: "Alice",
	}

	person2 := models.Person{
		ID:   2,
		Name: "Alice",
	}

	// Act
	diffs := FindDifferences(person1, person2)

	// Assert
	assert.Len(t, diffs, 1, "Should find exactly one difference")
	assert.Equal(t, "ID", diffs[0].Path)
	assert.Equal(t, 1, diffs[0].Expected)
	assert.Equal(t, 2, diffs[0].Actual)
}

func TestFindDifferences_DifferentNestedStructFields_ShouldReturnCorrectPath(t *testing.T) {
	// Arrange
	person1 := models.Person{
		ID:   1,
		Name: "Alice",
		Profile: models.Profile{
			Bio: "Engineer",
			Address: models.Address{
				City:    "São Paulo",
				Country: "Brasil",
			},
		},
	}

	person2 := models.Person{
		ID:   1,
		Name: "Alice",
		Profile: models.Profile{
			Bio: "Developer",
			Address: models.Address{
				City:    "Rio de Janeiro",
				Country: "Brazil",
			},
		},
	}

	// Act
	diffs := FindDifferences(person1, person2)

	// Assert
	assert.Len(t, diffs, 3, "Should find three differences")

	pathsFound := make(map[string]bool)
	for _, diff := range diffs {
		pathsFound[diff.Path] = true
	}

	assert.True(t, pathsFound["Profile.Bio"], "Should find difference in Profile.Bio")
	assert.True(t, pathsFound["Profile.Address.City"], "Should find difference in Profile.Address.City")
	assert.True(t, pathsFound["Profile.Address.Country"], "Should find difference in Profile.Address.Country")
}

func TestFindDifferences_DifferentSliceElements_ShouldReturnDifferences(t *testing.T) {
	// Arrange
	person1 := models.Person{
		ID:     1,
		Name:   "Alice",
		Emails: []string{"alice@company.com", "alice@personal.com"},
		Profile: models.Profile{
			Tags: []string{"go", "backend", "api"},
		},
	}

	person2 := models.Person{
		ID:     1,
		Name:   "Alice",
		Emails: []string{"alice@company.com", "alice@gmail.com"},
		Profile: models.Profile{
			Tags: []string{"go", "frontend", "api"},
		},
	}

	// Act
	diffs := FindDifferences(person1, person2)

	// Assert
	assert.GreaterOrEqual(t, len(diffs), 2, "Should find at least two differences in slices")

	foundEmailDiff := false
	foundTagDiff := false

	for _, diff := range diffs {
		if diff.Path == "Emails.[1]" {
			foundEmailDiff = true
		}
		if diff.Path == "Profile.Tags.[1]" {
			foundTagDiff = true
		}
	}

	assert.True(t, foundEmailDiff && foundTagDiff, "Should find differences in slice elements")
}

func TestFindDifferences_EmptySlicesVsNilSlices_ShouldHandleCorrectly(t *testing.T) {
	// Arrange
	person1 := models.Person{
		ID:     1,
		Name:   "Alice",
		Emails: []string{},
	}

	person2 := models.Person{
		ID:     1,
		Name:   "Alice",
		Emails: nil,
	}

	// Act
	diffs := FindDifferences(person1, person2)

	// Assert
	assert.NotNil(t, diffs, "Should not panic with empty vs nil slices")
}

func TestFindDifferences_ComplexNestedStructure_ShouldFindAllDifferences(t *testing.T) {
	// Arrange
	expected := models.Person{
		ID:     1,
		Name:   "Alice",
		Emails: []string{"alice@company.com", "alice@personal.com"},
		Profile: models.Profile{
			Bio:  "Senior Engineer",
			Tags: []string{"go", "backend", "api", "microservices"},
			Address: models.Address{
				City:    "São Paulo",
				Country: "Brasil",
			},
		},
	}

	actual := models.Person{
		ID:     1,
		Name:   "Alice Johnson",
		Emails: []string{"alice@company.com", "alice@gmail.com"},
		Profile: models.Profile{
			Bio:  "Lead Developer",
			Tags: []string{"go", "frontend", "api", "web"},
			Address: models.Address{
				City:    "Rio de Janeiro",
				Country: "Brazil",
			},
		},
	}

	// Act
	diffs := FindDifferences(expected, actual)

	// Assert
	assert.NotEmpty(t, diffs, "Should find multiple differences")

	pathsFound := make(map[string]models.FieldDiff)
	for _, diff := range diffs {
		pathsFound[diff.Path] = diff
	}

	if nameDiff, exists := pathsFound["Name"]; exists {
		assert.Equal(t, "Alice", nameDiff.Expected)
		assert.Equal(t, "Alice Johnson", nameDiff.Actual)
	}

	if bioDiff, exists := pathsFound["Profile.Bio"]; exists {
		assert.Equal(t, "Senior Engineer", bioDiff.Expected)
		assert.Equal(t, "Lead Developer", bioDiff.Actual)
	}

	if cityDiff, exists := pathsFound["Profile.Address.City"]; exists {
		assert.Equal(t, "São Paulo", cityDiff.Expected)
		assert.Equal(t, "Rio de Janeiro", cityDiff.Actual)
	}

	if countryDiff, exists := pathsFound["Profile.Address.Country"]; exists {
		assert.Equal(t, "Brasil", countryDiff.Expected)
		assert.Equal(t, "Brazil", countryDiff.Actual)
	}
}

func TestFindDifferences_SimplePersonaStruct_ShouldWork(t *testing.T) {
	// Arrange
	pessoa1 := models.Pessoa{
		Nome:   "João",
		Idade:  30,
		Ativo:  true,
		Emails: []string{"joao@example.com"},
	}

	pessoa2 := models.Pessoa{
		Nome:   "João Silva",
		Idade:  31,
		Ativo:  false,
		Emails: []string{"joao@company.com"},
	}

	// Act
	diffs := FindDifferences(pessoa1, pessoa2)

	// Assert
	assert.NotEmpty(t, diffs, "Should find differences between different Pessoa structs")

	pathsFound := make(map[string]bool)
	for _, diff := range diffs {
		pathsFound[diff.Path] = true
	}

	assert.True(t, pathsFound["Nome"] || pathsFound["Idade"] || pathsFound["Ativo"],
		"Should find differences in at least one of the basic fields")
}

func TestFindDifferences_DifferentBoolFields_ShouldReturnCorrectDifferences(t *testing.T) {
	// Arrange
	pessoa1 := models.Pessoa{
		Nome:  "João",
		Idade: 30,
		Ativo: true,
	}

	pessoa2 := models.Pessoa{
		Nome:  "João",
		Idade: 30,
		Ativo: false,
	}

	// Act
	diffs := FindDifferences(pessoa1, pessoa2)

	// Assert
	assert.Len(t, diffs, 1, "Should find exactly one difference")
	assert.Equal(t, "Ativo", diffs[0].Path)
	assert.Equal(t, true, diffs[0].Expected)
	assert.Equal(t, false, diffs[0].Actual)
}

func TestFindDifferences_DifferentSliceLengths_ShouldReturnDifferences(t *testing.T) {
	// Arrange
	person1 := models.Person{
		ID:     1,
		Name:   "Alice",
		Emails: []string{"alice@company.com", "alice@personal.com"},
	}

	person2 := models.Person{
		ID:     1,
		Name:   "Alice",
		Emails: []string{"alice@company.com"},
	}

	// Act
	diffs := FindDifferences(person1, person2)

	// Assert
	assert.Len(t, diffs, 1, "Should find exactly one difference for different slice lengths")
	assert.Equal(t, "Emails", diffs[0].Path)
}

func TestFindDifferences_ZeroValues_ShouldDetectCorrectly(t *testing.T) {
	// Arrange
	pessoa1 := models.Pessoa{
		Nome:  "",
		Idade: 0,
		Ativo: false,
	}

	pessoa2 := models.Pessoa{
		Nome:  "João",
		Idade: 25,
		Ativo: true,
	}

	// Act
	diffs := FindDifferences(pessoa1, pessoa2)

	// Assert
	assert.Len(t, diffs, 3, "Should find differences for all non-zero vs zero values")

	pathsFound := make(map[string]bool)
	for _, diff := range diffs {
		pathsFound[diff.Path] = true
	}

	assert.True(t, pathsFound["Nome"], "Should find difference in Nome")
	assert.True(t, pathsFound["Idade"], "Should find difference in Idade")
	assert.True(t, pathsFound["Ativo"], "Should find difference in Ativo")
}

func TestFindDifferences_SliceElementPath_ShouldShowCorrectPath(t *testing.T) {
	// Arrange
	person1 := models.Person{
		ID:   1,
		Name: "Alice",
		Profile: models.Profile{
			Tags: []string{"go", "backend", "api"},
		},
	}

	person2 := models.Person{
		ID:   1,
		Name: "Alice",
		Profile: models.Profile{
			Tags: []string{"go", "frontend", "api"},
		},
	}

	// Act
	diffs := FindDifferences(person1, person2)

	// Assert
	assert.NotEmpty(t, diffs, "Should find differences in slice elements")

	foundTagDiff := false
	for _, diff := range diffs {
		if diff.Path == "Profile.Tags.[1]" {
			foundTagDiff = true
			assert.Equal(t, "backend", diff.Expected)
			assert.Equal(t, "frontend", diff.Actual)
		}
	}

	assert.True(t, foundTagDiff, "Should find difference with correct element path")
}

func TestFindDifferences_EmptyStructs_ShouldReturnNoDifferences(t *testing.T) {
	// Arrange
	addr1 := models.Address{}
	addr2 := models.Address{}

	// Act
	diffs := FindDifferences(addr1, addr2)

	// Assert
	assert.Empty(t, diffs, "Empty structs should have no differences")
}

func TestFindDifferences_NestedEmptyVsFilledStruct_ShouldDetectDifferences(t *testing.T) {
	// Arrange
	person1 := models.Person{
		ID:   1,
		Name: "Alice",
		Profile: models.Profile{
			Address: models.Address{},
		},
	}

	person2 := models.Person{
		ID:   1,
		Name: "Alice",
		Profile: models.Profile{
			Address: models.Address{
				City:    "São Paulo",
				Country: "Brasil",
			},
		},
	}

	// Act
	diffs := FindDifferences(person1, person2)

	// Assert
	assert.NotEmpty(t, diffs, "Should find differences between empty and filled nested structs")

	pathsFound := make(map[string]bool)
	for _, diff := range diffs {
		pathsFound[diff.Path] = true
	}

	assert.True(t, pathsFound["Profile.Address.City"], "Should find difference in nested City")
	assert.True(t, pathsFound["Profile.Address.Country"], "Should find difference in nested Country")
}

// ===== TESTS FOR INTEGER TYPES =====

func TestFindDifferences_AllIntegerTypes_ShouldDetectCorrectly(t *testing.T) {
	// Arrange
	data1 := models.DataTypes{
		IntValue:   42,
		Int8Value:  8,
		Int16Value: 16,
		Int32Value: 32,
		Int64Value: 64,
	}

	data2 := models.DataTypes{
		IntValue:   43,
		Int8Value:  9,
		Int16Value: 17,
		Int32Value: 33,
		Int64Value: 65,
	}

	// Act
	diffs := FindDifferences(data1, data2)

	// Assert
	assert.Len(t, diffs, 5, "Should find differences in all integer fields")

	pathsFound := make(map[string]models.FieldDiff)
	for _, diff := range diffs {
		pathsFound[diff.Path] = diff
	}

	// Verify specific types are preserved
	if intDiff, exists := pathsFound["IntValue"]; exists {
		assert.Equal(t, 42, intDiff.Expected)
		assert.Equal(t, 43, intDiff.Actual)
		assert.IsType(t, int(0), intDiff.Expected)
	}

	if int8Diff, exists := pathsFound["Int8Value"]; exists {
		assert.Equal(t, int8(8), int8Diff.Expected)
		assert.Equal(t, int8(9), int8Diff.Actual)
		assert.IsType(t, int8(0), int8Diff.Expected)
	}
}

func TestFindDifferences_UnsignedIntegerTypes_ShouldDetectCorrectly(t *testing.T) {
	// Arrange
	data1 := models.DataTypes{
		UintValue:   uint(100),
		Uint8Value:  uint8(200),
		Uint16Value: uint16(300),
		Uint32Value: uint32(400),
		Uint64Value: uint64(500),
	}

	data2 := models.DataTypes{
		UintValue:   uint(101),
		Uint8Value:  uint8(201),
		Uint16Value: uint16(301),
		Uint32Value: uint32(401),
		Uint64Value: uint64(501),
	}

	// Act
	diffs := FindDifferences(data1, data2)

	// Assert
	assert.Len(t, diffs, 5, "Should find differences in all unsigned integer fields")

	pathsFound := make(map[string]models.FieldDiff)
	for _, diff := range diffs {
		pathsFound[diff.Path] = diff
	}

	if uintDiff, exists := pathsFound["UintValue"]; exists {
		assert.Equal(t, uint(100), uintDiff.Expected)
		assert.Equal(t, uint(101), uintDiff.Actual)
		assert.IsType(t, uint(0), uintDiff.Expected)
	}
}

func TestFindDifferences_FloatTypes_ShouldDetectCorrectly(t *testing.T) {
	// Arrange
	data1 := models.DataTypes{
		Float32Value: float32(3.14),
		Float64Value: float64(2.71828),
	}

	data2 := models.DataTypes{
		Float32Value: float32(3.15),
		Float64Value: float64(2.71829),
	}

	// Act
	diffs := FindDifferences(data1, data2)

	// Assert
	assert.Len(t, diffs, 2, "Should find differences in float fields")

	pathsFound := make(map[string]models.FieldDiff)
	for _, diff := range diffs {
		pathsFound[diff.Path] = diff
	}

	if float32Diff, exists := pathsFound["Float32Value"]; exists {
		assert.Equal(t, float32(3.14), float32Diff.Expected)
		assert.Equal(t, float32(3.15), float32Diff.Actual)
		assert.IsType(t, float32(0), float32Diff.Expected)
	}
}

func TestFindDifferences_ZeroVsNonZeroIntegers_ShouldDetectCorrectly(t *testing.T) {
	// Arrange
	zero := models.DataTypes{} // All zero values
	nonZero := models.DataTypes{
		IntValue:     1,
		Int8Value:    1,
		Int16Value:   1,
		Int32Value:   1,
		Int64Value:   1,
		UintValue:    1,
		Uint8Value:   1,
		Uint16Value:  1,
		Uint32Value:  1,
		Uint64Value:  1,
		Float32Value: 1.0,
		Float64Value: 1.0,
		BoolValue:    true,
		StringValue:  "test",
	}

	// Act
	diffs := FindDifferences(zero, nonZero)

	// Assert
	assert.Len(t, diffs, 14, "Should find differences in all non-zero fields")
}

// ===== TESTS FOR MAPS =====

func TestFindDifferences_SimpleStringMaps_ShouldDetectCorrectly(t *testing.T) {
	// Arrange
	container1 := models.MapContainer{
		StringMap: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	container2 := models.MapContainer{
		StringMap: map[string]string{
			"key1": "value1_modified",
			"key2": "value2",
		},
	}

	// Act
	diffs := FindDifferences(container1, container2)

	// Assert
	assert.Len(t, diffs, 1, "Should find exactly one difference in map")
	assert.Equal(t, "StringMap.[key1]", diffs[0].Path)
	assert.Equal(t, "value1", diffs[0].Expected)
	assert.Equal(t, "value1_modified", diffs[0].Actual)
}

func TestFindDifferences_MapsDifferentKeys_ShouldDetectCorrectly(t *testing.T) {
	// Arrange
	container1 := models.MapContainer{
		StringMap: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	container2 := models.MapContainer{
		StringMap: map[string]string{
			"key1": "value1",
			"key3": "value3", // Key different
		},
	}

	// Act
	diffs := FindDifferences(container1, container2)

	// Assert
	assert.Len(t, diffs, 1, "Should detect maps with different keys as different")
	assert.Equal(t, "StringMap.[key2]", diffs[0].Path)
}

func TestFindDifferences_NilVsEmptyMap_ShouldDetectDifference(t *testing.T) {
	// Arrange
	container1 := models.MapContainer{
		StringMap: nil,
	}

	container2 := models.MapContainer{
		StringMap: make(map[string]string),
	}

	// Act
	diffs := FindDifferences(container1, container2)

	// Assert
	assert.Len(t, diffs, 1, "Should detect difference between nil and empty map")
	assert.Equal(t, "StringMap", diffs[0].Path)
}

func TestFindDifferences_NestedMaps_ShouldDetectCorrectly(t *testing.T) {
	// Arrange
	container1 := models.MapContainer{
		NestedMap: map[string]map[string]int{
			"group1": {
				"item1": 10,
				"item2": 20,
			},
		},
	}

	container2 := models.MapContainer{
		NestedMap: map[string]map[string]int{
			"group1": {
				"item1": 15, // Value modified
				"item2": 20,
			},
		},
	}

	// Act
	diffs := FindDifferences(container1, container2)

	// Assert
	assert.Len(t, diffs, 1, "Should find difference in nested map")
	assert.Equal(t, "NestedMap.[group1].[item1]", diffs[0].Path)
	assert.Equal(t, 10, diffs[0].Expected)
	assert.Equal(t, 15, diffs[0].Actual)
}

func TestFindDifferences_MapWithComplexValues_ShouldDetectCorrectly(t *testing.T) {
	// Arrange
	container1 := models.MapContainer{
		PersonMap: map[string]models.Person{
			"person1": {
				ID:   1,
				Name: "Alice",
			},
		},
	}

	container2 := models.MapContainer{
		PersonMap: map[string]models.Person{
			"person1": {
				ID:   1,
				Name: "Bob", // Name modified
			},
		},
	}

	// Act
	diffs := FindDifferences(container1, container2)

	// Assert
	assert.Len(t, diffs, 1, "Should find difference in complex map value")
	assert.Equal(t, "PersonMap.[person1].Name", diffs[0].Path)
	assert.Equal(t, "Alice", diffs[0].Expected)
	assert.Equal(t, "Bob", diffs[0].Actual)
}

func TestFindDifferences_EmptyMaps_ShouldReturnNoDifferences(t *testing.T) {
	// Arrange
	container1 := models.MapContainer{
		StringMap: make(map[string]string),
		IntMap:    make(map[string]int),
	}

	container2 := models.MapContainer{
		StringMap: make(map[string]string),
		IntMap:    make(map[string]int),
	}

	// Act
	diffs := FindDifferences(container1, container2)

	// Assert
	assert.Empty(t, diffs, "Empty maps should have no differences")
}
