package past

import (
	"fmt"
)

//  funct nome parâmetro retorno{}
func soma(a, b int) int {
	return a + b
}

func subtracao(a, b int) int {
	return a - b
}

func multiplicacao(a, b int) int {
	return a * b
}

func divisao(a, b int) int {
	return a / b
}




func Main_Funcion(){	
	fmt.Println("Calculus:")
	fmt.Println(
		soma(2,8),
		subtracao(2,8),
		multiplicacao(2,8),
		divisao(2,8),
	)
	potencia := func (a int) int{
		return a * a
	}
	fmt.Println(potencia(9))


}