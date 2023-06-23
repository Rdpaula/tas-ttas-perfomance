package main

import (
	"sync"
	"time"
)

var value int64

func criticalSection(lock Lock, iterations int) {
	lock.Lock()
	for i := 0; i < iterations; i++ {
		value++
	}
	lock.Unlock()
}

func Run(threads int, executions int, load int, lock Lock) time.Duration {
	value = 0
	initialTime := time.Now()

	var wg sync.WaitGroup
	wg.Add(threads)

	for t := 0; t < threads; t++ {
		go func() {
			defer wg.Done()
			for i := 0; i < executions; i++ {
				criticalSection(lock, load)
			}
		}()
	}

	wg.Wait()

	return time.Since(initialTime)
}
