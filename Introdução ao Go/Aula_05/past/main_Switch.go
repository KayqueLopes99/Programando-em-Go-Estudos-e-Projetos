package past
import (
	"fmt"
	"time"
)

func Main_Switch(){
// switch case default
	fmt.Println("--- Sistema De Dias da Semana ---")
	today := time.Now().Weekday()

	switch today{
	case time.Monday:
		fmt.Println("Hoje é Segunda!")
	case time.Tuesday:
		fmt.Println("Hoje é Terça!")

	case time.Wednesday:
		fmt.Println("Hoje é Quarta!")

	case time.Thursday:
		fmt.Println("Hoje é Quinta!")

	case time.Friday:
		fmt.Println("Hoje é Sexta!")

	case time.Saturday:
		fmt.Println("Hoje é Sabado!")
		
	case time.Sunday:
		fmt.Println("Hoje é Domingo!")
	
	default:
		fmt.Println("ERRO.")


	}





}
