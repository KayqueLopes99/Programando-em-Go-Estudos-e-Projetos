package past

import (
	"fmt"
	"time"
)

// Escrevendo channels <- Valor Lendo Valor <- channels
func Main_Channels() {
    channels := make(chan int) // CANAL DE int.

	// Espera da leitura.
	go func(){
		// Escrever
		for i := 0; i < 5; i++{
			channels <- i
		}
		close(channels) // Fechando Sempre!!!!
		fmt.Println("Escrita Finalizada!")
	}() 

	time.Sleep(1 * time.Second)
	for valor := range channels{
		fmt.Println("Leitura:", valor)
	}

}


