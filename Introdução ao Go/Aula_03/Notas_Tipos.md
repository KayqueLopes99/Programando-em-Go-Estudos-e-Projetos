## Tipos Básicos

- 1. **`int`** (inteiro): Representa números inteiros. O tamanho de um `int` depende do sistema (32 ou 64 bits). Em Go, o tipo `int` é adequado para valores inteiros em geral, mas existem tipos específicos para inteiros com tamanho fixo.

   **Exemplo de sintaxe**:
   ```go
   var idade int = 30
   fmt.Println(idade)
   ```

- 2. **`string`** (cadeia de caracteres): Usado para armazenar texto. As strings em Go são imutáveis e representam sequências de bytes, geralmente codificadas em UTF-8.

   **Exemplo de sintaxe**:
   ```go
   var nome string = "João"
   fmt.Println(nome)
   ```

- 3. **`float`** (números de ponto flutuante): Usado para armazenar números decimais. Go tem dois tipos de `float`:
   - `float32`: precisão de 32 bits.
   - `float64`: precisão de 64 bits (o tipo `float64` é o mais comum).

   **Exemplo de sintaxe**:
   ```go
   var peso float64 = 70.5
   fmt.Println(peso)
   ```

- **`bool`** (booleano): Representa valores lógicos, `true` ou `false`.

   **Exemplo de sintaxe**:
   ```go
   var ativo bool = true
   fmt.Println(ativo)
   ```

---

### **Manipulação do Tamanho em Go**
```go
var num8 int8 = 127   // Inteiro de 8 bits
var num16 int16 = 32000 // Inteiro de 16 bits
var num32 int32 = 2147483647 // Inteiro de 32 bits
var num64 int64 = 9223372036854775807 // Inteiro de 64 bits
```
- **Tamanho Igual nas Operações**: tamanhos equivalentes para evitar problemas
