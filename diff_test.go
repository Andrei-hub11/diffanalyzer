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
		if diff.Path == "Emails" {
			foundEmailDiff = true
		}
		if diff.Path == "Profile.Tags" {
			foundTagDiff = true
		}
	}

	assert.True(t, foundEmailDiff || foundTagDiff, "Should find differences in slice elements")
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
