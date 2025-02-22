## Métodos
- São semelhantes a funções, mas são associados a **structs** ou outros tipos. 
- Eles permitem que uma struct tenha **comportamentos específicos**, como os métodos de classes em outras linguagens.  

- 1. Criando um Método em uma Struct 
```go
package main

import "fmt"

// Definição da struct
type Pessoa struct {
    Nome  string
    Idade int
}

// Método associado à struct Pessoa
func (p Pessoa) apresentar() {
    fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.Nome, p.Idade)
}

func main() {
    pessoa1 := Pessoa{"Lucas", 25}
    pessoa1.apresentar() // Chamando o método
}
```
---

- Método Com Ponteiro (`*`) para Modificar Valores
- Se usarmos um método com um **receiver por valor (`p Pessoa`)**, a struct **não será modificada**.  
- Para **modificar a struct**, usamos um **ponteiro (`*Pessoa`)**.  
