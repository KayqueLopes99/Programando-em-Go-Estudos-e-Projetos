## Ponteiros - ***point**
- O Endereço de um Dado(Valor) que será armazenado em local da memória(Apontar).
- Usados para armazenar endereços de memória. Isso permite **modificar valores diretamente na memória** e evitar cópias desnecessárias de dados.  
-  Um ponteiro é uma variável que armazena o **endereço de memória** de outra variável.  
### **Declaração de um Ponteiro**
```go
var p *int  // Ponteiro para um inteiro
```
- `*int` indica um ponteiro para um `int`.  
- `p` pode armazenar o **endereço de um inteiro**.  
+ Obtendo e Utilizando Ponteiros**
- Para **pegar o endereço de memória** de uma variável, usamos `&`.  
- Para **acessar o valor armazenado no endereço** do ponteiro, usamos `*`.  

### **Exemplo: Criando e Usando Ponteiros**
```go
package main

import "fmt"

func main() {
    var numero int = 42
    var p *int = &numero // Pegamos o endereço de memória de 'numero'

    fmt.Println("Valor de numero:", numero)    // 42
    fmt.Println("Endereço de numero:", p)      // Exibe o endereço de memória
    fmt.Println("Valor apontado por p:", *p)   // 42 (desreferenciação)
}
```

- Se passarmos um ponteiro para uma função, podemos modificar o valor original da variável.  