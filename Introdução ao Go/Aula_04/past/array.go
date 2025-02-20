package past

import "fmt"

func MainArray(){
	var gavetas [5]string
	// Inserir
	gavetas[0] = "Copos"
	gavetas[1] = "Pratos"
	gavetas[2] = "Panos"
	gavetas[3] = "Panelas"
	gavetas[4] = "Xicaras"

	// Leitura
	fmt.Println(gavetas[0], gavetas[1], gavetas[2], gavetas[3], gavetas[4])
}