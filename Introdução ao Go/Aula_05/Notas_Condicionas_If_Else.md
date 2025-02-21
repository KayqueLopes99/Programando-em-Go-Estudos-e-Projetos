## Operadores lógicos
- Os operadores lógicos são && (E), || (OU) e ! (NÃO). Eles são usados para comparar valores e avaliam expressões em valores booleanos, retornando true ou false. 

## **Estruturas de Controle em Go: `if`, `else if` e `else`**  

## **1. `if`: Condição Simples**
- O `if` executa um bloco de código se a condição for verdadeira.

- Sintaxe:
```go
if condição {
    // Código executado se a condição for verdadeira
}
```


+ **2. `if` com Inicialização de Variável**
```go
if num := 10; num > 5 {
    fmt.Println("Número maior que 5.")
}
// Aqui, 'num' não está mais acessível
```


## **3. `if...else`: Condição com Alternativa**
- Se a condição do `if` for falsa, o bloco dentro de `else` será executado.
- O `else if` permite testar múltiplas condições.

### **Exemplo:**
```go
package main

import "fmt"

func main() {
    nota := 75

    if nota >= 90 {
        fmt.Println("Nota A")
    } else if nota >= 80 {
        fmt.Println("Nota B")
    } else if nota >= 70 {
        fmt.Println("Nota C")
    } else {
        fmt.Println("Reprovado")
    }
}
```
