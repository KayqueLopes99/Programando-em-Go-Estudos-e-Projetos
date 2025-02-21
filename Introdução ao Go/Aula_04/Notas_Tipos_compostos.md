## Tipos Compostos em Go

- Os tipos compostos em Go permitem armazenar múltiplos valores dentro de uma única variável. 
- Os principais tipos compostos são **arrays, slices e maps**.

## **1. Arrays** (Fixos, Leitura e Inserção)

- Um **array** é uma estrutura de tamanho fixo que armazena elementos do mesmo tipo. 
- O tamanho do array faz parte do seu tipo e não pode ser alterado depois de declarado.

- Sintaxe de Declaração
```go
var numeros [5]int // Array de 5 inteiros (todos iniciam com 0)
```

+ Inserção 
- Os elementos podem ser acessados ou modificados pelo índice (começando de `0`).
```go
numeros[0] = 10 // Atribui o valor 10 ao primeiro elemento
numeros[1] = 20
fmt.Println(numeros) // Saída: [10 20 0 0 0]
```

+ Leitura de Elementos
```go
fmt.Println(numeros[0]) // Lê o primeiro elemento -> Saída: 10
```

+ Inicialização Direta
```go
nomes := [3]string{"Alice", "Bob", "Carol"}
fmt.Println(nomes) // Saída: [Alice Bob Carol]
```

---

## **2. Slices** (Dinâmico: Inserção, Divisão e Remoção)

- Diferente dos arrays, os **slices** têm tamanho dinâmico e são mais utilizados em Go.
- Um slice é baseado em um array interno, mas permite crescer ou diminuir conforme necessário.

+ Criando um Slice
```go
var numeros []int // Declara um slice de inteiros (inicialmente vazio)
```
- Ou com inicialização:
```go
numeros := []int{1, 2, 3, 4, 5}
fmt.Println(numeros) // Saída: [1 2 3 4 5]
```

+ Inserção de Elementos
- Usa a função `append()` para adicionar elementos a um slice:
```go
numeros = append(numeros, 6) // Adiciona o número 6 ao slice
fmt.Println(numeros) // Saída: [1 2 3 4 5 6]
```

+ Divisão (Slicing): Podemos extrair partes de um slice usando a sintaxe `slice[início:fim]`:
```go
parte := numeros[1:4] // Elementos do índice 1 ao 3 (exclui o 4)
fmt.Println(parte) // Saída: [2 3 4]
```

+ Remoção de Elemento
- Como Go não tem uma função específica para remover elementos, usamos `append()` combinando slices:
```go
indice := 2
numeros = append(numeros[:indice], numeros[indice+1:]...)
fmt.Println(numeros) // Remove o elemento no índice 2
```

---

## **3. Maps** (Chave-Valor: Inserção, Remoção e Leitura)

- estrutura de dados chave-valor. 
- A chave pode ser de qualquer tipo comparável (`string`, `int`, etc.), e o valor pode ser qualquer tipo.

+ Criando um Map
```go
pessoas := make(map[string]int) // Cria um map com chave string e valor int
```
- Ou inicializando com valores:
```go
pessoas := map[string]int{
    "Alice": 25,
    "Bob":   30,
}
fmt.Println(pessoas) // Saída: map[Alice:25 Bob:30]
```

+ Inserção e Atualização
```go
pessoas["Carol"] = 28 // Adiciona um novo elemento
fmt.Println(pessoas) // Saída: map[Alice:25 Bob:30 Carol:28]
```

- Leitura de Elemento
```go
idade := pessoas["Alice"]
fmt.Println(idade) // Saída: 25
```
- Se a chave não existir, o Go retorna o valor padrão do tipo (`0` para `int`, `""` para `string`, etc.).

### **Verificando se uma Chave Existe**
```go
idade, existe := pessoas["David"]
if existe {
    fmt.Println("Idade de David:", idade)
} else {
    fmt.Println("David não encontrado") // Saída: David não encontrado
}
```

+ Remoção de Elemento
-nUsamos a função `delete()` para remover um item do map:
```go
delete(pessoas, "Bob")
fmt.Println(pessoas) // Saída: map[Alice:25 Carol:28]
```