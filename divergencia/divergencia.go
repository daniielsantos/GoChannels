package divergencia

import (
	"fmt"
	"sync"
	"time"
)

func GetDivergencia() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	workers := 50
	go manda(100, canal1)
	go outra(workers, canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}
}

func manda(n int, canal chan int) {
	for x := 0; x < n; x++ {
		canal <- x
	}
	close(canal)
}

func outra(workers int, canal1, canal2 chan int) {
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * 1000) //time.Duration(rand.Intn(1e3)))
	return n
}
