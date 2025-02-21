package past
import (
	"fmt"
)

func Main_if(){
	nota_01 := 100
	nota_02 := 100
	nota_03 := 90
	media := (nota_01 + nota_02 + nota_03)/3

	players := map[string]int{
		"Kayque": 90,
	}

	// Se está na estrutura.
	if value, ok := players["Kayque"]; ok{
		fmt.Println("Pontos:", value, ok)

	}


	if media >= 70 && media <= 100{
		fmt.Println("Aprovado!")
	}else if media < 70 && media > 50{
		fmt.Println("Recuperação!")
	}else{
		fmt.Println("Reprovado!")
	}

	fmt.Println("Média:", media)

}
