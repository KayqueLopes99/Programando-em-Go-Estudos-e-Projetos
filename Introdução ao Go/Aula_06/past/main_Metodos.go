package past


// Método -> Struct + Functions.
import (
	"fmt"
)

type Pessoa struct{
	Nome string
	Idade int
}


// Recebe uma cópia da struct, pois não tem ponteiro. Coloquei ponteiro para evitar essas copias. 
func (p *Pessoa) Apresentar() {
	fmt.Printf("Olá me chamo %s sou Programador de Go tenho %d Anos.\n", p.Nome, p.Idade)
}

func Main_Metodo(){	
	p1 := Pessoa{Nome: "Kayque", Idade: 20}
	p2 := Pessoa{Nome: "Samuel", Idade: 20}
	p1.Apresentar()
	p2.Apresentar()

}