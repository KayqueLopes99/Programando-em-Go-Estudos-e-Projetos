package main

import (
	"fmt"
    "github.com/kayquelopes9/aula_03/past"
)


func main()  {
	fmt.Println("Escolha qual código executar:")
    past.MainFloat() // Executa o código de float
    past.MainInt()    // Executa o código de inteiro
}