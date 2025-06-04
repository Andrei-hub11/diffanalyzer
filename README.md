# Go Struct Diff Analyzer

Este projeto implementa um comparador de structs em Go que identifica diferenças entre objetos complexos.

## Fluxo da Função FindDifferences

```mermaid
flowchart TD
    A[FindDifferences<br/>expected, actual] --> B[compare<br/>expected, actual, path=""]

    B --> C{Tipos iguais?}
    C -->|Não| D[Adicionar diff<br/>tipos diferentes]
    C -->|Sim| E{Qual Kind?}

    E -->|Struct| F[Iterar campos<br/>do struct]
    E -->|String| G[Comparar strings]
    E -->|Slice/Array| H[Iterar elementos<br/>do slice]

    F --> I[Para cada campo:<br/>buildPath + compare recursivo]
    I --> B

    G --> J{Strings iguais?}
    J -->|Não| K[Adicionar diff<br/>string diferente]
    J -->|Sim| L[Continuar]

    H --> M[Para cada índice:<br/>DeepEqual]
    M --> N{Elementos iguais?}
    N -->|Não| O[Adicionar diff<br/>elemento diferente]
    N -->|Sim| P[Próximo elemento]

    D --> Q[Return diffs]
    K --> Q
    L --> Q
    O --> Q
    P --> Q

    subgraph "Exemplo: Person"
        R[Person struct] --> S[ID: int]
        R --> T[Name: string]
        R --> U[Emails: slice string]
        R --> V[Profile: struct]

        V --> W[Bio: string]
        V --> X[Tags: slice string]
        V --> Y[Address: struct]

        Y --> Z[City: string]
        Y --> AA[Country: string]
    end

    classDef structBox fill:#e1f5fe
    classDef stringBox fill:#f3e5f5
    classDef sliceBox fill:#e8f5e8
    classDef diffBox fill:#ffebee

    class R,V,Y structBox
    class S,T,W,Z,AA stringBox
    class U,X sliceBox
    class D,K,O diffBox
```

## Tipos de Dados Suportados

- ✅ **Struct**: Comparação recursiva de campos
- ✅ **String**: Comparação direta de valores
- ✅ **Int**: Comparação de números inteiros
- ✅ **Slice/Array**: Comparação elemento por elemento
- ✅ **Bool**: Comparação de valores booleanos
- ✅ **Float**: Comparação de números decimais
- ✅ **Ptr**: Suporte a ponteiros
- ✅ **Map**: Comparação de mapas

## Exemplo de Uso

```go
expected := Person{
    ID:     1,
    Name:   "Alice",
    Emails: []string{"alice@company.com", "alice@personal.com"},
    Profile: Profile{
        Bio:  "Engineer",
        Tags: []string{"go", "backend", "api"},
        Address: Address{
            City:    "São Paulo",
            Country: "Brasil",
        },
    },
}

actual := Person{
    ID:     1,
    Name:   "Alice",
    Emails: []string{"alice@company.com", "alice@gmail.com"},
    Profile: Profile{
        Bio:  "Developer",
        Tags: []string{"go", "frontend", "api"},
        Address: Address{
            City:    "São Paulo",
            Country: "Brazil",
        },
    },
}

diffs := FindDifferences(expected, actual)
```

## Saída Esperada

```
Field differences:
  └─ Emails: "alice@personal.com" ≠ "alice@gmail.com"
  └─ Profile.Bio: "Engineer" ≠ "Developer"
  └─ Profile.Tags: "backend" ≠ "frontend"
  └─ Profile.Address.Country: "Brasil" ≠ "Brazil"
```
