package past

import (
	"fmt"
	"time"
)


func sayHello(){
	for i := 0; i < 3; i++{
		fmt.Println("Hello")
		time.Sleep(150 * time.Millisecond)
	}
}


func sayWorld(){
	for i := 0; i < 3; i++{
		fmt.Println("World")
		time.Sleep(100 * time.Millisecond)
	}
}



// Executam de maneira concorrente. 

func Main_Goreotine_Sem_Ordem() {
    go sayHello()
	go sayWorld()
	time.Sleep(1 * time.Second)
}