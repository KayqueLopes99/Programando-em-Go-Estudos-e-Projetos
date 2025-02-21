
## **Repetição com `for` em Go (Simulando `while`)**  

- `for` é a única estrutura de repetição**. 
- Não existe `while` ou `do-while`, mas podemos usar o `for` de diferentes formas para simular esses comportamentos.  

- Sintaxe:
- O `for` tradicional tem três partes:  
1. **Inicialização** : Define uma variável de controle.  
2. **Condição** : O laço continua enquanto essa condição for verdadeira.  
3. **Incremento** : Atualiza a variável de controle.  

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println("Número:", i)
    }
}
```


## **2. `for` como `while` (Sem Inicialização e Incremento)**
- Para simular um `while`, omitimos a **inicialização** e o **incremento**, deixando apenas a **condição**.

```go
package main

import "fmt"

func main() {
    numero := 1

    for numero <= 5 { // Enquanto `numero` for menor ou igual a 5
        fmt.Println("Número:", numero)
        numero++ // Incremento dentro do loop
    }
}
```

## **3. `continue` e `break`**
- `break` : Sai do loop imediatamente.  
- `continue` : Pula para a próxima iteração, ignorando o código abaixo dele.  


## Range:
- **percorrer elementos de estruturas de dados** como **arrays, slices, maps e strings**. Ele retorna **dois valores**:  
1. O **índice** (posição) no caso de arrays, slices e strings.  
2. A **chave** no caso de maps.  
3. O **valor** correspondente ao índice ou chave.  

- O `range` percorre um **array ou slice**, retornando o **índice** e o **valor**.

```go
package main

import "fmt"

func main() {
    numeros := []int{10, 20, 30, 40, 50}

    for indice, valor := range numeros {
        fmt.Printf("Índice: %d, Valor: %d\n", indice, valor)
    }
}
