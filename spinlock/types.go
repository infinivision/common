package spinlock

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type lock struct {
	v uint32
}

func New() sync.Locker {
	return &lock{0}
}

func (l *lock) Lock() {
	for !atomic.CompareAndSwapUint32(&l.v, 0, 1) {
		runtime.Gosched()
	}
}

func (l *lock) Unlock() {
	atomic.StoreUint32(&l.v, 0)
}
