package main

import (
	"fmt"
	"runtime"
	"sync"
)

// This is inconsistent and no functional code. It will be fixed in -mutex.go-
// This is an example of data race, what we shouldn't do if we don't want to
// get random stuff on our code.
func main() {
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())

	contador := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := contador
			v++
			runtime.Gosched()
			contador = v
			wg.Done()
		}()
		fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
