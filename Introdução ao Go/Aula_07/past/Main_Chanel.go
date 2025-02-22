package past

import (
	"fmt"
	"time"
)

// Algoritmos de concorrencia.
func producer(c chan int){
	for i := 0; i < 6; i++{
		c <- i

	}
	close(c)
}


func consumer(c chan int){
	for v := range c{
		fmt.Println(v)

	}
	fmt.Println("Consumer finalizada!")
}

// Escrevendo channels <- Valor Lendo Valor <- channels
func Main_Channel() {
    channels := make(chan int) // CANAL DE int.
	time.Sleep(1 * time.Second)
	fmt.Println("Outro modelo:")
	go producer(channels)
	go consumer(channels)
	go consumer(channels)
	time.Sleep(1 * time.Second)
}
