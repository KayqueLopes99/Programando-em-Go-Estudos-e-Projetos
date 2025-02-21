package past

import (
	"fmt"
)
type Endereco struct {
	Numero int
	Zona string
	Local string
	
}

// Struct com dados você pode colocar dados depois.
type Cliente struct {
	Nome string
	Idade int
	Endereco
	Email string
}


func Main_Struct(){
	fmt.Println("Clientes:")

	cliente_01 := Cliente{
		Nome: "Kayque",
		Idade: 20,
		Endereco: Endereco{
			Numero: 230,
			Zona: "Rural",
			Local: "Próximo de Vênus",
		},
		Email: "Kaiquelopes.601@gmail.com",
	}

	cliente_02 := Cliente{
		Nome: "Isabelly",
		Idade: 21,
		Endereco: Endereco{
			Numero: 231,
			Zona: "Rural",
			Local: "Próximo do Sol",
		},
		Email: "Isa.601@gmail.com",
	}

	fmt.Println("Cliente:", cliente_01)
	fmt.Println("Cliente:", cliente_02)

	
	
}