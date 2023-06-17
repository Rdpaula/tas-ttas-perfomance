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

func Run(threads int, executions int, criticalLoad int, lockType string) int{
	value = 0

	var lock Lock

	switch lockType {
	case "TAS":
		lock = &TASLock{}
	case "TTAS":
		lock = &TTASLock{}
	default:
		fmt.Println("Lock nao especificado")
		os.Exit(1)
	}

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
	return value
}
