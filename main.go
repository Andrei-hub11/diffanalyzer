package main

import (
	"fmt"

	"github.com/seu-usuario/meu-projeto/models"
)

func main() {
	fmt.Println("=== EXAMPLE 1: Basic Person Comparison ===")

	expected := models.Person{
		ID:     1,
		Name:   "Alice",
		Emails: []string{"alice@company.com", "alice@personal.com"},
		Profile: models.Profile{
			Bio:  "Engineer",
			Tags: []string{"go", "backend", "api"},
			Address: models.Address{
				City:    "São Paulo",
				Country: "Brasil",
			},
		},
	}

	actual := models.Person{
		ID:     1,
		Name:   "Alice",
		Emails: []string{"alice@company.com", "alice@gmail.com"},
		Profile: models.Profile{
			Bio:  "Developer",
			Tags: []string{"go", "frontend", "api"},
			Address: models.Address{
				City:    "São Paulo",
				Country: "Brazil",
			},
		},
	}

	diffs := FindDifferences(expected, actual)
	printDifferences("Person Comparison", diffs)

	fmt.Println("\n=== EXAMPLE 2: Data Types Comparison ===")

	data1 := models.DataTypes{
		IntValue:     42,
		Int8Value:    8,
		BoolValue:    true,
		Float32Value: 3.14,
		StringValue:  "hello",
	}

	data2 := models.DataTypes{
		IntValue:     43,
		Int8Value:    8,
		BoolValue:    false,
		Float32Value: 3.15,
		StringValue:  "world",
	}

	diffs2 := FindDifferences(data1, data2)
	printDifferences("Data Types Comparison", diffs2)

	fmt.Println("\n=== EXAMPLE 3: Map Comparison ===")

	container1 := models.MapContainer{
		StringMap: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		},
		IntMap: map[string]int{
			"count": 10,
			"total": 100,
		},
	}

	container2 := models.MapContainer{
		StringMap: map[string]string{
			"key1": "value1_modified",
			"key2": "value2",
			"key4": "new_value", // new key
		},
		IntMap: map[string]int{
			"count": 15, // value modified
			"total": 100,
		},
	}

	diffs3 := FindDifferences(container1, container2)
	printDifferences("Map Comparison", diffs3)

	fmt.Println("\n=== EXAMPLE 4: Nested Map Comparison ===")

	nested1 := models.MapContainer{
		NestedMap: map[string]map[string]int{
			"group1": {
				"item1": 10,
				"item2": 20,
			},
			"group2": {
				"item3": 30,
			},
		},
	}

	nested2 := models.MapContainer{
		NestedMap: map[string]map[string]int{
			"group1": {
				"item1": 15, // modified
				"item2": 20,
			},
			"group2": {
				"item3": 30,
				"item4": 40, // new item
			},
		},
	}

	diffs4 := FindDifferences(nested1, nested2)
	printDifferences("Nested Map Comparison", diffs4)

	fmt.Println("\n=== EXAMPLE 5: Slice Length Differences ===")

	person1 := models.Person{
		ID:     1,
		Name:   "Bob",
		Emails: []string{"bob@company.com", "bob@personal.com"},
	}

	person2 := models.Person{
		ID:     1,
		Name:   "Bob",
		Emails: []string{"bob@company.com"}, // slice smaller
	}

	diffs5 := FindDifferences(person1, person2)
	printDifferences("Slice Length Comparison", diffs5)

	fmt.Println("\n=== EXAMPLE 6: Nil vs Empty Comparison ===")

	nilPerson := models.Person{
		ID:     1,
		Name:   "Test",
		Emails: nil, // nil slice
	}

	emptyPerson := models.Person{
		ID:     1,
		Name:   "Test",
		Emails: []string{}, // empty slice
	}

	diffs6 := FindDifferences(nilPerson, emptyPerson)
	printDifferences("Nil vs Empty Comparison", diffs6)

	fmt.Println("\n=== EXAMPLE 7: Complex Person with Maps ===")

	complexContainer1 := models.MapContainer{
		PersonMap: map[string]models.Person{
			"employee1": {
				ID:   1,
				Name: "John",
				Profile: models.Profile{
					Bio: "Senior Developer",
					Address: models.Address{
						City:    "New York",
						Country: "USA",
					},
				},
			},
		},
	}

	complexContainer2 := models.MapContainer{
		PersonMap: map[string]models.Person{
			"employee1": {
				ID:   1,
				Name: "John Smith", // name modified
				Profile: models.Profile{
					Bio: "Lead Developer", // bio modified
					Address: models.Address{
						City:    "San Francisco", // city modified
						Country: "USA",
					},
				},
			},
		},
	}

	diffs7 := FindDifferences(complexContainer1, complexContainer2)
	printDifferences("Complex Person Map Comparison", diffs7)

	fmt.Println("\n=== EXAMPLE 8: Pessoa (Portuguese) ===")

	pessoas := []models.Pessoa{
		{
			Nome:   "João",
			Idade:  30,
			Ativo:  true,
			Emails: []string{"joao@example.com", "joao@empresa.com"},
		},
		{
			Nome:   "Maria",
			Idade:  25,
			Ativo:  false,
			Emails: []string{"maria@example.com"},
		},
	}

	for i, pessoa := range pessoas {
		fmt.Printf("Pessoa %d: %s\n", i+1, formatClean(pessoa))
	}

	pessoaModificada := models.Pessoa{
		Nome:   "João Silva",                    // name modified
		Idade:  31,                              // age modified
		Ativo:  false,                           // status modified
		Emails: []string{"joao@newcompany.com"}, // email modified
	}

	diffs8 := FindDifferences(pessoas[0], pessoaModificada)
	printDifferences("Pessoa Comparison", diffs8)

	fmt.Println("\n=== EXAMPLE 9: Slice of Structs Comparison ===")

	expectedCollection := models.ItemCollection{
		Items: []models.Item{
			{ID: 1, Status: "active", Value: 100},
			{ID: 3, Status: "active", Value: 150},
			{ID: 5, Status: "active", Value: 300},
		},
	}

	actualCollection := models.ItemCollection{
		Items: []models.Item{
			{ID: 1, Status: "active", Value: 100},
			{ID: 3, Status: "active", Value: 140}, // Value changed
			{ID: 5, Status: "active", Value: 250}, // Value changed
		},
	}

	diffs9 := FindDifferences(expectedCollection, actualCollection)
	printDifferences("Slice of Structs Comparison", diffs9)

	fmt.Println("\n=== EXAMPLE 10: Direct Slice Comparison ===")

	expectedItems := []models.Item{
		{ID: 1, Status: "active", Value: 100},
		{ID: 3, Status: "active", Value: 150},
		{ID: 5, Status: "active", Value: 300},
	}

	actualItems := []models.Item{
		{ID: 1, Status: "active", Value: 100},
		{ID: 3, Status: "active", Value: 140}, // Value changed
		{ID: 5, Status: "active", Value: 250}, // Value changed
	}

	diffs10 := FindDifferences(expectedItems, actualItems)
	printDifferences("Direct Slice Comparison", diffs10)
}

func printDifferences(title string, diffs []models.FieldDiff) {
	fmt.Printf("\n%s:\n", title)
	if len(diffs) == 0 {
		fmt.Println("  No differences found!")
		return
	}

	fmt.Printf("  Found %d difference(s):\n", len(diffs))
	for _, diff := range diffs {
		expectedStr := formatDiffValue(diff.Expected)
		actualStr := formatDiffValue(diff.Actual)
		fmt.Printf("  └─ %s: %s ≠ %s\n", diff.Path, expectedStr, actualStr)
	}
}

func formatDiffValue(value interface{}) string {
	if value == nil {
		return "<nil>"
	}

	switch v := value.(type) {
	case string:
		return fmt.Sprintf("%q", v)
	case bool:
		return fmt.Sprintf("%t", v)
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%v", v)
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%v", v)
	case float32, float64:
		return fmt.Sprintf("%v", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
