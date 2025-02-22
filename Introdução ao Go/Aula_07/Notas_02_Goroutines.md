## Goroutines com concorrencia e sem ordem e channels
### 0.1 - Goroutines e Concorrência em Go
-  Uma **Goroutine** é uma **função leve** que roda **concorrentemente** (em paralelo com o código principal).  
- Ela é criada com a palavra-chave **`go`** antes da chamada da função.  

```go
package main

import (
    "fmt"
    "time"
)

func imprimirMensagem() {
    fmt.Println("Olá do Goroutine!")
}

func main() {
    go imprimirMensagem() // Executa a função como Goroutine

    fmt.Println("Executando no main...")

    time.Sleep(time.Second) // Espera 1 segundo para a Goroutine terminar
}
```


### 0.2 Goroutines Sem Ordem
- Como Goroutines rodam de forma independente, **não há garantia de ordem de execução**.   
- A ordem dos prints pode variar a cada execução.  

### 0.3 **Channels em Go**  
- Goroutines executam de forma **concorrente**, mas para **sincronizar e trocar informações** entre elas, usamos **Channels**.  

- Um **Channel** (`chan`) permite que uma Goroutine **envie e receba valores** de outra Goroutine.  

- Sintaxe: 
```go
canal := make(chan int) // Cria um canal que transmite inteiros

// Escrevendo channels <- Valor || Lendo Valor <- channels
```
- Enviando e Recebendo Dados com Channel**
### **Exemplo: Comunicação Entre Goroutines**
```go
package main

import "fmt"

func enviarMensagem(canal chan string) {
    canal <- "Olá, Channel!" // Envia mensagem para o canal
}

func main() {
    canal := make(chan string) // Criação do canal

    go enviarMensagem(canal) // Goroutine enviando dados

    mensagem := <-canal // Recebe dados do canal
    fmt.Println(mensagem) // "Olá, Channel!"
}
```
🔹 **Explicação:**  
- `canal := make(chan string)`: Cria um **canal de strings**.  
- `canal <- "Olá, Channel!"`: **Envia** o valor `"Olá, Channel!"` para o canal.  
- `mensagem := <-canal`: **Recebe** o valor do canal.  
- O programa **espera** até que o valor seja recebido antes de continuar!  
