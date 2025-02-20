// Comentário.
package main // Pacote com um conjunto de bibliotecas.

import (
	"fmt"
    "github.com/kayquelopes9/introducao/meet"
)

// Ponto de Partida.
// go run name.go <---
func main() {
	fmt.Println("Hello World!")
	meet.SayHello()  // Chama a função SayHello do pacote meet
	meet.Say("Olá kayque")  // Chama a função Say passando uma mensagem
}