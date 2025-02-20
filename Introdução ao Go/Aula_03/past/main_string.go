package past

import (
	"fmt"
	"strings"
)

func MainString(){
	var hello string = "Olá Mundo!"
	var question string = " Como você está?"

	var meet = hello + question

	// Deixar maiusculas.
	fmt.Println(strings.ToUpper(meet))
	fmt.Println(meet)
	fmt.Println(strings.Contains(meet, "Mundo"))

}