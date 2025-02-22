package main

import (
	"fmt"

	"github.com/kayquelopes9/aula_07/past"
	// "github.com/kayquelopes9/aula_07/past"
)

func main(){
	fmt.Println("Ponteiros:")
	past.Main_Ponteiro()
	fmt.Println("Goroutinnes:")
	past.Main_Goreotine()
	past.Main_Goreotine_Sem_Ordem()
	past.Main_Channels()
	past.Main_Channel()


}