package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var value int

func criticalSection(lock Lock, iterations int) {
	for i := 0; i < iterations; i++ {
		lock.Lock()
		value++
		lock.Unlock()
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Uso: go run main.go <quantidade-threads> <vezes-execucao> <critical-load>")
		os.Exit(1)
	}

	threads, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Erro ao converter a quantidade de threads:", err)
		os.Exit(1)
	}

	executions, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Erro ao converter a quantidade de execuções:", err)
		os.Exit(1)
	}

	criticalLoad, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Erro ao converter o valor de critical load:", err)
		os.Exit(1)
	}

	value = 0

	lock := &TTASLock{}

	var wg sync.WaitGroup
	wg.Add(threads)

	for t := 0; t < threads; t++ {
		go func() {
			defer wg.Done()
			for i := 0; i < executions; i++ {
				criticalSection(lock, criticalLoad)
			}
		}()
	}

	wg.Wait()

	fmt.Println("Valor final:", value)
}
