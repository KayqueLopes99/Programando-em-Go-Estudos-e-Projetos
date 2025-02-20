## Go
- Go é uma linguagem de programação desenvolvida pela Google, também conhecida como Golang.
- Ela foi projetada para ser simples, eficiente e concisa.

### Estrutura Básica do Código

0.1 **Pacote (package)**: 
   - O Go usa pacotes para organizar o código. Um pacote é uma coleção de arquivos Go. 
   - No seu exemplo, o código começa com `package main`, que indica que esse arquivo contém a função principal (ou `main`), ponto de entrada do programa.
   - O pacote `meet` contém funções que são utilizadas em outros arquivos dentro do mesmo projeto.

0.2. **Função `main`**:
   - A função `main` é o ponto de entrada de qualquer aplicação Go. Quando o programa é executado com `go run nome-do-arquivo.go`, o Go começa a execução pela função `main`.
   - Dentro da `main`, você está chamando duas funções do pacote `meet`: `SayHello` e `Say`.

0.3. **Funções exportáveis (com letra maiúscula)**:
   - Em Go, se uma função ou variável começa com letra maiúscula, ela é **exportada**. Isso significa que ela pode ser acessada fora do pacote onde foi declarada.
   - Exemplo: `SayHello` e `Say` são funções exportadas do pacote `meet` porque começam com letra maiúscula.
   
0.4. **Funções não exportáveis (com letra minúscula)**:
   - Se uma função ou variável começa com letra minúscula, ela é **não exportada** e só pode ser usada dentro do próprio pacote.
   - No seu código, todas as funções exportadas podem ser usadas fora do pacote onde estão, mas as funções privadas só funcionariam dentro de `meet`.

0.5. **Modularização e Go Modules**:
   - O comando `go mod init github.com/kayquelopes9/introducao` inicializa um **modulo Go**, que é uma maneira de gerenciar dependências e versões no Go.
   - A modularização permite que o seu projeto cresça de forma mais organizada e que você possa dividir funcionalidades em pacotes, como você fez com o pacote `meet`.

0.5 **// Comentário.**

### Fluxo do Código:

1. **Pacote `main`**:
   - O pacote `main` importa o pacote `meet` e usa as funções `SayHello` e `Say` para imprimir mensagens no console.
   
2. **Pacote `meet`**:
   - O pacote `meet` define as funções `Say` e `SayHello`, que são usadas no pacote `main` para imprimir mensagens personalizadas.
### Como Rodar o Código
- Use o comando `go run nome-do-arquivo.go`. 


## Fmt:
- O `import "fmt"` é uma declaração que traz o pacote **fmt** para dentro do seu código.
- O **fmt** é um dos pacotes padrão do Go e é utilizado para formatar e imprimir dados na saída (geralmente no terminal ou console).

1. **`fmt.Println()`**: 
   - Imprime o que é passado como argumento, seguido por uma nova linha.
   - Exemplo:
     ```go
     fmt.Println("Hello, World!")
     ```
     Saída: `Hello, World!`

2. **`fmt.Print()`**: 
   - Imprime o que é passado como argumento, sem adicionar uma nova linha no final.
   - Exemplo:
     ```go
     fmt.Print("Hello")
     fmt.Print(" World!")
     ```
     Saída: `Hello World!`

3. **`fmt.Printf()`**: 
   - Permite formatar a saída, usando uma string de formato. Isso é útil quando você precisa imprimir variáveis de maneira mais controlada.
   - Exemplo:
     ```go
     name := "Kayque"
     age := 25
     fmt.Printf("My name is %s and I am %d years old.\n", name, age)
     ```
     Saída: `My name is Kayque and I am 25 years old.`
