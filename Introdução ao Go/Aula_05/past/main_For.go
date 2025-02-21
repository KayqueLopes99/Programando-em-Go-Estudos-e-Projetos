package past
import (
	"fmt"
)

func Main_for(){
	sum := 0

	for i := 0; i < 10; i++{
		fmt.Println(i)
	}

	// Simula o While.
	for sum < 20 {
		fmt.Println(sum)
		sum += 2
	}
	// Com slices e range.
	nums := []int{1,2,3,4,5}

	for key, value := range nums{
		fmt.Println(key, value)
	}

	// map e range
	users := map[string]string{
	"Nome": "Kayque",
	"Idade": "20",
}

for key, value := range users{
	fmt.Println(key, value)
}
	fmt.Println(sum)
}
