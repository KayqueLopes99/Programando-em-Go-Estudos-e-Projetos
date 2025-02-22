package past

import (
    "fmt"
    "time"
)

func imprimirMensagem() {
    fmt.Println("Olá do Goroutine!")
}

func Main_Goreotine() {
    go imprimirMensagem() // Executa a função como Goroutine

    fmt.Println("Executando no main...")

    time.Sleep(time.Second) // Espera 1 segundo para a Goroutine terminar
}