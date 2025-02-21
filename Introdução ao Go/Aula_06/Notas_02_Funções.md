## **Funções em Go** 🚀  
- São blocos de código reutilizáveis que podem receber **parâmetros** e retornar **valores**.  

- Sintaxe:  
Uma função em Go tem a seguinte estrutura:  
```go
func nomeDaFuncao(parametro Tipo) TipoDeRetorno {
    // Código da função
    return valor
}
```
##  Funções Anônimas (Lambda)
Podemos criar funções **sem nome** e armazená-las em variáveis.

```go
package main

import "fmt"

func main() {
    saudacao := func(nome string) {
        fmt.Println("Olá,", nome)
    }

    saudacao("Maria")
}
