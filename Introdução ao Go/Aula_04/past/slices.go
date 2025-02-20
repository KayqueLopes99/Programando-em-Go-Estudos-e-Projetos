package past

import "fmt"

func MainSlices(){ // Flexível
	var gavetas []string
	gavetas = append(gavetas, "Figures de Anime", "Pokémons", "Brinquedos")

	// ver tamahho
    fmt.Println(len(gavetas))
	fmt.Println(gavetas[0:3]) // Ultimo indice indice + 1.

	// Excluir
	gavetas = gavetas[:1]


}