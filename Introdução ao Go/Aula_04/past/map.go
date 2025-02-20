package past

import "fmt"

func MainMap(){ // Flexível
	var pessoas = map[string]int{}//Chave[] e Valor{}

	// Inserir
	pessoas["Kayque"] = 20
	pessoas["Isabelly"] = 21
	fmt.Println(pessoas["Kayque"])

	// ok é boll
	if idade, ok :=  pessoas["Kayque"]; ok{
		fmt.Println("Pessoa existe no Map.", idade, ok)
	}else {
		fmt.Println("Pessoa não existe no Map.")

	}
	// Excluir
	delete(pessoas, "Isabelly")
	fmt.Println(pessoas)
}