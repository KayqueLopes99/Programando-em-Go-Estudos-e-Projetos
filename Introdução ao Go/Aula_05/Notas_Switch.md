## `switch`: Controle de Fluxo Condicional**  

- O `switch` em Go é uma alternativa mais elegante ao `if...else if...else`, especialmente quando há múltiplas condições a serem testadas. Ele permite comparar um valor contra diferentes casos (`case`) e executar o bloco correspondente.  

- Sintaxe 
```go
switch expressão {
case valor1:
    // Código executado se expressão == valor1
case valor2:
    // Código executado se expressão == valor2
default:
    // Código executado se nenhum case for correspondente
}
```

---

```go
package main

import "fmt"

func main() {
    dia := "segunda"

    switch dia {
    case "segunda":
        fmt.Println("Início da semana!")
    case "sexta":
        fmt.Println("Sextou!")
    case "sábado", "domingo":
        fmt.Println("Fim de semana!")
    default:
        fmt.Println("Dia comum.")
    }
}
 
- O `default` é opcional e é executado caso nenhum `case` seja atendido.
- Se não houver uma expressão após `switch`, ele funciona como múltiplos `if-else`, com condições nos cases.
