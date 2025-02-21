## Structs:
- As **structs** são usadas para **agrupar dados** relacionados em um único tipo. Elas são como "objetos" em outras linguagens, mas sem métodos embutidos.  

- A sintaxe:  
```go
type NomeDaStruct struct {
    Campo1 Tipo
    Campo2 Tipo
}
```


- Alterando Valores em uma Struct
```go
package main

import "fmt"

type Pessoa struct {
    Nome  string
    Idade int
}

func main() {
    p := Pessoa{"Carlos", 30}

    // Alterando valores
    p.Nome = "Roberto"
    p.Idade = 40

    fmt.Println("Nome atualizado:", p.Nome)
    fmt.Println("Idade atualizada:", p.Idade)
}
```


- Aninhamento de Structs (Struct Dentro de Struct)
- Podemos **colocar uma `struct` dentro de outra** para representar relações mais complexas.  

```go
package main

import "fmt"

// Struct Endereco
type Endereco struct {
    Rua    string
    Cidade string
}

// Struct Pessoa que contém um Endereco
type Pessoa struct {
    Nome  string
    Idade int
    Endereco
}

func main() {
    p := Pessoa{
        Nome:  "Lucas",
        Idade: 28,
        Endereco: Endereco{
            Rua:    "Rua das Flores",
            Cidade: "São Paulo",
        },
    }

    // Acessando os campos da struct aninhada
    fmt.Println("Nome:", p.Nome)
    fmt.Println("Cidade:", p.Cidade) // Podemos acessar diretamente
    fmt.Println("Rua:", p.Endereco.Rua) // Ou referenciar explicitamente
}
```