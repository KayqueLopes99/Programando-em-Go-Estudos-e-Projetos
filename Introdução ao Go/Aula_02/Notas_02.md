## Variáveis e constantes para o armazenamento de dados. 
---
### Variáveis
#### **Declaração de variáveis em Go:**

+ **Declaração com valor inicial:**
- Quando você não especifica um valor, Go atribui um valor zero para o tipo da variável (valores padrão).
- Exemplo: 

```go
var x int  // x recebe 0 por padrão, pois é um tipo inteiro
    var name string  // name recebe "" (string vazia)
```

+ **Declaração curta (em um escopo de função):**
- Go permite uma forma mais curta de declarar e inicializar variáveis. Isso é útil dentro de funções.
- Sintaxe: 
     
```go
x := 10   // Declare e atribua valor, sem usar "var"
name := "John" // "name" é uma string, sem especificar tipo
```
- Go determina automaticamente o tipo da variável com base no valor atribuído (inferência de tipo).

---

### Constantes

- As constantes em Go são valores imutáveis após sua declaração e devem ser atribuídas com um valor durante a declaração.

#### **Declaração de constantes em Go:**

+ 1. **Declaração de uma constante com tipo especificado:**
- Pode declarar uma constante passando o tipo e valor explicitamente.
- Sintaxe:

```go
const PI float64 = 3.14159
```

+ 2. **Declaração sem tipo:**
- Pode omitir o tipo e Go deduzirá automaticamente o tipo da constante a partir do valor fornecido.
- Sintaxe:
```go
const name = "Golang"  // Go deduz que é uma string
```

---

```go
package main

import "fmt"

func main() {
    // Declarando variáveis
    var x int  // x recebe valor 0
    y := 15    // y recebe 15, tipo int inferido automaticamente
    name := "John"
    
    // Exibindo valores
    fmt.Println("x:", x)    // 0
    fmt.Println("y:", y)    // 15
    fmt.Println("name:", name) // John
    
    // Declarando constante
    const PI = 3.14159   // Go deduz que PI é do tipo float64
    
    // Exibindo constante
    fmt.Println("PI:", PI) // 3.14159
}
```
