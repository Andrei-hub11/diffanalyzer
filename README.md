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
    classDef structBox fill:#1a365d,stroke:#ffffff,stroke-width:2px,color:#ffffff
    classDef stringBox fill:#553c9a,stroke:#ffffff,stroke-width:2px,color:#ffffff
    classDef sliceBox fill:#1e4d72,stroke:#ffffff,stroke-width:2px,color:#ffffff
    classDef diffBox fill:#c53030,stroke:#ffffff,stroke-width:2px,color:#ffffff
    classDef processBox fill:#2d5a87,stroke:#ffffff,stroke-width:2px,color:#ffffff
    classDef decisionBox fill:#4a5568,stroke:#ffffff,stroke-width:2px,color:#ffffff
    classDef actionBox fill:#38a169,stroke:#ffffff,stroke-width:2px,color:#ffffff

    class A,B,I,Q processBox
    class C,E,J,N decisionBox
    class D,K,O diffBox
    class F,G,H,L,M,P actionBox
    class R,V,Y structBox
    class S,T,W,Z,AA stringBox
    class U,X sliceBox
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
