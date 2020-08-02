package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Implementing mutex to solve the data race problem
func main() {
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())

	contador := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	// Declaring a mutex
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() // Locking access to other goroutines
			v := contador
			v++
			runtime.Gosched()
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Número de Goroutinas:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
