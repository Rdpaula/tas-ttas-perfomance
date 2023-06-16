package main

import "sync/atomic"

type TTASLock struct {
	locked int32
}

func (l *TTASLock) Lock() {
	for {
		for atomic.LoadInt32(&l.locked) == 1 {
			// Espera ocupada
		}
		if atomic.SwapInt32(&l.locked, 1) == 0 {
			return
		}
	}
}

func (l *TTASLock) Unlock() {
	atomic.StoreInt32(&l.locked, 0)
}
