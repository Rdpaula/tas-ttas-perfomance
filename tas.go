package main

import "sync/atomic"

type TASLock struct {
	locked int32
}

func (l *TASLock) Lock() {
	for atomic.SwapInt32(&l.locked, 1) == 1 {
		// Espera ocupada
	}
}

func (l *TASLock) Unlock() {
	atomic.StoreInt32(&l.locked, 0)
}
