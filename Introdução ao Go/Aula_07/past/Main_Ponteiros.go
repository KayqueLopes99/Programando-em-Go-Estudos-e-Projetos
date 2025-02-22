package past

import (
	"fmt"
)

type Pessoa struct{
	Name string
}




func Main_Ponteiro(){
	var p1 Pessoa = Pessoa{Name: "Kayque"}
	var p2 Pessoa = Pessoa{Name: "Daniel"}
	fmt.Println(p1.Name)

	var p3 *Pessoa = &p1 //Ponteiro para guardar no endereço & .
	p3.Name = "Samuel"

	fmt.Println(p3.Name)
	fmt.Println(p2.Name)
	
}