package main

import (
	"fmt"

	"github.com/seu-usuario/meu-projeto/models"
)

func main() {
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

	for _, pessoa := range pessoas {
		fmt.Println(formatClean(pessoa))
	}

	diffs := FindDifferences(expected, actual)

	fmt.Println("Field differences:")
	for _, diff := range diffs {
		fmt.Printf("  └─ %s: %q ≠ %q\n", diff.Path, diff.Expected, diff.Actual)
	}
}
