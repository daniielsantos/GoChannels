package covergencia

import (
	"fmt"
	"math/rand"
	"time"
)

func GetConvergencia() {
	canal := converge(trabalho("Maca"), trabalho("Pera"))
	for x := 0; x < 16; x++ {
		fmt.Println(<-canal)
	}
	close(canal)
}

func trabalho(s string) chan string {
	canal := make(chan string)
	go func(s string, c chan string) {
		for i := 1; ; i++ {
			c <- fmt.Sprintf("Funcao %v diz: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, canal)
	return canal
}

func converge(x, y chan string) chan string {
	novo := make(chan string)
	go func() {
		for {
			novo <- <-x
		}
	}()

	go func() {
		for {
			novo <- <-y
		}
	}()
	return novo
}
